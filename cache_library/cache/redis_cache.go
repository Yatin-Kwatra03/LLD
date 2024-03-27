package cache

import (
	"fmt"
	"sync"

	"github.com/personal-projects/LLD/cache_library/cache/eviction_policy"
	"github.com/personal-projects/LLD/cache_library/cache/storage_policy"
)

const (
	defaultCacheCapacity = 5
)

type redisCache struct {
	storagePolicy  storage_policy.IStorage
	evictionPolicy eviction_policy.IEviction
	capacity       int32
	mu             sync.Mutex
}

func newRedisCache(arg string) *redisCache {
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
		capacity:       defaultCacheCapacity,
	}
}

var _ ICache = &redisCache{}

func (s *redisCache) Set(key string, value string) error {
	var (
		err error
	)

	s.mu.Lock()
	defer func() {
		if err != nil {
			s.RollbackChanges()
		} else {
			s.UpdateBackUp()
		}

		s.mu.Unlock()
	}()

	if err = s.freeUpSpaceIfRequired(); err != nil {
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

func (s *redisCache) Get(key string) (string, error) {
	var (
		val string
		err error
	)

	s.mu.Lock()
	defer func() {
		if err != nil {
			s.RollbackChanges()
		} else {
			s.UpdateBackUp()
		}

		s.mu.Unlock()
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

func (s *redisCache) Update(key string, value string) error {
	var (
		err error
	)

	s.mu.Lock()
	defer func() {
		if err != nil {
			s.RollbackChanges()
		} else {
			s.UpdateBackUp()
		}

		s.mu.Unlock()
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

func (s *redisCache) Delete(keyToBeEvicted string) error {
	var (
		err error
	)

	s.mu.Lock()
	defer func() {
		if err != nil {
			s.RollbackChanges()
		} else {
			s.UpdateBackUp()
		}

		s.mu.Unlock()
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

func (s *redisCache) CurrentCacheSize() (int32, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.storagePolicy.NoOfEntitiesCached(), nil
}

func (s *redisCache) RemainingCacheCapacity() (int32, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.capacity - s.storagePolicy.NoOfEntitiesCached(), nil
}

func (s *redisCache) freeUpSpaceIfRequired() error {
	if s.capacity == s.storagePolicy.NoOfEntitiesCached() {
		// we need to evict a key in order to add current one
		keyToBeEvicted, err := s.evictionPolicy.GetKeyToEvict()
		if err != nil {
			return fmt.Errorf("error while fetching key to be evicted : %w", err)
		}

		if err = s.Delete(keyToBeEvicted); err != nil {
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
