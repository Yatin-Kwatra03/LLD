package cache

import "errors"

func GetCacheForUseCase(usecase string) (ICache, error) {
	// we can have some mapping from this usecase to cache type
	return nil, errors.New("method unimplemented")
}
