# 🚀 Notification System Using Factory Pattern with Goroutines

## 📋 Problem Statement
Amazon's warehouse management system requires a solution to handle different types of **notification services** to inform customers about their order status. The system should support multiple notification channels such as:

- **Email Notification**  
- **SMS Notification**  
- **Push Notification**  

Each notification type must follow a **common interface** to ensure consistency. The goal is to create a **flexible** and **scalable** system that can easily integrate **new notification methods** without modifying existing logic.

---

## ❓ Question
> Design a **Notification System** using the **Factory Design Pattern** that creates the correct notification object dynamically based on the user’s preference. The system should:

✅ Use the **Factory Pattern** to instantiate the correct notification method.  
✅ Implement a **common interface** for all notification methods.  
✅ Allow adding **new notification types** with minimal code changes.  
✅ Demonstrate **polymorphism** by dynamically invoking the `SendNotification()` method.  
✅ Utilize **Goroutines** to efficiently handle **bulk notifications** concurrently.  

---

## 🧩 Solution Overview
The solution will utilize:

✅ **Factory Design Pattern** to create the appropriate notification object dynamically.  
✅ **Interface-based Design** to enforce a common method signature for all notification types.  
✅ **Goroutines** to efficiently send bulk notifications concurrently.  
✅ **sync.WaitGroup** to manage concurrent Goroutines and ensure all notifications are completed.  

---

## 📚 Key Takeaways
✅ Uses the **Factory Design Pattern** to create different notification methods dynamically.  
✅ Utilizes **Goroutines** for concurrent notification dispatch, improving system performance.  
✅ Ensures **scalability** — new notification types can be easily integrated by adding a new struct and registering it in the Factory.  
✅ Demonstrates **polymorphism** by dynamically invoking the `SendNotification()` method.  

---

## 💬 Bonus Challenge
🔹 Implement **priority-based notifications** to ensure urgent messages are sent first.  
🔹 Add **logging functionality** to track successful and failed notifications.  
🔹 Introduce a **retry mechanism** for failed notifications.  
🔹 Implement a **rate limiter** to control the number of notifications sent per second to avoid spamming users.  
