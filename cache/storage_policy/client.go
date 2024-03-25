package storage_policy

// GetStorage : provides the storage as per the use case of the client
// note - arg is a very generic way to decide the type of cache required.
// actually logic could be pretty complex as well
func GetStorage(arg string) (IStorage, error) {
	storageClient := NewStorageFactory()
	switch arg {
	case "cache vendor data":
		return storageClient.getStorage(StorageType_STORAGE_TYPE_HASHMAP)
	case "permanent storage of data":
		return storageClient.getStorage(StorageType_STORAGE_TYPE_DATABASE)
	default:
		return storageClient.getStorage(StorageType_STORAGE_TYPE_UNSPECIFIED)
	}
}
