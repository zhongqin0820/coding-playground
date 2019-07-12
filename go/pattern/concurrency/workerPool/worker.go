package workerPool

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

type Request struct {
	Data    interface{}
	handler RequestHandler
}

type RequestHandler func(interface{})

func NewStringRequest(s string, id int, wg *sync.WaitGroup) Request {
	return Request{
		Data: fmt.Sprintf(s, id),
		handler: func(i interface{}) {
			defer wg.Done()
			s, ok := i.(string)
			if !ok {
				log.Fatal("Invalid casting to string")
			}
			fmt.Println(s)
		},
	}
}

// WorkerLauncher interface
type WorkerLauncher interface {
	// in chan Request is the single entrance point to the pipeline.
	LaunchWorker(in chan Request)
}

// worker pipeline
type PreffixSuffixWorker struct {
	id      int
	prefixS string
	suffixS string
}

func (w *PreffixSuffixWorker) LaunchWorker(in chan Request) {
	// each work is a pipeline of Request(pipeline pattern)
	w.prefix(w.append(w.uppercase(in)))
}

func (w *PreffixSuffixWorker) uppercase(in <-chan Request) <-chan Request {
	out := make(chan Request)
	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)
			if !ok {
				msg.handler(nil)
				continue
			}
			msg.Data = strings.ToUpper(s)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PreffixSuffixWorker) append(in <-chan Request) <-chan Request {
	out := make(chan Request)
	go func() {
		for msg := range in {
			uppercaseString, ok := msg.Data.(string)
			if !ok {
				msg.handler(nil)
				continue
			}
			msg.Data = fmt.Sprintf("%s%s", uppercaseString, w.suffixS)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PreffixSuffixWorker) prefix(in <-chan Request) {
	go func() {
		for msg := range in {
			uppercasedStringWithSuffix, ok := msg.Data.(string)
			if !ok {
				msg.handler(nil)
				continue
			}
			msg.handler(fmt.Sprintf("%s%s", w.prefixS,
				uppercasedStringWithSuffix))
		}
	}()
}
