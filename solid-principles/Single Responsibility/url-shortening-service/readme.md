# 📋 Google System Design Interview Question: Single Responsibility Principle (SRP)

## ❓ Question

**Design a URL Shortening Service** similar to **bit.ly** that follows the **Single Responsibility Principle (SRP)**. The service should:

✅ Encode long URLs into **short URLs**.  
✅ Decode short URLs back to their **original URLs**.  
✅ Track **analytics** such as the number of visits for each URL.  
✅ Ensure that **URL encoding**, **decoding**, and **analytics tracking** are **separate concerns**.  

---

## 🧩 Solution in Golang Implementing SRP

### Key Components

1. **URLShortener**  
   - Responsible for **encoding** and **decoding** URLs.  

2. **AnalyticsService**  
   - Tracks **visit counts** for each URL.  

3. **StorageService**  
   - Handles **URL storage** and **retrieval**.  

---

## 🔎 Explanation

✅ **Single Responsibility Principle (SRP):**  
- **StorageService** handles **data storage** and **retrieval**, ensuring data management is isolated.  
- **URLShortener** is dedicated solely to **encoding** and **decoding** URLs.  
- **AnalyticsService** manages only **visit tracking** without interfering with other logic.  

✅ **Scalability:**  
- Each component can be **extended** or **modified** independently.  
- Adding new features like **URL expiry**, **detailed analytics**, or **multiple storage backends** can be implemented with minimal code changes.  

✅ **Dependency Injection:**  
- Dependencies are injected through **constructors**, promoting **loose coupling** and better testability.  

---

## 📌 Potential Follow-Up Questions Google May Ask

1. **How would you extend the system to support URL expiry?**  
   - Introduce a **TTL (Time-to-Live)** feature in the `StorageService` to automatically delete expired URLs.  

2. **How would you design this system for high availability and scalability?**  
   - Implement **load balancing**, **distributed storage** (e.g., Redis, DynamoDB), and **CDN caching** to enhance scalability.  

3. **What changes would you make to ensure data persistence in a distributed environment?**  
   - Use a **replicated database** or **eventual consistency strategies** to maintain data integrity.  

4. **How can caching strategies improve the performance of the URL resolution process?**  
   - Implement **in-memory caching** with **Redis** or **Memcached** to reduce database hits and improve lookup performance.  

---