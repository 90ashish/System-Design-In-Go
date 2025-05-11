package pubsub

import "sync"

// Topic holds a set of subscribers and delivers messages to them.
// Thread-safe via RWMutex.
type Topic struct {
	Name        string
	subscribers map[Subscriber]struct{}
	mu          sync.RWMutex
}

// NewTopic constructs a Topic (Factory Pattern).
func NewTopic(name string) *Topic {
	return &Topic{
		Name:        name,
		subscribers: make(map[Subscriber]struct{}),
	}
}

// AddSubscriber registers a Subscriber (Observer Pattern).
func (t *Topic) AddSubscriber(s Subscriber) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.subscribers[s] = struct{}{}
}

// RemoveSubscriber unregisters a Subscriber.
func (t *Topic) RemoveSubscriber(s Subscriber) {
	t.mu.Lock()
	defer t.mu.Unlock()
	delete(t.subscribers, s)
}

// Publish sends message to all current subscribers concurrently.
func (t *Topic) Publish(msg *Message) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	for s := range t.subscribers {
		// Optionally dispatch in separate goroutines for real async:
		s.OnMessage(msg) // add goroutine
	}
}
