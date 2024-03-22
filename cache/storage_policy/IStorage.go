package storage_policy

// IStorage : interface to be implemented by each storage_policy class
type IStorage interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Update(key string, value string) error
	Delete(key string) error
}
