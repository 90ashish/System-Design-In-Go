# 🚀 Problem Statement

Amazon’s **Stock Notification System** requires a robust solution to notify multiple subscribers whenever stock updates occur. The system should efficiently handle a large number of subscribers and ensure notifications are delivered concurrently for improved performance.

The system must support the following subscriber types:

- **Email Subscribers** (Receive stock updates via email)  
- **SMS Subscribers** (Receive stock updates via SMS)  

The system should ensure:

✅ All subscribers are notified when a stock update occurs.  
✅ New subscriber types (like **push notifications**, **Slack integration**, etc.) can be added with minimal code changes.  
✅ Notifications should be delivered **concurrently** to improve performance for large subscriber lists.  
✅ Use **Goroutines** to handle concurrent notification delivery efficiently.  
✅ Utilize **sync.WaitGroup** to manage and synchronize multiple Goroutines.  

---

## ❓ Question

Design a **Stock Notification System** using the **Observer Design Pattern** that efficiently notifies multiple subscribers about stock updates using **Goroutines**. The system should:

✅ Use the **Observer Pattern** to manage multiple subscriber services.  
✅ Implement a **common interface** for all subscribers.  
✅ Allow adding new subscriber types with minimal code changes.  
✅ Utilize **Goroutines** to notify subscribers concurrently for improved performance.  
✅ Use **sync.WaitGroup** to ensure all Goroutines complete before the program proceeds.  

---

## 🧩 Solution Overview

The solution will utilize:

✅ **Observer Design Pattern** to manage and notify multiple subscribers.  
✅ An **Interface-Based Design** to enforce a common method signature for all subscriber types.  
✅ A **Publisher (Subject)** that stores the list of subscribers and notifies them when data is updated.  
✅ **Goroutines** to concurrently send notifications to multiple subscribers.  
✅ **sync.WaitGroup** to ensure all Goroutines complete before proceeding.  

---

## 📚 Key Takeaways

✅ Uses the **Observer Design Pattern** to ensure subscribers are automatically notified about stock updates.  
✅ Maintains **loose coupling** between the **Publisher** and **Subscribers**, improving flexibility and scalability.  
✅ Ensures concurrent delivery using **Goroutines**, improving performance for large-scale subscriber lists.  
✅ Implements **sync.WaitGroup** to manage Goroutines and ensure synchronization.  
✅ Supports seamless addition of new subscriber services without modifying the existing logic.  

---

## 💬 Bonus Challenge

🔹 Add a **Priority Queue** for urgent stock updates to ensure high-priority notifications are sent first.  
🔹 Implement a **Logging System** to track successful and failed notifications.  
🔹 Introduce a **Retry Mechanism** for failed notifications to improve reliability.  
🔹 Add **Webhook Integration** to notify external systems when stock updates occur.  
