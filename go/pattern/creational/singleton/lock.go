package singleton

import (
	"sync"
	"time"
)

type lock struct {
	value int
}

// struct method of singleton
func (s *lock) AddOne() int {
	var l sync.Mutex
	l.Lock()
	defer l.Unlock()
	s.value += 1
	return s.value
}

func (s *lock) GetValue() int {
	return s.value
}

// global instance in the package
var lockInstance *lock

// global sync lock
var instanceMutex sync.Mutex

func GetLockInstance() *lock {
	instanceMutex.Lock()
	defer instanceMutex.Unlock()
	if lockInstance == nil {
		// delay init
		time.Sleep(5e9)
		lockInstance = new(lock)
	}
	return lockInstance
}
