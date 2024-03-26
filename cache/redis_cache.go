package cache

import (
	"errors"
	"fmt"
	"sync"

	"github.com/personal-projects/LLD/cache/eviction_policy"
	"github.com/personal-projects/LLD/cache/storage_policy"
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
	s.mu.Lock()
	defer s.mu.Unlock()

	// todo: add transactional block to maintain data consistency

	if err := s.freeUpSpaceIfRequired(); err != nil {
		return fmt.Errorf("error while freeing up space for the new key")
	}

	if err := s.storagePolicy.Set(key, value); err != nil {
		return fmt.Errorf("error while adding new key in storage : %w", err)
	}

	if err := s.evictionPolicy.NotifyGet(key); err != nil {
		return fmt.Errorf("error while notifying eviction policy for the new key : %w", err)
	}

	return nil
}

func (s *redisCache) Get(key string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return "", errors.New("method unimplemented")

}

func (s *redisCache) Update(key string, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return errors.New("method unimplemented")
}

func (s *redisCache) Delete(keyToBeEvicted string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// todo: add transactional block to maintain data consistency

	if err := s.storagePolicy.Delete(keyToBeEvicted); err != nil {
		return fmt.Errorf("error while deleting key from storage")
	}

	if err := s.evictionPolicy.Evict(keyToBeEvicted); err != nil {
		return fmt.Errorf("error while evicting key = %s : %w", keyToBeEvicted, err)
	}

	return nil
}

func (s *redisCache) CurrentCacheSize() (int32, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return 0, errors.New("method unimplemented")
}

func (s *redisCache) RemainingCacheCapacity() (int32, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return 0, errors.New("method unimplemented")
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
