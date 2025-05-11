package main

import (
	"pubsub-system/internal/pubsub"
	"time"
)

func main() {
	// Create two topics
	topic1 := pubsub.NewTopic("Topic1")
	topic2 := pubsub.NewTopic("Topic2")

	// Create publishers (Factory Pattern)
	pub1 := pubsub.NewPublisher()
	pub2 := pubsub.NewPublisher()

	// Create subscribers
	sub1 := pubsub.NewPrintSubscriber("Subscriber1")
	sub2 := pubsub.NewPrintSubscriber("Subscriber2")
	sub3 := pubsub.NewPrintSubscriber("Subscriber3")

	// Register which topics each publisher may publish to
	pub1.RegisterTopic(topic1)
	pub2.RegisterTopic(topic2)

	// Subscribers express interest (Observer Pattern)
	topic1.AddSubscriber(sub1)
	topic1.AddSubscriber(sub2)
	topic2.AddSubscriber(sub2)
	topic2.AddSubscriber(sub3)

	// Publish some messages
	pub1.Publish(topic1, pubsub.NewMessage("Message1 for Topic1"))
	pub1.Publish(topic1, pubsub.NewMessage("Message2 for Topic1"))
	pub2.Publish(topic2, pubsub.NewMessage("Message1 for Topic2"))

	// Unsubscribe sub2 from topic1
	topic1.RemoveSubscriber(sub2)

	// Publish again
	pub1.Publish(topic1, pubsub.NewMessage("Message3 for Topic1"))
	pub2.Publish(topic2, pubsub.NewMessage("Message2 for Topic2"))

	// Give goroutines time to print
	time.Sleep(1 * time.Second)
}
