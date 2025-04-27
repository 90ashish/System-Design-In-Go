📋 Problem Statement

Google’s Order Fulfillment System for its Google Shopping Platform requires a reliable way to manage complex, long-running transactions across multiple distributed services.

For example, an order placement workflow involves the following steps:

    Payment Service → Process the payment.

    Inventory Service → Reserve the items in stock.

    Shipping Service → Arrange the shipment.

    Notification Service → Notify the customer about the order status.

If any step fails, the system should trigger compensating transactions to roll back the previously completed steps to maintain data consistency.
❓ Question

    Design an Order Fulfillment System using the Saga Pattern that ensures:

✅ Reliable handling of long-running transactions across multiple services.
✅ Ability to trigger compensating transactions in case of failures.
✅ Ensures eventual consistency across services.
✅ Efficient handling of concurrent transactions.
✅ Implement a rollback mechanism to maintain system integrity.
🧩 Solution Overview

The solution will utilize:

✅ Saga Pattern with a Choreography-based approach for distributed coordination.
✅ Each service will publish events that trigger subsequent steps or compensating actions.
✅ Services will communicate asynchronously to ensure loose coupling and scalability.
✅ Ensures eventual consistency by reversing previously successful steps when a failure occurs.

📚 Key Takeaways

✅ Utilizes the Saga Pattern with a Choreography-based approach for distributed transactions.
✅ Ensures eventual consistency across multiple services.
✅ Uses compensating transactions to rollback previous successful steps during failures.
✅ Implements a scalable, asynchronous design by allowing each service to work independently.
✅ Ideal for complex e-commerce workflows, booking systems, or financial transactions.

💬 Bonus Challenge

🔹 Implement timeout logic to automatically rollback if a service doesn’t respond in time.
🔹 Add logging and monitoring to improve traceability of success and rollback events.
🔹 Introduce a retry mechanism for transient failures (e.g., network issues).
🔹 Extend the system to track partial order fulfillment for in-stock items.