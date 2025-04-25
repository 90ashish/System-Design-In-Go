4. Messaging & Asynchronous Processing

Messaging and async processing let services talk without waiting on each other, improving resilience and scalability. Below is a breakdown of key concepts, each with a simple example.
4.1 Message Queues

A message queue is a buffer that holds messages until consumers are ready to process them.
System	Type	Strengths	Simple Example
Kafka	Distributed log	High throughput, long retention	Stream user activity to topics
RabbitMQ	Broker with queues	Rich routing, easy setup	Distribute tasks to worker pools
NATS	Lightweight broker	Ultra-low latency, simple API	Real-time sensor updates
SQS	Managed queue	Fully managed on AWS, unlimited scale	Buffer webhook events

    Example Flow (RabbitMQ)

        Producer: publishes a task message to order_queue.

        Broker: holds messages in order_queue.

        Worker(s): pull messages, process orders, then ACK to remove them.

Producer --> [ order_queue ] --> Worker A
                                 Worker B

4.2 Pub/Sub

Publish/Subscribe decouples senders (“publishers”) from receivers (“subscribers”) via topics.

    Google Pub/Sub: Fully managed, global.

    AWS SNS + SQS: SNS publishes to multiple SQS queues or HTTP endpoints.

    Example (AWS SNS + SQS)

        SNS Topic “user-events”

        Two SQS queues subscribe:

            email-service-queue

            analytics-service-queue

        A user signup event is published; both services receive it independently.

             ┌───────────────┐
Publisher →  │  SNS Topic    │
             └───────────────┘
              │           │
              ↓           ↓
     ┌────────────────┐ ┌────────────────┐
     │ email-service  │ │ analytics-svc  │
     └────────────────┘ └────────────────┘

4.3 Event-Driven Patterns
Event Sourcing

    What it is: Store every change as an immutable event, then rebuild state by replaying events.

    Use Case: Audit trails, temporal queries.

    Example:

        User signs up → UserCreated{userId, name, email}

        User changes email → UserEmailChanged{userId, newEmail}

    The current profile is the result of applying all events in order.

CQRS (Command Query Responsibility Segregation)

    What it is: Separate models for writes (commands) and reads (queries).

    Use Case: Complex domains where read and write demands differ greatly.

    Example:

        Command Model: Handles CreateOrder and CancelOrder.

        Query Model: Optimized read-only tables or caches for “list my orders” queries.

4.4 Processing Semantics

How we guarantee messages are handled correctly, even if failures occur.

    Idempotency:

        Definition: Performing the same operation multiple times has the same effect as once.

        Example: Credit an account only once per transaction ID.

    Delivery Guarantees:

        At-Least-Once: Messages may be delivered more than once (need idempotent consumers).

        At-Most-Once: Messages are never redelivered, but some may be lost.

        Exactly-Once: Delivered and processed once; hard to achieve without special protocols.

    Example (At-Least-Once with Kafka)

        Producer writes to topic, retries on failure.

        Consumer commits offsets only after processing, so a crash can replay messages.

4.5 Error Handling

Robust systems automatically retry and quarantine failed messages.

    Retries

        Simple backoff: Retry after fixed or exponential delay.

        Example Pseudocode:

        for attempt in 1..3:
            if process(message): break
            wait(2**attempt)  # exponential backoff

    Dead-Letter Queues (DLQ)

        What it is: A special queue for messages that repeatedly fail.

        Use Case: Inspect or reprocess later without blocking the main queue.

        Example: After 5 retries, move to order-processing-dlq.

[ order_queue ] --failures--> [ retry x5 ] --fail--> [ order_dlq ]

Further Reading

    Kafka: The Definitive Guide by Neha Narkhede et al. — deep dive into Kafka’s architecture, APIs, and best practices.