# Interview Question: Payment Gateway Integration Using OCP

## Question

Design a Payment Gateway Integration System that supports multiple providers (e.g., PayPal, Stripe, etc.) using the **Open/Closed Principle (OCP)**. The system should allow adding new payment providers without modifying the existing core logic. Explain your design and provide a sample implementation in Golang.

---

## Explanation

### Abstraction with Interfaces

- The `PaymentProvider` interface defines a contract with the method `ProcessPayment`.  
- Both `PayPalProvider` and `StripeProvider` implement this interface.  
- This allows the `PaymentGateway` to work with any provider that satisfies the interface.

### Separation of Concerns

- The `PaymentGateway` struct is responsible for orchestrating the payment process.  
- It doesn't need to know the details of each payment provider, thus staying closed for modifications but open for extension by injecting different implementations.

### Extensibility

- Adding a new payment provider (e.g., `SquareProvider`) only requires implementing the `PaymentProvider` interface.  
- The existing `PaymentGateway` code remains unchanged, adhering to the **Open/Closed Principle**.

### Usage in Main

- The `main()` function demonstrates how to process a payment using both PayPal and Stripe providers.  
- This design ensures that new providers can be integrated seamlessly.

---

This implementation effectively demonstrates how to design an extensible payment integration system using the **Open/Closed Principle** in Golang.
