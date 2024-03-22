package eviction_policy

// Eviction : is a factory of implementations of different eviction policies
type Eviction struct {
	lru *lru
	// extensible
}

func newEviction() *Eviction {
	// method unimplemented
	return nil
}
