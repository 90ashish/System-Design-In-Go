## 📋 Problem Statement

Google's Content Delivery System requires a flexible solution to efficiently deliver media content using different content compression strategies. To enhance performance and scalability, the system should also support **concurrent compression tasks using Goroutines**.

### Available Compression Strategies
- **ZIP Compression** (Best for text-based content)  
- **GZIP Compression** (Best for web assets like CSS/JS files)  
- **LZ Compression** (Best for large media files like videos or images)  

---

## ✅ System Requirements
The system should:

✅ Dynamically select the appropriate compression strategy based on content type.  
✅ Ensure the compression system supports **concurrent compression tasks** for improved performance.  
✅ Ensure new compression algorithms can be added without modifying the existing logic.  

---

## ❓ Question

Design a **Content Compression System** using the **Strategy Design Pattern** that dynamically applies the correct compression method based on content type. The system should:

✅ Use the **Strategy Pattern** to switch between different compression algorithms dynamically.  
✅ Implement a **common interface** for all compression strategies.  
✅ Allow adding new compression strategies with minimal code changes.  
✅ Utilize **Goroutines** to efficiently handle multiple concurrent compression tasks.  
✅ Use **sync.WaitGroup** to synchronize concurrent Goroutines.  

---

## 🧩 Solution Overview

The solution will utilize:

✅ **Strategy Design Pattern** to define multiple compression strategies dynamically.  
✅ A **Context Class** that applies the chosen strategy for content compression.  
✅ **Goroutines** to concurrently handle multiple compression tasks, improving system performance.  
✅ **sync.WaitGroup** to ensure proper synchronization and prevent premature program termination.  

---

## 📚 Key Takeaways

✅ Uses the **Strategy Design Pattern** to dynamically switch between different compression strategies.  
✅ Utilizes **Goroutines** for concurrent file compression, improving system throughput and efficiency.  
✅ Ensures **scalability** — new compression strategies can be easily integrated without modifying existing logic.  
✅ Demonstrates **polymorphism** by dynamically invoking the `Compress()` method for different compression strategies.  
✅ Ensures proper synchronization with **sync.WaitGroup** to manage concurrent Goroutines.  

---

## 💬 Bonus Challenge

🔹 Add support for **encryption algorithms** as an additional security feature for compressed files.  
🔹 Implement a **progress bar system** for tracking file compression progress in real-time.  
🔹 Introduce a **priority-based compression system** to ensure urgent files are processed first.  
🔹 Enhance the solution with **batch compression support** for handling large volumes of files.  
