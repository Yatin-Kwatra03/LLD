package storage_policy

import "errors"

type hashmap struct {
	// add storage_policy entities
}

func newHashmap() *hashmap {
	// method unimplemented
	return nil
}

var _ IStorage = &hashmap{}

func (s *hashmap) Set(key string, value string) error {
	return errors.New("method unimplemented")
}

func (s *hashmap) Get(key string) (string, error) {
	return "", errors.New("method unimplemented")
}

func (s *hashmap) Update(key string, value string) error {
	return errors.New("method unimplemented")
}

func (s *hashmap) Delete(key string) error {
	return errors.New("method unimplemented")
}
