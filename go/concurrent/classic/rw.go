package concurrent

import (
	"sync"
	"time"
)

func sleep() {
	time.Sleep(10 * time.Second)
}

func NaiveReader(c chan int, m *sync.RWMutex, wg *sync.WaitGroup) {
	sleep()
	m.RLock()
	c <- 1
	sleep()
	c <- -1
	m.RUnlock()
	wg.Done()

}

func NaiveWriter(c chan int, m *sync.RWMutex, wg *sync.WaitGroup) {
	sleep()
	m.Lock()
	c <- 1
	sleep()
	c <- -1
	m.Unlock()
	wg.Done()
}
