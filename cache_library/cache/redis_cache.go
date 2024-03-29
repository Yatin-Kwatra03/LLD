package cache

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/personal-projects/LLD/cache_library/cache/eviction_policy"
	"github.com/personal-projects/LLD/cache_library/cache/storage_policy"
	"github.com/personal-projects/LLD/cache_library/pkg"
)

type redisCache struct {
	storagePolicy  storage_policy.IStorage
	evictionPolicy eviction_policy.IEviction
	capacity       int32
	// makes more sense to use reentrant / recursive locks in our case, but it
	// is not always advisable to use reentrant locks, as they can lead to
	// deadlock.
	// we can keep the locking mechanism at db level as well.
	// But good to have it on the service level since
	// based on use case we can handle locking and unlocking accordingly.
	// Also, if we do at database level, we'll need to do for eviction policy
	// as well separately. Here we only need to do it at one place.
	mu *pkg.ReentrantLock
}

func newRedisCache(arg string, cacheCapacity int32) *redisCache {
	storageClient, err := storage_policy.GetStorage(arg)
	if err != nil {
		panic(fmt.Sprintf("error while initializing storage client : %v", err))
	}

	evictionClient, err := eviction_policy.GetEvictionPolicy(arg)
	if err != nil {
		panic(fmt.Sprintf("error while initializing eviction client : %v", err))
	}

	return &redisCache{
		storagePolicy:  storageClient,
		evictionPolicy: evictionClient,
		capacity:       cacheCapacity,
		mu:             pkg.NewReentrantLock(),
	}
}

var _ ICache = &redisCache{}

func (s *redisCache) Set(key, value string, optionalParams ...string) error {
	var (
		err error
	)

	optionalParams = populateOwnerIdIfRequired(optionalParams)
	ownerId := getLockOwnerId(optionalParams)
	s.mu.Lock(ownerId)
	defer func() {
		if err != nil {
			s.RollbackChanges()
		} else {
			s.UpdateBackUp()
		}

		s.mu.Unlock(ownerId)
	}()

	if err = s.freeUpSpaceIfRequired(optionalParams); err != nil {
		err = fmt.Errorf("error while freeing up space for the new key : %w", err)
		return err
	}

	if err = s.storagePolicy.Set(key, value); err != nil {
		err = fmt.Errorf("error while adding new key in storage : %w", err)
		return err
	}

	if err = s.evictionPolicy.NotifyGet(key); err != nil {
		err = fmt.Errorf("error while notifying eviction policy for the new key : %w", err)
		return err
	}

	return nil
}

func (s *redisCache) Get(key string, optionalParams ...string) (string, error) {
	var (
		val string
		err error
	)

	optionalParams = populateOwnerIdIfRequired(optionalParams)
	ownerId := getLockOwnerId(optionalParams)
	s.mu.Lock(ownerId)
	defer func() {
		if err != nil {
			s.RollbackChanges()
		} else {
			s.UpdateBackUp()
		}

		s.mu.Unlock(ownerId)
	}()

	if val, err = s.storagePolicy.Get(key); err != nil {
		err = fmt.Errorf("error while fetching value of key : %w", err)
		return "", err
	}

	if err = s.evictionPolicy.NotifyGet(key); err != nil {
		err = fmt.Errorf("error while notifying eviction policy for the fetched key : %w", err)
		return "", err
	}

	return val, nil
}

func (s *redisCache) Update(key, value string, optionalParams ...string) error {
	var (
		err error
	)

	optionalParams = populateOwnerIdIfRequired(optionalParams)
	ownerId := getLockOwnerId(optionalParams)
	s.mu.Lock(ownerId)
	defer func() {
		if err != nil {
			s.RollbackChanges()
		} else {
			s.UpdateBackUp()
		}

		s.mu.Unlock(ownerId)
	}()

	if err = s.storagePolicy.Update(key, value); err != nil {
		err = fmt.Errorf("error while updating key value in storage : %w", err)
		return err
	}

	if err = s.evictionPolicy.NotifyGet(key); err != nil {
		err = fmt.Errorf("error while notifying eviction policy for the updated key : %w", err)
		return err
	}

	return nil
}

func (s *redisCache) Delete(keyToBeEvicted string, optionalParams ...string) error {
	var (
		err error
	)

	optionalParams = populateOwnerIdIfRequired(optionalParams)
	ownerId := getLockOwnerId(optionalParams)
	s.mu.Lock(ownerId)
	defer func() {
		if err != nil {
			s.RollbackChanges()
		} else {
			s.UpdateBackUp()
		}

		s.mu.Unlock(ownerId)
	}()

	if err = s.storagePolicy.Delete(keyToBeEvicted); err != nil {
		err = fmt.Errorf("error while deleting key from storage")
		return err
	}

	if err = s.evictionPolicy.Evict(keyToBeEvicted); err != nil {
		err = fmt.Errorf("error while evicting key = %s : %w", keyToBeEvicted, err)
		return err
	}

	return nil
}

func (s *redisCache) CurrentCacheSize(optionalParams ...string) (int32, error) {
	optionalParams = populateOwnerIdIfRequired(optionalParams)
	ownerId := getLockOwnerId(optionalParams)
	s.mu.Lock(ownerId)
	defer s.mu.Unlock(ownerId)

	return s.storagePolicy.NoOfEntitiesCached(), nil
}

func (s *redisCache) RemainingCacheCapacity(optionalParams ...string) (int32, error) {
	optionalParams = populateOwnerIdIfRequired(optionalParams)
	ownerId := getLockOwnerId(optionalParams)
	s.mu.Lock(ownerId)
	defer s.mu.Unlock(ownerId)

	return s.capacity - s.storagePolicy.NoOfEntitiesCached(), nil
}

func (s *redisCache) freeUpSpaceIfRequired(optionalParams []string) error {
	if s.capacity == s.storagePolicy.NoOfEntitiesCached() {

		log.Println("key eviction required")

		// we need to evict a key in order to add current one
		keyToBeEvicted, err := s.evictionPolicy.GetKeyToEvict()
		if err != nil {
			return fmt.Errorf("error while fetching key to be evicted : %w", err)
		}

		log.Println("key to be evicted: ", keyToBeEvicted)

		if err = s.Delete(keyToBeEvicted, optionalParams...); err != nil {
			return fmt.Errorf("error while deleting key : %s : %w", keyToBeEvicted, err)
		}
	}
	return nil
}

func (s *redisCache) RollbackChanges() {
	s.storagePolicy.RollbackChanges()
	s.evictionPolicy.RollbackChanges()
}

func (s *redisCache) UpdateBackUp() {
	s.storagePolicy.UpdateBackUp()
	s.evictionPolicy.UpdateBackUp()
}

func populateOwnerIdIfRequired(optionalParams []string) []string {
	if len(optionalParams) == 0 {
		optionalParams = []string{strconv.Itoa(rand.Int())}
	}
	return optionalParams
}

func getLockOwnerId(optionalParams []string) string {
	return optionalParams[0]
}
