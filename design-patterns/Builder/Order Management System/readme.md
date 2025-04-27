# 🚀 Order Management System Using Builder Pattern

## 📋 Problem Statement
Amazon's **Order Management System** requires a solution to construct complex **order objects** with multiple optional configurations. Each order can have various attributes such as:

- **Shipping Address**  
- **Billing Address**  
- **Gift Wrapping**  
- **Discount Coupons**  
- **Order Notes**  

Since the combinations of these configurations may vary significantly across orders, constructing such objects using a **Builder Pattern** is ideal to simplify object creation and improve maintainability.

---

## ❓ Question
> Design an **Order Management System** using the **Builder Design Pattern** that can dynamically construct order objects with various configurations. The system should:

✅ Use the **Builder Pattern** to construct complex order objects step by step.  
✅ Support **optional configurations** such as gift wrapping, order notes, and discount codes.  
✅ Ensure **immutable order objects** to maintain data integrity.  
✅ Allow **chaining methods** to improve readability and enhance usability.  

---

## 🧩 Solution Overview
The solution will utilize:

✅ **Builder Design Pattern** to construct complex order objects step by step.  
✅ A **Director Class** to manage the order-building process and enforce a clear sequence of steps.  
✅ **Method Chaining** to enhance code readability and usability.  
✅ Ensures **immutability** by making the constructed order object immutable.  

---

## 📚 Key Takeaways
✅ Uses the **Builder Design Pattern** to construct complex order objects step by step.  
✅ Ensures **optional configurations** are flexible and easy to add without modifying existing code.  
✅ Supports **method chaining** for improved readability and user-friendliness.  
✅ Ensures **immutability** by constructing complete order objects only once via the `.Build()` method.  
✅ Demonstrates how the **Builder Pattern** effectively handles objects with numerous optional fields.  

---

## 💬 Bonus Challenge
🔹 Add support for **multiple product items** with individual details like quantity and price.  
🔹 Implement **validation logic** to ensure required fields (e.g., shipping address) are mandatory.  
🔹 Introduce **dynamic pricing rules** for discounts, taxes, and promotions.  
🔹 Implement a **tracking system** to provide real-time updates on order status.  

