package singleton

import (
	"sync"
	"time"
)

// definition of simpleSingleton
type simple struct {
	value int
}

// struct method of singleton
func (s *simple) AddOne() int {
	var l sync.Mutex
	l.Lock()
	defer l.Unlock()
	s.value += 1
	return s.value
}

func (s *simple) GetValue() int {
	return s.value
}

// simpleInstance is a private variable in singleton package
var simpleInstance *simple

func GetSimpleInstance() *simple {
	if simpleInstance == nil {
		// delay init
		time.Sleep(5e9)
		simpleInstance = new(simple)
	}
	return simpleInstance
}
