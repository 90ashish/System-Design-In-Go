# Designing a Pub-Sub System in Java

## Requirements

- The Pub-Sub system should allow publishers to publish messages to specific topics.
- Subscribers should be able to subscribe to topics of interest and receive messages published to those topics.
- The system should support multiple publishers and subscribers.
- Messages should be delivered to all subscribers of a topic in real-time.
- The system should handle concurrent access and ensure thread safety.
- The Pub-Sub system should be scalable and efficient in terms of message delivery.

---

## Workflow

1. **Initialize Components**  
   - Create one or more `Topic` instances.  
   - Create `Publisher` instances.  
   - Create `Subscriber` instances (e.g. `PrintSubscriber`).

2. **Register Publishers**  
   - Call `Publisher.RegisterTopic(topic)` to let each publisher know which topics it can publish to.

3. **Subscribe to Topics**  
   - Call `Topic.AddSubscriber(subscriber)` for each subscriber–topic pairing.

4. **Publish Messages**  
   - Call `Publisher.Publish(topic, message)` to send a `Message` to a topic.  
   - The topic locks its subscriber list, then invokes each subscriber’s `OnMessage` (in its own goroutine).

5. **Deliver & Handle**  
   - Each `Subscriber.OnMessage` is called concurrently, so subscribers receive and process messages in “real time.”

6. **Unsubscribe (Optional)**  
   - Call `Topic.RemoveSubscriber(subscriber)` to stop sending future messages to that subscriber.

---

## Patterns & Concurrency

- **Observer Pattern**:  
  Topic maintains a list of subscribers and notifies them via `onMessage()` when a new message is published.

- **Factory Pattern**:  
  `NewXxx` methods (e.g., `NewTopic`, `NewPublisher`, `NewSubscriber`) encapsulate object creation logic for better modularity.

- **Thread Safety**:  
  Each `Topic` uses `sync.RWMutex` (in Go) or `ReentrantReadWriteLock` (in Java) to ensure concurrent access to the subscriber set is thread-safe.

- **Scalability**:  
  Publishing spawns a new thread (or goroutine in Go) per subscriber to ensure non-blocking message delivery.