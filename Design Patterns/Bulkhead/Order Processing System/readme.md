📋 Problem Statement

Amazon's Order Processing System needs to ensure high availability and reliability during peak sales events like Big Billion Days or Prime Day. The system handles multiple services that include:

    Payment Service for processing transactions.

    Inventory Service to manage stock updates.

    Notification Service to send order confirmations and alerts.

During peak traffic, if one of these services (e.g., the Payment Service) fails or slows down, it must not impact the overall system. The system should remain responsive to handle other services efficiently.
❓ Question

    Design an Order Processing System using the Bulkhead Pattern that:

✅ Isolates critical services to prevent system-wide failures.
✅ Ensures that failures in one service (e.g., Payment Service) don't affect others.
✅ Uses Goroutines and Buffered Channels to implement Bulkhead isolation.
✅ Implements timeouts and fallback mechanisms to improve resilience.
🧩 Solution Overview

The solution will utilize:

✅ The Bulkhead Pattern to isolate system services using separate Goroutines and buffered channels.
✅ Buffered Channels to limit the number of concurrent requests each service can handle.
✅ Timeout Mechanism to prevent long-running requests from blocking system performance.
✅ Fallback Logic to gracefully handle service failures.

📚 Key Takeaways

✅ Uses the Bulkhead Pattern to isolate services, preventing system-wide failure.
✅ Each service has a buffered channel to limit the number of concurrent requests, ensuring controlled load distribution.
✅ Ensures resilience by preventing cascading failures when one service becomes unresponsive.
✅ Uses timeouts to limit long-running tasks, enhancing system performance.
✅ Demonstrates concurrency using Goroutines and sync.WaitGroup to manage parallel requests.

💬 Bonus Challenge

🔹 Implement retry logic for failed requests to improve fault tolerance.
🔹 Add a circuit breaker mechanism to temporarily block requests to failing services.
🔹 Introduce a logging system to track and analyze rejected requests.
🔹 Add rate limiting to control the frequency of requests during peak traffic.