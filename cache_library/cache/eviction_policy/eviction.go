package eviction_policy

// EvictionFactory : is a factory of implementations of different eviction policies
// so we can have n number of concrete implementations for it.
type EvictionFactory struct {
	lru *lru
}

func newEviction() *EvictionFactory {
	return &EvictionFactory{
		lru: newLru(),
	}
}
