package eviction_policy

import "github.com/personal-projects/LLD/cache_library/pkg"

// IEviction : any type of eviction policy will need to add support
// for these functionalities
type IEviction interface {
	pkg.Rollback
	NotifyGet(key string) error
	Evict(key string) error
	GetKeyToEvict() (string, error)
}
