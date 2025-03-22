# 📋 Problem Statement

Google’s **File Encryption System** requires a flexible solution to provide multiple layers of encryption and compression for securing user data before storage or transmission.

The system should support multiple functionalities such as:

- **Basic File Saving**  
- **Encryption Layer** (Encrypts data before saving)  
- **Compression Layer** (Compresses data to reduce file size)  
- **Checksum Layer** (Adds a checksum to verify file integrity)  

The system should allow users to dynamically apply one or more layers of protection **without modifying the core file-saving logic**.

---

## ❓ Question

Design a **File Handling System** using the **Decorator Design Pattern** that allows users to dynamically add layers of encryption, compression, and checksum verification when saving files. The system should:

✅ Use the **Decorator Pattern** to dynamically add layers of functionality.  
✅ Implement a **common interface** for all file handlers to ensure consistency.  
✅ Allow adding new features with minimal code changes.  
✅ Ensure the system is flexible enough for users to apply **one or multiple layers** as needed.  

---

## 🧩 Solution Overview

The solution will utilize:

✅ **Decorator Design Pattern** to wrap file-saving functionality with additional layers like encryption, compression, and checksum.  
✅ An **Interface-Based Design** to enforce a common method signature for all file handlers.  
✅ Each **Decorator** will wrap the core file-saving logic and provide enhanced functionality.  
✅ Ensures **scalability** by easily integrating additional layers of functionality without modifying the core code.  

---

## 📚 Key Takeaways

✅ Uses the **Decorator Design Pattern** to dynamically apply layers of encryption, compression, and checksum verification.  
✅ Ensures **flexibility** — users can choose **one or multiple layers** of protection as needed.  
✅ Implements the **Open/Closed Principle** — adding new features only requires creating a new decorator without modifying the core logic.  
✅ Demonstrates **polymorphism** by invoking the `.Save()` method across multiple decorated objects.  
✅ Provides a **scalable** and **maintainable** solution for enhancing file security and data integrity.  

---

## 💬 Bonus Challenge

🔹 Add a **BackupDecorator** that automatically backs up the file after saving.  
🔹 Introduce a **VersionControlDecorator** that tracks file versions during updates.  
🔹 Implement a **LoggingDecorator** that logs file-saving actions for improved traceability.  
🔹 Enhance the solution with **Error Recovery Logic** for failed encryption or compression tasks.  
