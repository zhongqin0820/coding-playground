package singleton

import (
	"sync"
)

type singleton struct {
	value int
}

// global instance in the package
var instance *singleton

// global sync lock
var instanceMutex sync.Mutex

func GetInstance() *singleton {
	if instance == nil {
		instanceMutex.Lock()
		instance = new(singleton)
		instanceMutex.Unlock()
	}
	return instance
}

// struct method of singleton
func (s *singleton) AddOne() int {
	s.value += 1
	return s.value
}

func (s *singleton) GetValue() int {
	return s.value
}
