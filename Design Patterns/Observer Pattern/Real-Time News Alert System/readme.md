# 📋 Problem Statement

Google’s Real-Time News Alert System requires a scalable solution to notify multiple subscribers (users and services) whenever breaking news is published. The system must efficiently handle high volumes of alerts and distribute notifications concurrently for improved performance.

The system should notify multiple observers such as:

- **Mobile App Notification Service** (Push notifications for app users)  
- **Email Notification Service** (Emails for subscribers)  
- **News Aggregator Service** (3rd-party platforms that pull news data)  

The system should:

✅ Automatically notify all observers when a breaking news alert is published.  
✅ Allow adding new observer services with minimal code changes.  
✅ Utilize **Goroutines** to efficiently deliver notifications to multiple observers concurrently.  
✅ Ensure **thread safety** and synchronization when notifying observers.  

---

## ❓ Question

Design a **Real-Time News Alert System** using the **Observer Design Pattern** that efficiently broadcasts breaking news alerts to multiple subscribers using Goroutines. The system should:

✅ Use the **Observer Pattern** to manage multiple subscriber services.  
✅ Implement a **common interface** for all observers.  
✅ Utilize **Goroutines** to notify observers concurrently.  
✅ Use **sync.WaitGroup** to manage concurrent Goroutines.  
✅ Ensure that adding new observers requires minimal code changes.  

---

## 🧩 Solution Overview

The solution will utilize:

✅ **Observer Design Pattern** to notify multiple observers dynamically.  
✅ A **Subject Class (News Publisher)** that maintains and notifies a list of observers.  
✅ Multiple **Observer Classes** that implement the `Update()` method for notification delivery.  
✅ **Goroutines** for concurrent observer notification to improve performance.  
✅ **sync.WaitGroup** to synchronize Goroutines and ensure proper completion.  

---

## 📚 Key Takeaways

✅ Uses the **Observer Design Pattern** to notify multiple subscribers automatically when breaking news alerts are published.  
✅ Ensures **loose coupling** between the **News Publisher** and **Subscriber Services**, improving flexibility and scalability.  
✅ Uses **Goroutines** for concurrent notification delivery, improving performance for high-traffic systems.  
✅ Ensures **thread safety** with **sync.WaitGroup** to manage concurrent Goroutines.  
✅ New services can be integrated seamlessly without modifying the existing publisher logic.  

---

## 💬 Bonus Challenge

🔹 Add an **SMS Notification Service** to send text messages for urgent breaking news.  
🔹 Implement **filter-based subscriptions** so users only receive news for specific topics.  
🔹 Introduce a **priority system** to ensure critical alerts are sent first.  
🔹 Add **logging support** to track successful and failed notifications for improved traceability.  
