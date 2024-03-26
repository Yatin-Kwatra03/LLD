package eviction_policy

// IEviction : any type of eviction policy will need to add support
// for these functionalities
type IEviction interface {
	NotifyGet(key string) error
	Evict(key string) error
	GetKeyToEvict() (string, error)
}
