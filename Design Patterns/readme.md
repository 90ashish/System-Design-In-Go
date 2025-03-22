# 📚 Design Patterns in Golang

Welcome to the **Design Patterns in Golang** repository! 🚀  
This repository provides comprehensive solutions and practical examples for implementing key **Design Patterns** using **Golang**. Each pattern includes clear explanations, sample code, and practical use cases to help you understand and apply these design principles effectively.

---

## 🔍 Design Patterns Covered

This repository showcases the following key design patterns in **Golang**:

### 1. 🟦 Singleton Pattern
- Ensures that a class has **only one instance** and provides a **global point of access** to it.  
- Ideal for configurations, database connections, or managing shared resources.  

### 2. 🟨 Factory Method Pattern
- Provides an interface for creating objects in a **superclass**, but allows subclasses to **alter the type of objects** that will be created.  
- Useful for creating families of related objects without specifying their exact class.  

### 3. 🟧 Builder Pattern
- A step-by-step creation pattern that constructs complex objects by **separating construction logic from representation**.  
- Suitable for constructing objects with numerous optional parameters.  

### 4. 🟪 Strategy Pattern
- Allows selecting a **behavior or algorithm** at runtime by defining a family of algorithms and making them interchangeable.  
- Effective in scenarios requiring multiple algorithms for similar tasks.  

### 5. 🟩 Observer Pattern
- Establishes a **one-to-many dependency** between objects, ensuring that when one object changes state, all its dependents are automatically notified.  
- Commonly used in event-driven systems and notifications.  

### 6. 🟫 Decorator Pattern
- Allows behavior to be added to individual objects without modifying their class by **wrapping objects with new functionality**.  
- Useful for adding flexible, dynamic features to objects.  

---

## 🏆 Advanced System Design Patterns

In addition to core design patterns, this repository also covers advanced **System Design Patterns** essential for building scalable and resilient distributed systems:

### 7. ✅ Circuit Breaker Pattern
- Provides **failure protection** by detecting failures and preventing repeated requests to a failing service.  
- Improves system stability and prevents cascading failures in microservices.  

### 8. ✅ Strangler Pattern
- Facilitates **safe migration** from monolithic architectures to microservices by gradually replacing parts of the old system with new services.  
- Ensures minimal disruption during refactoring.  

### 9. ✅ Saga Pattern
- Manages **distributed transactions** by coordinating multiple microservices through a sequence of compensating transactions.  
- Ideal for ensuring data consistency across distributed systems.  

### 10. ✅ Bulkhead Pattern
- Isolates system components into **independent compartments**, ensuring that failures in one component do not affect others.  
- Enhances system reliability by limiting the impact of failures.  

---

## 🚀 Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/90ashish/System-Design-In-Go.git
   cd Design\ Patterns 
   ```

2. Explore individual patterns in their respective folders for detailed explanations and example code.

3. Run sample code using:

```
go run main.go
```