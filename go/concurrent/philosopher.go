package concurrent

import (
	"log"
	"math/rand"
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
