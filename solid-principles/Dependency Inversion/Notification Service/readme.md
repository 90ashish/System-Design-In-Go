# Amazon System Design Interview Question: Order Status Notification Service Using the Dependency Inversion Principle

---

## Question:

Amazon needs a service that updates an order’s status and then notifies the customer through their preferred channel (**Email**, **SMS**, **Push**, etc.).  
Applying the **Dependency Inversion Principle (DIP)**, design this system so that:

- The **high-level module** (`OrderService`) does not depend on concrete notification channels.
- Both high-level (`OrderService`) and low-level modules (e.g., `EmailNotifier`, `SMSNotifier`) depend on an **abstraction** (`Notifier` interface).
- **Abstractions do not depend on details**; details depend on abstractions.

---

## How This Follows the Dependency Inversion Principle

### High-Level Module (`OrderService`)
- Depends only on the `Notifier` abstraction, not on any concrete notifier (**Email**, **SMS**, **Push**).

### Low-Level Modules (`EmailNotifier`, `SMSNotifier`, `PushNotifier`)
- Implement the `Notifier` interface, depending on the abstraction.

### Abstraction Independence
- The `Notifier` interface sits between high- and low-level code.
- Neither side depends on the other’s details — formatting or delivery logic changes never touch `OrderService`, and business logic changes never touch the concrete notifiers.

### Extensibility
- Adding a new notification channel (e.g., `SlackNotifier`, `WebhookNotifier`) involves only creating a new struct that implements `Notifier` and injecting it into `OrderService`.
- No modifications to existing business-logic code, ensuring the system is **open for extension** but **closed for modification**.

---
