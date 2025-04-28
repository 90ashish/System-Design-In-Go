# Google System Design Interview Question: Notification System Using OCP

## Question

Design a notification system that supports multiple channels (such as Email, SMS, and Push Notifications) using the **Open/Closed Principle (OCP)**. The system should allow adding new notification channels without modifying the existing code. Explain your design and implement a sample solution in Golang.

---

## Solution in Golang

The key idea is to create an abstraction (interface) for sending notifications. Each notification channel (Email, SMS, Push) implements this interface. The main notification service depends on the abstraction, so new channels can be added easily by implementing the interface, without changing the service.

---

## Explanation

### Abstraction via Interface

- The `Notification` interface defines a single method `Send()`, which each notification channel implements.  
- This abstraction allows the `NotificationService` to remain agnostic of the underlying implementation.

### Implementations

- `EmailNotification`, `SMSNotification`, and `PushNotification` are concrete implementations of the `Notification` interface.  
- Each one encapsulates its own logic for sending notifications.  
- To add a new notification channel (e.g., `WhatsAppNotification`), you only need to implement the `Notification` interface in a new struct without changing existing code.

### NotificationService

- The `NotificationService` depends on the `Notification` interface, following the **Dependency Inversion Principle**.  
- It calls the `Send()` method on the interface, making it open for extension but closed for modification (**OCP**).

### Usage in Main

- The `main` function demonstrates how to use the `NotificationService` with different notification channels.  
- The service is extended by passing different implementations of the `Notification` interface.

---

## Benefits of OCP in This Design

- **Extensibility**: New notification channels can be added by simply implementing the `Notification` interface without altering the `NotificationService` or other existing code.
- **Maintainability**: Each notification channel has its own implementation, making it easier to modify or fix issues related to a specific channel without affecting the rest of the system.
- **Scalability**: The system can easily scale by adding additional functionality to new channels without disrupting the existing architecture.
