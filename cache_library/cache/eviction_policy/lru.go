package eviction_policy

import (
	"errors"
)

type node struct {
	key  string
	next *node
	last *node
}

func newNode(key string) *node {
	return &node{
		key:  key,
		next: nil,
		last: nil,
	}
}

type lru struct {
	head          *node
	tail          *node
	nodeReference map[string]*node

	// in order to maintain data integrity we'll need some rollback strategy
	// so, we'll keep backup of all the data, if any failure happens, we'll revert
	// the changes by using the backup data
	backUpHead          *node
	backUpTail          *node
	backUpNodeReference map[string]*node
}

func newLru() *lru {
	return &lru{
		head:                nil,
		tail:                nil,
		nodeReference:       make(map[string]*node),
		backUpHead:          nil,
		backUpTail:          nil,
		backUpNodeReference: make(map[string]*node),
	}
}

var _ IEviction = &lru{}

// NotifyGet : will be used as the common method to inform the eviction policy
// if any new element is added or an old element is accessed. Handling for both
// the cases will be same for eviction, because for eviction policy the current
// element will become the most recent accessed element.
func (s *lru) NotifyGet(key string) error {
	existingEntity, ok := s.nodeReference[key]
	if !ok { // it's a new entity
		return s.addElementToTail(newNode(key))
	}

	isCurrentNodeHead := existingEntity == s.head
	isCurrentNodeTail := existingEntity == s.tail

	switch {
	case isCurrentNodeHead && isCurrentNodeTail:
		// no need to do anything since it's the only element
		return nil
	case isCurrentNodeHead:
		s.head = s.head.next
		s.head.last = nil
		return s.addElementToTail(existingEntity)

	case isCurrentNodeTail:
		// since it is already at the last / most recently accessed element
		return nil

	default:
		// link my last to my next (
		existingEntity.last.next = existingEntity.next
		// link my next to my last
		existingEntity.next.last = existingEntity.last
		// add me to the last
		return s.addElementToTail(existingEntity)
	}

}

func (s *lru) RollbackChanges() {
	s.head = s.backUpHead
	s.tail = s.backUpTail
	s.nodeReference = s.backUpNodeReference
}

func (s *lru) UpdateBackUp() {
	s.backUpHead = s.head
	s.backUpTail = s.tail
	s.backUpNodeReference = s.nodeReference
}

func (s *lru) Evict(key string) error {
	existingEntity, ok := s.nodeReference[key]
	if !ok {
		return errors.New("key doesn't exist")
	}

	isCurrentNodeHead := existingEntity == s.head
	isCurrentNodeTail := existingEntity == s.tail

	switch {
	case isCurrentNodeHead && isCurrentNodeTail:
		s.head = nil
		s.tail = nil
		delete(s.nodeReference, key)
	case isCurrentNodeHead:
		s.head = s.head.next
		s.head.last = nil
		delete(s.nodeReference, key)

	case isCurrentNodeTail:
		s.tail = s.tail.last
		s.tail.next = nil
		delete(s.nodeReference, key)
	default:
		// link my last to my next (
		existingEntity.last.next = existingEntity.next
		// link my next to my last
		existingEntity.next.last = existingEntity.last
	}

	return nil
}

func (s *lru) GetKeyToEvict() (string, error) {
	if len(s.nodeReference) == 0 {
		return "", errors.New("no key found in record")
	}
	return s.head.key, nil
}

func (s *lru) addElementToTail(addNode *node) error {
	// first element case
	if s.tail == nil {
		addNode.last = nil
		addNode.next = nil
		s.head = addNode
		s.tail = addNode
		s.nodeReference[addNode.key] = addNode

		return nil
	}

	// add to last
	s.tail.next = addNode
	// link last to current tail
	addNode.last = s.tail
	// make next of tail to nil
	addNode.next = nil

	// update tail
	s.tail = addNode

	// update in the node reference map
	s.nodeReference[addNode.key] = addNode

	return nil
}
