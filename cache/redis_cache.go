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

	return errors.New("method unimplemented")
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

func (s *redisCache) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return errors.New("method unimplemented")
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
