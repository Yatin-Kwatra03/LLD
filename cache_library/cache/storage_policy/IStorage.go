package storage_policy

import "github.com/personal-projects/LLD/cache_library/pkg"

// IStorage : interface to be implemented by each storage_policy class
type IStorage interface {
	pkg.Rollback
	Set(key string, value string) error
	Get(key string) (string, error)
	Update(key string, value string) error
	Delete(key string) error
	NoOfEntitiesCached() int32
}
