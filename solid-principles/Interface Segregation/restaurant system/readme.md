# Interview Question: Implementing Interface Segregation in a Restaurant System

## Question

Consider the following partial code that demonstrates the **Interface Segregation Principle (ISP)**. In this example, a monolithic interface `RestaurantEmployee1` forces a single type (e.g., `Waiter1`) to implement functions it doesn't really need. To address this, the code introduces segregated interfaces (`Waiter` and `Chef`) with corresponding concrete types (`frontdesk` and `backdesk`) that only implement the methods relevant to their roles.

### Your task is to:

- ✅ Complete the provided code so that it can be compiled and run.
- ✅ Implement a `main` function to demonstrate how each type performs its role (i.e., serving customers, taking orders, and cooking food).
- ✅ Output appropriate messages for each responsibility.

---

## Explanation

### Monolithic Interface Issue

- The `RestaurantEmployee1` interface forces `Waiter1` to implement methods like `washDishes()` and `cookFood()` even though these actions are not part of a waiter's responsibility.
- In the implementation of `Waiter1`, we return messages indicating that these methods are not its job.

### Interface Segregation Principle (ISP)

- To resolve this issue, we created **separate interfaces**:

  - `Waiter` — with methods `serveCustomers()` and `takeOrder()`
  - `Chef` — with the method `cookFood()`

- This design allows each concrete type to implement **only** the functions it needs:

  - `frontdesk` type implements the `Waiter` interface, handling serving and taking orders.
  - `backdesk` type implements the `Chef` interface, handling cooking.

---

## Demonstration in `main`

- The `main` function first shows the **limitations** of the monolithic design using `Waiter1`.
- It then demonstrates the **segregated interface design** by:

  - Creating instances of `frontdesk` and `backdesk`
  - Invoking their respective methods
  - Printing the output for each responsibility

---

## Conclusion

This implementation clearly illustrates how the **Interface Segregation Principle (ISP)** can be applied to create a **cleaner**, **more modular** design in a restaurant system. By ensuring that types only implement what they actually use, we reduce unnecessary dependencies and improve maintainability and clarity in the codebase.
