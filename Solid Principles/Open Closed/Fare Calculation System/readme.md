# Uber System Design Interview Question: Fare Calculation System Using OCP

## Question

Uber’s fare calculation involves multiple components—base fare, distance fare, time fare, surge pricing, and potentially other factors in the future. Design a fare calculation system in Golang that adheres to the **Open/Closed Principle (OCP)**. Your design should allow new pricing strategies to be added without modifying the existing code.

---

## Requirements

### Additive Strategies

- Calculate parts of the fare that add up (e.g., base fare, distance fare, time fare).

### Multiplicative Strategies

- Apply modifiers such as surge pricing that scale the base fare.

### Extensibility

- New pricing strategies (e.g., toll fees, promotional discounts) should be integrated by simply adding new implementations rather than changing the core fare calculation logic.

---

## Explanation

### Interfaces for Abstraction

- `AdditivePricingStrategy`: Each additive strategy (base fare, distance, time) implements the `Calculate` method.  
- `MultiplicativePricingStrategy`: The surge pricing strategy implements the `Multiply` method.

### Separation of Concerns

- Each strategy is responsible for one aspect of the fare calculation.  
- The `FareCalculator` aggregates these strategies without being concerned with the internal implementation details.

### Open/Closed Principle (OCP)

- **Open for Extension**: New strategies (like toll fees or discounts) can be added by implementing the appropriate interface.  
- **Closed for Modification**: The `FareCalculator` and existing strategies remain unchanged when new pricing components are introduced.

### Usage

- In the `main()` function, the fare is calculated for a sample ride.  
- The design allows new pricing strategies to be injected easily, ensuring the system is scalable and maintainable.

---
