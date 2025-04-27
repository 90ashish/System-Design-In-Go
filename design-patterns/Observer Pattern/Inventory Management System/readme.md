# 📋 Problem Statement

Amazon’s Inventory Management System requires a solution to notify multiple services when product stock levels change. The system must automatically update various components like:

- **Warehouse Service** (Manages inventory stock)  
- **User Notification Service** (Alerts customers when products are back in stock)  
- **Analytics Service** (Tracks sales trends and restock needs)  

The system should ensure that whenever a product's stock changes, all dependent services are notified immediately.

---

## ❓ Question

Design an Inventory Management System using the **Observer Design Pattern** that notifies multiple services whenever a product's stock is updated. The system should:

✅ Use the Observer Pattern to manage multiple dependent services.  
✅ Implement a common interface for all observer services.  
✅ Ensure that new observers can be added with minimal code changes.  
✅ Notify all observers automatically when product stock updates.  
✅ Maintain loose coupling between the **Subject** (Inventory) and its **Observers**.  

---

## 🧩 Solution Overview

The solution will utilize:

✅ **Observer Design Pattern** to broadcast product stock changes to multiple observers.  
✅ **Interface-Based Design** to enforce a common method signature for all observers.  
✅ A **Subject Class (Inventory)** that maintains and notifies a list of observers.  
✅ Ensures scalability by easily integrating additional observer services.  

---

## 📚 Key Takeaways

✅ Uses the **Observer Design Pattern** to automatically notify multiple services when product stock is updated.  
✅ Ensures **loose coupling** by separating the **Inventory (Subject)** from the dependent services (**Observers**).  
✅ Ensures **scalability** — new services can be added without modifying the existing Inventory logic.  
✅ Demonstrates **polymorphism** by dynamically invoking each observer's `Update()` method.  
✅ Utilizes `sync.Map` to ensure **thread-safe observer management** for concurrent updates.  

---

## 💬 Bonus Challenge

🔹 Add a **Sales Forecasting Service** to predict product demand based on inventory changes.  
🔹 Introduce **Real-Time Pricing Adjustments** based on stock levels and demand trends.  
🔹 Implement a **Low Stock Alert System** to notify warehouse managers when critical stock levels are reached.  
🔹 Add **Logging Support** to track inventory changes and observer notifications for enhanced traceability.  
