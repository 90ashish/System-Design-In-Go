## 📋 Problem Statement

Amazon's Shipping System requires a flexible solution to calculate shipping costs based on different shipping strategies. The platform offers various delivery options to optimize shipping time and costs. The available shipping methods include:

- **Standard Shipping** (5-7 days delivery)  
- **Express Shipping** (2-3 days delivery)  
- **Same-Day Delivery** (Delivery within 24 hours)  

### System Requirements
The system should be designed to:

✅ Dynamically select the appropriate shipping strategy based on customer preference.  
✅ Ensure new shipping methods can be added without modifying the existing logic.  
✅ Improve code maintainability and ensure the system follows the **Open/Closed Principle**.  

---

## ❓ Question

Design a **Shipping Cost Calculator** using the **Strategy Design Pattern** that dynamically selects the correct shipping method based on customer preference. The system should:

✅ Use the **Strategy Pattern** to switch between different shipping strategies dynamically.  
✅ Implement a **common interface** for all shipping strategies.  
✅ Allow adding new shipping strategies with minimal code changes.  
✅ Ensure each shipping strategy calculates the shipping cost independently.  
✅ Provide an interface that allows users to calculate the total shipping cost by simply selecting a method.  

---

## 🧩 Solution Overview

The solution will utilize:

✅ **Strategy Design Pattern** to define multiple shipping strategies dynamically.  
✅ An **Interface-Based Design** to enforce a common method signature for all shipping strategies.  
✅ A **Context Class** that applies the chosen strategy for shipping cost calculation.  
✅ Ensures **scalability** by easily integrating additional shipping strategies without modifying existing code.  

---

## 📚 Key Takeaways

✅ Uses the **Strategy Design Pattern** to dynamically switch between different shipping strategies.  
✅ Ensures **code maintainability** — new shipping strategies can be added with minimal code changes.  
✅ Demonstrates **polymorphism** by invoking the correct shipping strategy dynamically at runtime.  
✅ Ensures adherence to the **Open/Closed Principle** — existing code remains untouched when adding new strategies.  
✅ Improves **scalability** by providing a flexible architecture for integrating future shipping methods.  

---

## 💬 Bonus Challenge

🔹 **Add International Shipping** with dynamic currency conversion rates.  
🔹 Introduce **Discount Coupons** that can modify the final shipping cost.  
🔹 Implement **Shipping Insurance** as an optional feature to enhance customer protection.  
🔹 Add **Real-Time Tracking Integration** for improved visibility during order delivery.  
