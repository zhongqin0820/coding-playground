package singleton

import (
	"sync"
	"time"
)

type check struct {
	value int
}

// struct method of singleton
func (s *check) AddOne() int {
	var l sync.Mutex
	l.Lock()
	defer l.Unlock()
	s.value += 1
	return s.value
}

func (s *check) GetValue() int {
	return s.value
}

// global instance in the package
var checkInstance *check

// global sync lock
var checkMutex sync.Mutex

// check-lock-check pattern
func GetCheckInstance() *check {
	if checkInstance == nil {
		checkMutex.Lock()
		defer checkMutex.Unlock()
		if checkInstance == nil {
			// delay init
			time.Sleep(5e9)
			checkInstance = new(check)
		}

	}
	return checkInstance
}
