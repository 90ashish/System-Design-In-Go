# 🚀 Payment Processing System Using Factory Pattern

## 📋 Problem Statement
Imagine you are tasked with developing a **Payment Processing System** for Amazon's e-commerce platform. The platform supports multiple payment methods such as:

- **Credit Card**  
- **PayPal**  
- **Amazon Pay**  

Each payment method has its own unique way of processing transactions, but they should all adhere to a **common interface** to ensure scalability and maintainability.

---

## ❓ Question
> Design a **Payment Processing System** using the **Factory Design Pattern** that dynamically creates the correct payment method object based on user selection. The system should:

✅ Use the **Factory Pattern** to instantiate the correct payment method.  
✅ Implement a **common interface** for all payment methods.  
✅ Support adding **new payment methods** with minimal code changes.  
✅ Demonstrate **polymorphism** by invoking the `ProcessPayment()` method dynamically.  

---

## 🧩 Solution Overview
The solution will utilize:

✅ **Factory Design Pattern** to create the appropriate payment method at runtime.  
✅ **Interface-based Design** to enforce a common method signature for all payment methods.  
✅ Ensure **easy scalability** by allowing new payment methods to be added with minimal code changes.  

---

## 📚 Key Takeaways
✅ The **Factory Design Pattern** ensures a **flexible** and **scalable** solution for creating different payment methods.  
✅ New payment methods like **Google Pay**, **Apple Pay**, etc., can be easily added with minimal code changes.  
✅ Demonstrates **polymorphism** by invoking the `ProcessPayment()` method dynamically across different payment methods.  
✅ Enhances **code maintainability** by centralizing the creation logic in the **Factory Class**.  

---

## 💬 Bonus Challenge
🔹 Add support for **currency conversion** within each payment method.  
🔹 Implement **transaction logging** to maintain a detailed history of payments.  
🔹 Introduce an **error-retry mechanism** for failed transactions.  
