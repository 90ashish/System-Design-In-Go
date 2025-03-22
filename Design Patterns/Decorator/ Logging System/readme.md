# 📋 Problem Statement

Microsoft's **Logging System** requires a solution that supports multiple layers of log enhancement for a scalable and efficient backend system. The system should handle **concurrent log entries** using **Goroutines** to maximize performance.

The logging system should support:

- **Basic Logging** (Writes logs directly to a file or console)  
- **Timestamp Decorator** (Adds a timestamp to each log entry)  
- **Error Level Decorator** (Indicates the severity of the log: INFO, WARNING, ERROR)  
- **File Output Decorator** (Saves log entries to a specified file)  

The system should be designed to dynamically apply one or more decorators while efficiently handling concurrent log entries.

---

## ❓ Question

Design a **Logging System** using the **Decorator Design Pattern** that dynamically adds various log enhancements like timestamps, error levels, and file output support. The system should:

✅ Use the **Decorator Pattern** to dynamically add layers of functionality.  
✅ Implement a **common interface** for all logging handlers to ensure consistency.  
✅ Allow adding new log enhancement features with minimal code changes.  
✅ Utilize **Goroutines** to efficiently handle multiple concurrent log entries.  
✅ Use **sync.WaitGroup** to manage concurrent Goroutines safely.  

---

## 🧩 Solution Overview

The solution will utilize:

✅ **Decorator Design Pattern** to add dynamic log features like timestamps, error levels, and file outputs.  
✅ An **Interface-Based Design** to ensure consistent log entry structure.  
✅ **Goroutines** to process multiple log entries concurrently, improving performance.  
✅ **sync.WaitGroup** to ensure synchronization and proper completion of concurrent Goroutines.  

---

## 📚 Key Takeaways

✅ Uses the **Decorator Design Pattern** to dynamically apply multiple log enhancements like timestamps, error levels, and file output.  
✅ Utilizes **Goroutines** for concurrent log entries, ensuring efficient log handling during peak traffic.  
✅ Ensures **scalability** — new log features can be added without modifying existing code.  
✅ Demonstrates **polymorphism** by invoking `.Log()` across multiple decorated loggers.  
✅ Ensures proper synchronization with **sync.WaitGroup** to manage concurrent Goroutines.  

---

## 💬 Bonus Challenge

🔹 Add a **JSONFormatterDecorator** to format logs in JSON format for structured logging.  
🔹 Introduce a **LogRotationDecorator** to automatically rotate logs when file size exceeds a threshold.  
🔹 Implement a **Rate Limiter** to control the number of log entries processed per second.  
🔹 Add a **SlackNotifierDecorator** to send critical error logs directly to Slack for real-time monitoring.  
