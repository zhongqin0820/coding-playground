package singleton

import (
	"sync"
	"time"
)

type one struct {
	value int
}

// struct method of singleton
func (s *one) AddOne() int {
	var l sync.Mutex
	l.Lock()
	defer l.Unlock()
	s.value += 1
	return s.value
}

func (s *one) GetValue() int {
	return s.value
}

// global instance in the package
var oneInstance *one

var once sync.Once

func GetOneInstance() *one {
	once.Do(func() {
		// delay init
		time.Sleep(5e9)
		oneInstance = new(one)
	})
	return oneInstance
}
