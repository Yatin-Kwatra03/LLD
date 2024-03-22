package storage_policy

// StorageFactory : factory to decide the type
// of cache entity client will need
type StorageFactory struct {
	hashmap *hashmap
	// extensible
}

func NewStorageFactory() *StorageFactory {
	// method unimplemented
	return nil
}

type StorageType int32

const (
	StorageType_STORAGE_TYPE_UNSPECIFIED StorageType = 0
	StorageType_STORAGE_TYPE_HASHMAP     StorageType = 1
	StorageType_STORAGE_TYPE_SET         StorageType = 2
)

// getStorage : client exposed method to provide storage_policy entity
// based on the storage_policy type
func getStorage(storageType StorageType) IStorage {
	// return storage_policy entity based on the storage_policy type
	return nil
}
