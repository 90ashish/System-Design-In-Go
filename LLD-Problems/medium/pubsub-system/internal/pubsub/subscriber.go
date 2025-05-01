package pubsub

// Subscriber defines the callback invoked on message arrival.
// (Observer Pattern)
type Subscriber interface {
	OnMessage(msg *Message)
}
