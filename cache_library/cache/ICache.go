package cache

import "github.com/personal-projects/LLD/cache_library/pkg"

// ICache : is the interface exposed to internal client who wants
// to use the cache library. They will be able to use all the below
// mention methods.
// Note - optionalParams is required for internal lock handling.
// Client must not pass it while performing any cache operation
// Con: If client passes some random id in the request. Our lock
// implementation can get messed up.
// Solve: ownerId logic will abstract from client. So we can generate
// and finalize some format for that id, if id format is not honoured then
// we'll discard and generate our own id or maybe fail the call since client
// is trying to mess with the system.
type ICache interface {
	pkg.Rollback
	Set(key, value string, optionalParams ...string) error
	Get(key string, optionalParams ...string) (string, error)
	Update(key, value string, optionalParams ...string) error
	Delete(key string, optionalParams ...string) error
	CurrentCacheSize(optionalParams ...string) (int32, error)
	RemainingCacheCapacity(optionalParams ...string) (int32, error)
}
