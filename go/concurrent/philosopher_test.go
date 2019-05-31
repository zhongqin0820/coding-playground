package concurrent

import (
	"testing"
)

// go test -test.run TestPhilosopher -v
func TestPhilosopher(t *testing.T) {
	names := []string{"A", "B", "C", "D", "E"}
	philosophers := make([]*Philosoper, len(names))
	var phil *Philosoper
	for i, name := range names {
		phil = newPhilosoper(name, phil)
		philosophers[i] = phil
	}
	philosophers[0].neighbor = phil
	t.Logf("There are %v philosophers sitting at the table\n", len(names))
	//
	announce := make(chan *Philosoper)
	//
	for _, phil := range philosophers {
		go phil.dine(announce)
	}
	//
	for i := 0; i < len(names); i++ {
		phil := <-announce
		t.Logf("%s is done dining.\n", phil.name)
	}
}
