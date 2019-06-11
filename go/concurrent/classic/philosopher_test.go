package classic

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

type Philosoper struct {
	name     string
	chop     chan bool
	neighbor *Philosoper
}

func newPhilosoper(name string, neighbor *Philosoper) *Philosoper {
	phil := &Philosoper{name: name, chop: make(chan bool, 1), neighbor: neighbor}
	phil.chop <- true
	return phil
}

func (phil *Philosoper) think() {
	log.Printf("%s is thinking\n", phil.name)
	time.Sleep(time.Duration(rand.Int63n(1e9)))
}

func (phil *Philosoper) eat() {
	log.Printf("%s is eating\n", phil.name)
	time.Sleep(time.Duration(rand.Int63n(1e9)))
}

//
func (phil *Philosoper) getChop() {
	timeout := make(chan bool, 1)
	go func() { time.Sleep(1e9); timeout <- true }()
	<-phil.chop
	log.Printf("%s got his chopstick\n", phil.name)
	select {
	case <-phil.neighbor.chop:
		log.Printf("%s got %s's chopstick\n", phil.name, phil.neighbor.name)
	case <-timeout:
		// failed to get chop, need to return
		phil.chop <- true
		phil.think()
		phil.getChop()
	}
}

//
func (phil *Philosoper) returnChop() {
	phil.chop <- true
	phil.neighbor.chop <- true
}

// dine is the pipeline of the scheduler
func (phil *Philosoper) dine(announce chan<- *Philosoper) {
	phil.think()
	phil.getChop()
	phil.eat()
	phil.returnChop()
	announce <- phil
}

// go test -v -run=TestPhilosopher
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
