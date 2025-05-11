package pubsub

import "fmt"

// PrintSubscriber simply prints each message it receives.
type PrintSubscriber struct {
	Name string
}

// NewPrintSubscriber constructs a PrintSubscriber (Factory Pattern).
func NewPrintSubscriber(name string) *PrintSubscriber {
	return &PrintSubscriber{Name: name}
}

// OnMessage is called by Topic.Publish (Observer callback).
func (ps *PrintSubscriber) OnMessage(msg *Message) {
	fmt.Printf("[%s] received: %s\n", ps.Name, msg.Content)
}
