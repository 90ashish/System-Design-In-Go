# 📋 Problem Statement

Google’s **Search Aggregation System** requires a solution to ensure **high availability** and **resilience** when interacting with third-party APIs or unreliable services.

The system should:

- Continuously **fetch data** from multiple external services.  
- Implement a **failure detection mechanism** that stops sending requests to a failing service for a certain period to prevent system overload.  
- Automatically **reset and attempt reconnection** after a cooldown period.  

The goal is to design a solution that:

✅ Ensures **system stability** when dependent services are down.  
✅ Prevents **cascading failures** by blocking requests to failing services.  
✅ Implements **auto-recovery** after the failing service is restored.  

---

## ❓ Question

Design a **Search Aggregation System** using the **Circuit Breaker Pattern** that handles multiple third-party services with resilience and ensures stability during failures. The system should:

✅ Use the **Circuit Breaker Pattern** to monitor service health.  
✅ Transition between **Closed**, **Open**, and **Half-Open** states.  
✅ Automatically **block requests** to failing services for a cooldown period.  
✅ Implement **auto-recovery** when the failing service is restored.  
✅ Ensure **thread-safety** when multiple Goroutines make concurrent requests.  

---

## 🧩 Solution Overview

The solution will utilize:

✅ **Circuit Breaker Pattern** to manage the reliability of multiple external services.  
✅ A **State Management System** with the following states:

- **Closed State:** Service is healthy, and requests are sent normally.  
- **Open State:** Service is unhealthy; all requests are blocked for a cooldown period.  
- **Half-Open State:** Service is in recovery mode; a limited number of requests are tested to check the service’s stability.  

✅ **Goroutines** to handle multiple concurrent requests efficiently.  
✅ **sync.Mutex** to ensure **thread-safety** while modifying the circuit breaker state.  

---

## 📚 Key Takeaways

✅ Uses the **Circuit Breaker Pattern** to manage service failures and ensure system stability.  
✅ Ensures the system dynamically transitions between **Closed**, **Open**, and **Half-Open** states.  
✅ Utilizes **Goroutines** to handle multiple concurrent service calls efficiently.  
✅ Implements **sync.Mutex** to ensure thread safety when modifying Circuit Breaker states.  
✅ Prevents **cascading failures** by blocking requests to unreliable services for a defined cooldown period.  

---

## 💬 Bonus Challenge

🔹 Add **metrics tracking** to monitor request success and failure rates.  
🔹 Implement a **notification system** to alert admins when the circuit breaker is activated.  
🔹 Introduce **rate limiting** to control request throughput and avoid server overload.  
🔹 Add **adaptive recovery logic** to dynamically adjust cooldown periods based on service recovery patterns.  
