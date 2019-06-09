package snippet

import (
	"sync"
	"testing"
	"time"
)

var lock = &sync.Mutex{}

type DriverPg struct {
	conn string
}

var instance *DriverPg

func Connect() *DriverPg {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &DriverPg{conn: "DriverConnectPostgres"}
	}

	return instance
}

func TestConnect(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(20)
	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(1e9)
			t.Log(*Connect(), " - ", i)
			wg.Done()
		}
	}()

	go func() {
		t.Log(*Connect())
	}()

	wg.Wait()
}
