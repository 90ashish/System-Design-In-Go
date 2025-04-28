# 🚀 Database Connection Pool Manager Using Singleton Pattern

## 📋 Problem Statement
Imagine you're building a **Database Connection Pool Manager** for a high-traffic e-commerce platform. The platform's backend requires efficient management of database connections to ensure optimal performance. The following constraints must be met:

✅ Only **one instance** of the Connection Pool Manager should exist to avoid redundant resource allocation.  
✅ The Connection Pool Manager should support **thread-safe operations** since multiple Goroutines will request database connections simultaneously.  
✅ The Connection Pool Manager should limit the **maximum number of concurrent database connections** for performance optimization.  

---

## ❓ Question
> Design a **Database Connection Pool Manager** using the **Singleton Design Pattern** that ensures only one instance is created. The Pool Manager should:

✅ Allow **requesting** and **releasing** database connections.  
✅ Track the **number of active connections**.  
✅ Enforce a **maximum limit** on concurrent connections.  
✅ Ensure **thread-safety** for concurrent access.  

---

## 🧩 Solution Overview
The solution will utilize:

✅ **Singleton Design Pattern** to ensure only one instance of the Connection Pool Manager.  
✅ **sync.Mutex** to ensure **thread-safe** management of the pool.  
✅ A **channel-based connection pool** to manage the **maximum number of concurrent database connections efficiently**.  

---

## 📚 Key Takeaways
✅ Ensures **only one instance** of the Connection Pool Manager using the **Singleton Pattern**.  
✅ Uses **`sync.Mutex`** to ensure **thread-safety** during concurrent database connections.  
✅ Implements a **channel-based pool** for efficient management of connection limits.  
✅ Ideal for managing **database connections**, **HTTP client pools**, or **network socket pools** in high-performance applications.  

---

## 💬 Bonus Challenge
🔹 Enhance the code to include **connection timeout** logic for improved reliability.  
🔹 Implement **priority-based** connection handling for high-priority requests.  
