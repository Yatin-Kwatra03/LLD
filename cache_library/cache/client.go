package cache

import "errors"

func GetCache(arg1, arg2 string) (ICache, error) {
	// all the entities will be initialized based on the arg1
	// which can represent a useCase based on which internal
	// entities will be initialized
	cacheClient := newCache(arg1)

	switch arg2 {
	case "redis cache":
		return cacheClient.redisCache, nil
	default:
		return nil, errors.New("no cache supported for the given parameters")
	}
}
