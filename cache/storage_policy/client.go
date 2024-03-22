package storage_policy

import (
	"errors"
)

func GetStorageForUseCase(useCase string) (IStorage, error) {
	// for cache use case we can use hashMap implementation
	return nil, errors.New("method unimplemented")
}
