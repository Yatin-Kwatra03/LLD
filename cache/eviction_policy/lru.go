package eviction_policy

import (
	"errors"
	"sync"
)

type node struct {
	data int
	next *node
	last *node
}

func newNode(data int) *node {
	return &node{
		data: data,
		next: nil,
		last: nil,
	}
}

type lru struct {
	head          *node
	tail          *node
	nodeReference map[string]*node
	mu            sync.Mutex
}

func newLru() *lru {
	return &lru{
		head:          nil,
		tail:          nil,
		nodeReference: make(map[string]*node),
	}
}

var _ IEviction = &lru{}

func (s *lru) NotifyGet(key string) error {
	return errors.New("method unimplemented")
}

func (s *lru) Evict(key string) error {
	return errors.New("method unimplemented")
}

func (s *lru) GetKeyToEvict(string) (string, error) {
	return "", errors.New("method unimplemented")
}
