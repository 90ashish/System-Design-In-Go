# 📋 Problem Statement

PayPal’s **Payment Gateway System** requires a robust solution to handle **service failures** gracefully. The system interacts with multiple third-party services such as:

- **Bank API** (Handles fund transfers)  
- **Currency Exchange Service** (Converts foreign currency rates)  
- **Transaction Logging Service** (Stores payment history securely)  

Due to unpredictable network conditions or system downtime, these external services may become **unavailable** or **slow**. The payment gateway must be designed to:

✅ Automatically **stop requests** to failing services to prevent system overload.  
✅ **Monitor system health** and automatically resume requests once the service is stable.  
✅ Ensure that **failed transactions** are gracefully handled without affecting the core payment flow.  

---

## ❓ Question

Design a **Payment Gateway System** using the **Circuit Breaker Pattern** that handles service failures gracefully by:

✅ Implementing a **Circuit Breaker Pattern** to protect the system from repetitive failed requests.  
✅ Introducing **Open**, **Closed**, and **Half-Open** circuit states to control retry logic.  
✅ Implementing a **Timeout Mechanism** for slow services.  
✅ Ensuring the system automatically **recovers** when the failing service becomes available again.  
✅ Using **Goroutines** to simulate concurrent payment transactions.  

---

## 🧩 Solution Overview

The solution will utilize:

✅ **Circuit Breaker Pattern** to control request flow and handle service failures gracefully.  
✅ Three **Circuit States**:  
- **Closed State:** Service is healthy; requests are processed normally.  
- **Open State:** Service is unhealthy; all requests are blocked for a defined cooldown period.  
- **Half-Open State:** Service is in recovery mode; a limited number of test requests are sent to check stability.  

✅ A **Timeout Mechanism** to limit the waiting period for slow services.  
✅ **Goroutines** for concurrent request simulation and improved performance.  
✅ **sync.Mutex** to ensure **thread-safe circuit state transitions**.  

---

## 📚 Key Takeaways

✅ Uses the **Circuit Breaker Pattern** to prevent system overload from failing services.  
✅ Implements **Closed**, **Open**, and **Half-Open** states to manage retries intelligently.  
✅ Utilizes **Goroutines** for concurrent transactions, improving system throughput and scalability.  
✅ Ensures **thread safety** using **sync.Mutex** to manage the circuit breaker state changes.  
✅ Provides a **timeout mechanism** to retry failed services after a defined period.  

---

## 💬 Bonus Challenge

🔹 Add **Logging Support** to track circuit state changes and request patterns for better monitoring.  
🔹 Implement **Exponential Backoff Logic** for smarter retry intervals in the Half-Open state.  
🔹 Introduce a **Health Check Mechanism** to proactively monitor external service health.  
🔹 Add **Rate Limiting** for additional protection during high traffic surges.  
