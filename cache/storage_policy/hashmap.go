package storage_policy

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

var (
	keyNotFoundErr = errors.New("key not found in cache")
)

// hashmap : hashmap is one of the storage policy that
// can be used to store data in cache flows. Hashmap
// provide in memory cache support. (data is stored
// on RAM, so very fast access)
type hashmap struct {
	db map[string]string
	mu sync.Mutex
}

func newHashmap() *hashmap {
	return &hashmap{
		db: make(map[string]string),
	}
}

var _ IStorage = &hashmap{}

func (s *hashmap) Set(key string, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if key == "" {
		return errors.New("key cannot be empty")
	}
	// set the value in cache
	s.db[key] = value
	return nil
}

func (s *hashmap) Get(key string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// retrieve value from cache
	val, ok := s.db[key]
	if !ok {
		return "", keyNotFoundErr
	}
	return val, nil
}

func (s *hashmap) Update(key string, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// first we should check, does the value even exists in cache
	oldVal, err := s.Get(key)
	if err != nil {
		return err
	}

	log.Println(fmt.Sprintf("replacing the old value : %s with new value : %s for the key : %s", oldVal, value, key))

	// update the value in cache
	s.db[key] = value
	return nil
}

func (s *hashmap) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// we will check if the key exists in cache or not
	_, err := s.Get(key)
	if err != nil {
		// check if err is key not found, then we can gracefully handle it, since we wanted to delete the key anyway
		if err == keyNotFoundErr {
			log.Println(fmt.Sprintf("key %s doesn't exists in cache", key))
			return nil
		}
		return err
	}

	delete(s.db, key)
	return nil
}

func (s *hashmap) NoOfEntitiesCached() int32 {
	s.mu.Lock()
	defer s.mu.Unlock()

	return int32(len(s.db))
}
