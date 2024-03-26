package storage_policy

import (
	"errors"
	"fmt"
)

// StorageFactory :
// - factory to decide the type of cache entity client will need
// - we can add n number of concrete implementations here
type StorageFactory struct {
	hashmap *hashmap
}

func NewStorageFactory() *StorageFactory {
	return &StorageFactory{
		hashmap: newHashmap(),
	}
}

type StorageType int32

const (
	StorageType_STORAGE_TYPE_UNSPECIFIED StorageType = 0
	StorageType_STORAGE_TYPE_HASHMAP     StorageType = 1
	StorageType_STORAGE_TYPE_SET         StorageType = 2
	StorageType_STORAGE_TYPE_DATABASE    StorageType = 2
)

func (s *StorageFactory) getStorage(storageType StorageType) (IStorage, error) {
	switch storageType {
	case StorageType_STORAGE_TYPE_HASHMAP, StorageType_STORAGE_TYPE_UNSPECIFIED:
		return s.hashmap, nil
	default:
		return nil, errors.New(fmt.Sprintf("no implementation found for the %s storage type", string(storageType)))
	}
}
