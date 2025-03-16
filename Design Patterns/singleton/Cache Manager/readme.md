# 🚀 Singleton-Based Cache Manager with LRU Eviction

## 📋 Problem Statement
Imagine you're developing a **Cache Manager** for a large-scale distributed system. The cache should adhere to the following constraints:

✅ Only **one instance** of the cache should exist to maintain consistent data across multiple application instances.  
✅ The cache should support **thread-safe operations** to ensure data integrity during concurrent updates.  
✅ The cache should implement **Least Recently Used (LRU)** eviction policy to manage memory efficiently.  

---

## ❓ Question
> Design a **Cache Manager** using the **Singleton Design Pattern** that ensures only one instance is created. The Cache Manager should:

✅ Allow **adding** and **retrieving** key-value pairs.  
✅ Implement **LRU eviction** when the cache exceeds its capacity.  
✅ Ensure **thread-safety** for concurrent access.  
