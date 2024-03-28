package cache

// cache : this cache is a factory and
// can have n number of concrete implementations
// for the cache factory
type cache struct {
	redisCache *redisCache
}

func newCache(arg string, cacheCapacity int32) *cache {
	return &cache{
		redisCache: newRedisCache(arg, cacheCapacity),
	}
}
