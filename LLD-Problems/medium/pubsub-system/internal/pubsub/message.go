package pubsub

// Message represents the payload sent via topics.
type Message struct {
	Content string
}

// NewMessage is a simple factory for Message.
func NewMessage(content string) *Message {
	return &Message{Content: content}
}
