package cache

import "github.com/personal-projects/LLD/cache_library/pkg"

type ICache interface {
	pkg.Rollback
	Set(key string, value string) error
	Get(key string) (string, error)
	Update(key string, value string) error
	Delete(key string) error
	CurrentCacheSize() (int32, error)
	RemainingCacheCapacity() (int32, error)
}
