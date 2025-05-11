package pubsub

import "fmt"

// Publisher knows which Topics it may publish to.
type Publisher struct {
	topics map[*Topic]struct{}
}

// NewPublisher constructs a Publisher (Factory Pattern).
func NewPublisher() *Publisher {
	return &Publisher{topics: make(map[*Topic]struct{})}
}

// RegisterTopic grants this publisher the right to publish to topic.
func (p *Publisher) RegisterTopic(t *Topic) {
	p.topics[t] = struct{}{}
}

// Publish sends a message to a Topic if registered.
func (p *Publisher) Publish(t *Topic, msg *Message) {
	if _, ok := p.topics[t]; !ok {
		fmt.Printf("Publisher not registered for topic %s\n", t.Name)
		return
	}
	t.Publish(msg)
}
