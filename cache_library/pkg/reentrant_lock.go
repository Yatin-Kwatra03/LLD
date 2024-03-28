package pkg

import "sync"

// ReentrantLock / Recursive locks: use case for a reentrant lock arises
// when a function or method needs to acquire a lock multiple times within
// its execution. This is particularly useful in scenarios where the function
// might call other functions or methods that also need to acquire the same
// lock. By using a reentrant lock, a function can acquire the lock at the
// beginning of its execution and release it at the end, without worrying
// about deadlocks or blocking itself.
type ReentrantLock struct {
	mu     sync.Mutex
	cond   *sync.Cond
	locked bool
	owner  string // Change owner to string to hold the UUID
	count  int
}

func newReentrantLock() *ReentrantLock {
	reentrantLock := &ReentrantLock{}
	reentrantLock.cond = sync.NewCond(&reentrantLock.mu)
	return reentrantLock
}

func (s *ReentrantLock) Lock(ownerId string) {
	if !s.locked {
		s.mu.Lock()
		s.locked = true
		s.owner = ownerId
		s.count = 1
		return
	}

	// check if the owner is same, we can add a recurive lock
	if s.owner == ownerId {
		s.mu.Lock()
		s.count++
		return
	}

	for s.locked {
		s.cond.Wait()
	}

	s.mu.Lock()
	s.locked = true
	s.count++
	s.owner = ownerId
}

func (s *ReentrantLock) Unlock(ownerId string) {
	// ideally owner id should be same
	if !s.locked || s.owner != ownerId {
		return
	}

	s.mu.Unlock()
	s.count--
	if s.count == 0 {
		s.locked = false
		s.owner = "" // resets the owner
		// notifies the other waiting goroutines that lock is released
		s.cond.Signal()
	}
}
