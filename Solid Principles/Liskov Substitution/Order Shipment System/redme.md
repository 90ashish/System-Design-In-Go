# Amazon System Design Interview Question: Order Shipment System Using Liskov Substitution Principle (LSP)

## Question

Amazon processes millions of orders every day, and each order can be shipped using different methods based on customer preference, location, and shipping speed. Design an order shipment system that supports multiple shipping methods—such as Standard Shipping, Expedited Shipping, and International Shipping—while adhering to the **Liskov Substitution Principle (LSP)**. Your solution should allow any shipping method (which implements the shipping abstraction) to be used interchangeably without affecting the order processing logic.

---

## Solution in Golang

The following implementation demonstrates LSP by:

### Defining a Common Abstraction

- A `Shipping` interface is created with a method `Ship(order Order) string`.

### Implementing Concrete Shipping Methods

- Concrete types such as `StandardShipping`, `ExpeditedShipping`, and `InternationalShipping` implement the `Shipping` interface.

### Order Processing using Abstraction

- The `OrderProcessor` struct accepts any type that implements the `Shipping` interface.
- This ensures that any shipping method can be substituted without altering the order processing logic.

---

## Explanation

### Common Abstraction

- The `Shipping` interface defines a contract with the method `Ship(order Order) string`.
- This abstraction enables the order processor to interact with any shipping method without being aware of its concrete implementation.

### Liskov Substitution Principle (LSP)

#### Substitutability

- All shipping methods (`StandardShipping`, `ExpeditedShipping`, and `InternationalShipping`) implement the `Shipping` interface.
- As a result, they can be used interchangeably in the `OrderProcessor`.

#### Consistent Behavior

- The `OrderProcessor` calls the `Ship` method on the `Shipping` interface without concern for which concrete type it is working with.
- Any type that implements the interface can be substituted seamlessly.

### Extensibility

- New shipping methods (e.g., `SameDayShipping` or `DroneShipping`) can be added by implementing the `Shipping` interface.
- The core order processing logic in `OrderProcessor` remains unchanged, preserving system stability while extending functionality.

---

## Conclusion

This design exemplifies how to build a **flexible** and **maintainable** system using the **Liskov Substitution Principle**, ensuring that any new shipping method can be integrated without modifying the existing order processing code.
