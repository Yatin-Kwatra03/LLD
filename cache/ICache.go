package cache

type ICache interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Update(key string, value string) error
	Delete(key string) error
	CurrentCacheSize() (int32, error)
	RemainingCacheCapacity() (int32, error)
}
