# Designing a Vending Machine

## Requirements
- The vending machine should support multiple products with different prices and quantities.
- The machine should accept coins and notes of different denominations.
- The machine should dispense the selected product and return change if necessary.
- The machine should keep track of the available products and their quantities.
- The machine should handle multiple transactions concurrently and ensure data consistency.
- The machine should provide an interface for restocking products and collecting money.
- The machine should handle exceptional scenarios, such as insufficient funds or out-of-stock products.

## Classes, Interfaces and Enumerations
- **Product class**: Represents a product in the vending machine, with properties such as name and price.
- **Coin and Note enums**: Represent the different denominations of coins and notes accepted by the vending machine.
- **Inventory class**: Manages the available products and their quantities in the vending machine. It uses a concurrent hash map to ensure thread safety.
- **VendingMachineState interface**: Defines the behavior of the vending machine in different states, such as idle, ready, and dispense.
- **IdleState, ReadyState, and DispenseState classes**: Implement the VendingMachineState interface and define the specific behaviors for each state.
- **VendingMachine class**: The main class that represents the vending machine. It follows the Singleton pattern to ensure only one instance of the vending machine exists.
- **VendingMachine class**: Maintains the current state, selected product, total payment, and provides methods for state transitions and payment handling.
- **VendingMachineDemo class**: Demonstrates the usage of the vending machine by adding products to the inventory, selecting products, inserting coins and notes, dispensing products, and returning change.

---

# 📚 Design Patterns used in the Vending Machine Project

| Design Pattern | Where it is used | Why it is used |
| --- | --- | --- |
| State Pattern | IdleState, ReadyState, DispenseState, ReturnChangeState | To handle the machine's behavior changing based on current internal state |
| Strategy Pattern | VendingMachineState interface | Each State (Idle/Ready/Dispense/ReturnChange) defines different behavior for same actions (InsertCoin, SelectProduct, etc.) |
| Factory Pattern (Simple/Implicit) | NewProduct, NewInventory, NewVendingMachine | To create instances of Product, Inventory, and VendingMachine cleanly, hiding object creation complexity |
| Singleton Pattern (Semi-conceptually) | (Not Strict Singleton) but one VendingMachine instance created and reused in main | Although not strictly enforced, usually vending machines are single in real world, one object reused |
| Command Pattern (very lightly) | Methods like InsertCoin, SelectProduct act like commands being queued by user action | Each method represents a command that operates differently based on machine's current state |

---

# 🔥 Quick explanation of each

## 1. State Pattern (main design pattern)
- IdleState, ReadyState, DispenseState, ReturnChangeState all implement VendingMachineState interface.
- Based on current state, machine behaves differently for:
  - InsertCoin
  - SelectProduct
  - DispenseProduct
  - ReturnChange
- **Key feature**: `SetState()` method inside VendingMachine changes the behavior dynamically.

👉 Without State Pattern, your vending machine would have become a huge if-else ladder!

## 2. Strategy Pattern
- Each State acts like a strategy to perform an action.
- For example:
  - In IdleState, InsertCoin → “Please select a product first.”
  - In ReadyState, InsertCoin → adds money.
- You can swap strategies (i.e., swap states) without modifying VendingMachine logic.

## 3. Factory Pattern
- Functions like:
    ```go
    func NewVendingMachine() *VendingMachine
    func NewProduct(name string, price float64) *Product
    func NewInventory() *Inventory
    ```
- These act as simple factories to encapsulate object creation.

## 4. Singleton Pattern (light)
- In main, you create one VendingMachine instance, and keep operating on it.
- But strict Singleton enforcement (like private constructor + global instance) is not done here (which is fine for now).

## 5. Command Pattern (light, behavioral)
- Every user action (`SelectProduct()`, `InsertCoin()`, `DispenseProduct()`) is modeled as a command.
- Based on current state, commands behave differently.

---

# 🏆 So final answer:

The main design patterns used are:
- State Pattern (primary)
- Strategy Pattern (secondary)
- Simple Factory
- Light Singleton usage
- Light Command pattern

---

# ✨ In Interviews — How to answer if asked:

✅ "This design mainly uses the **State Pattern** to handle the complex dynamic behavior depending on internal state transitions of the machine.  
✅ Each state implements a **Strategy** through the `VendingMachineState` interface.  
✅ Simple **Factory Methods** like `NewProduct` hide object creation complexity.  
✅ The vending machine is modeled as a single instance (conceptually **Singleton**).  
✅ User operations (InsertCoin, SelectProduct) are loosely inspired by the **Command Pattern**."
