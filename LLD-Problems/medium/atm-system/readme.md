# Designing an ATM System

## Requirements

- The ATM system should support basic operations such as balance inquiry, cash withdrawal, and cash deposit.
- Users should be able to authenticate themselves using a card and a PIN (Personal Identification Number).
- The system should interact with a bank's backend system to validate user accounts and perform transactions.
- The ATM should have a cash dispenser to dispense cash to users.
- The system should handle concurrent access and ensure data consistency.
- The ATM should have a user-friendly interface for users to interact with.


## How It Works  
1. **Initialization**  
   - `main.go` wires up components via factory methods: console UI, card reader, dispenser, and a mock bank backend.  
   - Constructs the `ATM` (Facade) by injecting these dependencies.  

2. **Authentication Loop** (`ATM.Run`)  
   - Prompts for card number & PIN.  
   - Validates against the backend (`MockBackend.Validate`).  
   - On success, enters the session; on failure, restarts.  

3. **Session Loop** (`ATM.session`)  
   - Displays a menu:  
     1) Balance Inquiry  
     2) Cash Withdrawal  
     3) Cash Deposit  
     4) Exit  
   - Executes the chosen operation and then re-prompts until exit.  

4. **Operations**  
   - **Balance Inquiry**: calls `backend.GetBalance`, displays the result.  
   - **Cash Withdrawal**: prompts amount, calls `backend.Withdraw`, then `dispenser.Dispense`.  
   - **Cash Deposit**: prompts amount, calls `backend.Deposit`.  

5. **Concurrency & Consistency**  
   - `MockBackend` and each `Account` use `sync.RWMutex` to guard shared state.  

## Running the Project  
    ```bash
    go run cmd/app/main.go
    ```

## Design Patterns Used

- **Factory Pattern**: `NewXYZ()` functions create and configure components.
- **Strategy Pattern**: `UI`, `CardReader`, `CashDispenser`, `BankBackend` are interchangeable behaviors.
- **Dependency Injection (DIP)**: All dependencies are passed into constructors (`NewATM`, `NewConsoleCardReader`, etc.).
- **Facade Pattern**: `ATM` simplifies the workflow into `Run()` and `session()`.
- **Repository Pattern**: `BankBackend` abstracts data operations.
- **Thread Safety**: `sync.RWMutex` guards shared state in `MockBackend` and `Account`.