package advanced

import (
	"sync"
	"time"
)

// Clip codes from "github.com/docker/docker/pkg/pubsub"
type (
	subscriber chan interface{}
	//receives a content same as the definition in subscriber
	topicFunc func(v interface{}) bool
)

type Publisher struct {
	m           sync.RWMutex
	buffer      int
	timeout     time.Duration
	subscribers map[subscriber]topicFunc
}

// NewPublisher new an object of Publisher
// It contains a mapping of subscriber and the topic
func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     publishTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// Subscribe add a new publisher to subscribe all topics
func (p *Publisher) Subscibe() chan interface{} {
	return p.SubscribeTopic(nil)
}

// SubscribeTopic add a new publisher to filter the topic
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

// Evict delete a specific subscriber from the subscribers in publisher
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.n.Unlock()
	delete(p.subscribers, sub) //delete
	close(sub)                 //make the chan closed!
}

// Publish pulished a new topic
func (p *Publisher) Pulish(v interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	var wg sync.WaitGroup

	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}

	wg.Wait()
}

// Close closes all the subscriber from p
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

// sendTopic send a topicFunc(message) to a specific subcriber
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub <- v:
	case <-time.After(p.timeout):
		return
	}
}

// the usage of pub/sub
func main() {
	p := NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	all := p.Subscibe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		// v.(type) 进行类型检查
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})
	p.Pulish("hello, world!")
	p.Pulish("hello, golang")
	go func() {
		for msg := range golang {
			fmt.Println("golang: ", msg)
		}
	}()
	time.Sleep(3 * time.Second)
}
