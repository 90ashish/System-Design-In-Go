# 🚀 Payment Processing System Using Factory Pattern with Goroutines

## 📋 Problem Statement
Imagine you are tasked with developing a **Payment Processing System** for Amazon's e-commerce platform. The platform supports multiple payment methods such as:

- **Credit Card**  
- **PayPal**  
- **Amazon Pay**  

Each payment method has its own unique way of processing transactions, but they should all adhere to a **common interface** to ensure scalability and maintainability.

In addition, the system should efficiently handle **concurrent transactions** to improve performance during peak traffic.

---

## ❓ Question
> Design a **Payment Processing System** using the **Factory Design Pattern** with **Goroutines** to handle concurrent transactions. The system should:

✅ Use the **Factory Pattern** to instantiate the correct payment method.  
✅ Implement a **common interface** for all payment methods.  
✅ Support adding **new payment methods** with minimal code changes.  
✅ Demonstrate **polymorphism** by invoking the `ProcessPayment()` method dynamically.  
✅ Efficiently process **concurrent payment transactions** using Goroutines.  

---

## 🧩 Solution Overview
The solution will utilize:

✅ **Factory Design Pattern** to create the appropriate payment method at runtime.  
✅ **Goroutines** to ensure concurrent payment transactions are handled efficiently.  
✅ **sync.WaitGroup** to manage Goroutines and ensure all transactions are completed.  
✅ **Interface-based Design** to enforce a common method signature for all payment methods.  
✅ Ensures easy scalability by allowing new payment methods to be added with minimal code changes.  

---

## 🏗️ Solution Implementation
The solution includes:

- **Interface Implementation:** For defining the common payment method behavior.  
- **Factory Class:** For dynamically creating the appropriate payment method.  
- **Goroutines & sync.WaitGroup:** To manage concurrent payment processing.  

---

## 💬 Bonus Challenge
🔹 Add **timeout logic** for payment requests to prevent Goroutines from hanging indefinitely.  
🔹 Implement a **retry mechanism** for failed transactions.  
🔹 Introduce a **logging system** to track transaction statuses asynchronously using **channels**.  
🔹 Extend the solution to support **refund processing** or **transaction history** tracking.  
