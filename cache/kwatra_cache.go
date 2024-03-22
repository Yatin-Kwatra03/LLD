package cache

import "errors"

type kwatraCache struct {
}

var _ ICache = &kwatraCache{}

func (s *kwatraCache) Set(key string, value string) error {
	return errors.New("method unimplemented")
}

func (s *kwatraCache) Get(key string) (string, error) {
	return "", errors.New("method unimplemented")

}

func (s *kwatraCache) Update(key string, value string) error {
	return errors.New("method unimplemented")
}

func (s *kwatraCache) Delete(key string) error {
	return errors.New("method unimplemented")
}

func (s *kwatraCache) CurrentCacheSize() (int32, error) {
	return 0, errors.New("method unimplemented")
}

func (s *kwatraCache) RemainingCacheCapacity() (int32, error) {
	return 0, errors.New("method unimplemented")
}
