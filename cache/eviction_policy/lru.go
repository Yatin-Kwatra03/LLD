package eviction_policy

import "errors"

type lru struct {
	// add storage entites
}

func newLru() *lru {
	// method unimplemented
	return nil
}

var _ IEviction = &lru{}

func (s *lru) NotifyGet(key string) error {
	return errors.New("method unimplemented")
}

func (s *lru) Evict(key string) error {
	return errors.New("method unimplemented")
}

func (s *lru) GetKeyToEvict(string) (string, error) {
	return "", errors.New("method unimplemented")
}
