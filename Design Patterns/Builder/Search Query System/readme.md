# 🚀 Concurrent Search Query System Using Builder Pattern with Goroutines

## 📋 Problem Statement
Google’s **Search Engine Platform** requires a solution to build complex **Search Query Objects** that can support multiple optional parameters for flexible search functionality. The system should allow users to filter search results based on:

- **Keywords**  
- **Location**  
- **Date Range (Start Date and End Date)**  
- **File Type (PDF, DOCX, etc.)**  
- **Sort Order (Ascending or Descending)**  
- **Maximum Results Limit**  

Additionally, since search platforms handle **high traffic** with multiple requests being triggered simultaneously, the system should efficiently handle **concurrent search queries** to improve performance and reduce latency.

---

## ❓ Question
> Design a **Search Query System** using the **Builder Design Pattern** that efficiently constructs complex search query objects with various optional parameters. Additionally, implement **Goroutines** to handle concurrent search queries efficiently. The system should:

✅ Use the **Builder Pattern** to construct search query objects step by step.  
✅ Support **optional configurations** such as date range, file type, and sort order.  
✅ Ensure **immutable search query objects** to maintain data integrity.  
✅ Allow **method chaining** for improved readability and enhanced usability.  
✅ Introduce **Goroutines** to process **multiple search queries concurrently**.  
✅ Utilize **sync.WaitGroup** to ensure synchronization and proper completion of all Goroutines.  

---

## 🧩 Solution Overview
The solution will utilize:

✅ **Builder Design Pattern** to construct complex search query objects step by step.  
✅ **Interface-based Design** to enforce a common method signature for all search queries.  
✅ **Goroutines** to efficiently manage concurrent file uploads.  
✅ **sync.WaitGroup** to manage multiple Goroutines and ensure all uploads are completed before exiting.  
✅ Ensures **immutability** by creating the complete query object only once via `.Build()`.  

---

## 📚 Key Takeaways
✅ Uses the **Builder Design Pattern** to construct complex search query objects dynamically.  
✅ Utilizes **Goroutines** for concurrent query execution, improving system performance and reducing latency.  
✅ Ensures **scalability** — new search filters or query parameters can be easily added by defining additional methods in the **Builder Class**.  
✅ Demonstrates **polymorphism** by dynamically invoking the `DisplayQuery()` method for different types of search configurations.  

---

## 💬 Bonus Challenge
🔹 Add **timeouts** to Goroutines to prevent endless wait times in case of unresponsive systems.  
🔹 Implement **caching mechanisms** to speed up repeated queries.  
🔹 Introduce **priority-based query execution** to ensure urgent queries are processed first.  
🔹 Enhance the solution with **load balancing logic** to distribute search requests efficiently across multiple servers.  
