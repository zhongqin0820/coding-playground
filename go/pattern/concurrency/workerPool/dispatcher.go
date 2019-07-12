package workerPool

import "time"

// launch workers in parallel
// handle all the possible incoming channels
// In this case, the Dispatcher interface is merely acting as a Facade design pattern hiding some implementation details from the user.
type Dispatcher interface {
	// start Worker do Request
	LaunchWorker(WorkerLauncher)
	// hold the Request work queue
	MakeRequest(Request)
	// closes the incoming requests channel, provoking a chain reaction
	Stop()
}

//
func NewDispatcher(b int) Dispatcher {
	return &dispatcher{
		inCh: make(chan Request, b),
	}
}

// stores a channel of Request type
type dispatcher struct {
	inCh chan Request
}

func (d *dispatcher) LaunchWorker(w WorkerLauncher) {
	w.LaunchWorker(d.inCh)
}

func (d *dispatcher) Stop() {
	close(d.inCh)
}

func (d *dispatcher) MakeRequest(r Request) {
	select {
	case d.inCh <- r:
	case <-time.After(time.Second * 5):
		return
	}
}
