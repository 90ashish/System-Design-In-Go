# SOLID Principles (Simple, Beginner-Friendly Explanation)

---

## 1. **S — Single Responsibility Principle (SRP)**

🔵 **Meaning**:  
Every class, module, or function should have **only one responsibility** — **one reason to change**.

🔵 **Simple Words**:  
- One class = one job.
- If you have a "User" class, it should only manage user data — not sending emails or saving logs!

🔵 **Example**:
```go
// Bad - Doing too much
type User struct {
    Name string
}

func (u *User) SaveToDatabase() {}
func (u *User) SendEmailNotification() {}

// Good - Single Responsibility
type User struct {
    Name string
}

type UserRepository struct{}

func (r *UserRepository) Save(u User) {}

type EmailService struct{}

func (e *EmailService) SendNotification(u User) {}
```

---

## 2. **O — Open/Closed Principle (OCP)**

🟡 **Meaning**:  
Software should be **open for extension** but **closed for modification**.

🟡 **Simple Words**:  
- You can **add** new code, but you **shouldn’t change** existing working code.
- If you need a new feature, **extend** the behavior, don’t touch the already-tested part.

🔵 **Example**:
```go
// Good Example using interface
type PaymentMethod interface {
    Pay(amount float64) string
}

type CreditCard struct{}

func (c CreditCard) Pay(amount float64) string {
    return "Paid using Credit Card"
}

type PayPal struct{}

func (p PayPal) Pay(amount float64) string {
    return "Paid using PayPal"
}

// PaymentService doesn't need to change for new payment methods!
type PaymentService struct{}

func (ps PaymentService) ProcessPayment(pm PaymentMethod, amount float64) string {
    return pm.Pay(amount)
}
```

---

## 3. **L — Liskov Substitution Principle (LSP)**

🟠 **Meaning**:  
**Subclasses should behave like their parent classes** — without breaking anything.

🟠 **Simple Words**:  
- You should be able to use a child class wherever a parent class is expected — and it should **work normally**.
- No surprises, no broken logic!

✅ **Tip**:  
If an Ostrich can't fly, maybe it **shouldn't** implement a `Fly()` method. Design better interfaces!


🔵 **Example**:
```go
// Parent
type Bird interface {
    Fly()
}

// Child following LSP correctly
type Sparrow struct{}

func (s Sparrow) Fly() {
    fmt.Println("Sparrow flying!")
}

// Bad Child - violates LSP if it can't fly
type Ostrich struct{}

func (o Ostrich) Fly() {
    panic("Ostrich can't fly!") // This would surprise users!
}
```

---

## 4. **I — Interface Segregation Principle (ISP)**

🟢 **Meaning**:  
**Don't force classes to implement methods they don't need**.

🟢 **Simple Words**:  
- **Small, focused interfaces** are better.
- A class should only know about the methods that are **relevant** to it.

🔵 **Example**:
```go
// Bad - Big Interface
type Worker interface {
    Work()
    Eat()
}

type Robot struct{}

func (r Robot) Work() {}
func (r Robot) Eat() { 
    // What? Robots don't eat! Bad design.
}

// Good - Split Interface
type Workable interface {
    Work()
}

type Eatable interface {
    Eat()
}

type Human struct{}

func (h Human) Work() {}
func (h Human) Eat() {}

type Android struct{}

func (a Android) Work() {}
// Android doesn't implement Eat() — because it doesn't need to!
```


---

## 5. **D — Dependency Inversion Principle (DIP)**

🔴 **Meaning**:  
High-level modules should **not depend on low-level modules**.  
Both should depend on **abstractions** (like interfaces).

🔴 **Simple Words**:  
- The big boss (main service) should not care about the small workers (details).
- It should just trust that the workers know how to do their job (through an interface).

✅ **Tip**:  
If you add a new notifier (like Push Notification), you just create a new struct — no change needed in the main service.

```go
// Abstraction
type Notifier interface {
    Send(message string) error
}

// Low-level modules
type EmailNotifier struct{}

func (e EmailNotifier) Send(message string) error {
    fmt.Println("Sending Email:", message)
    return nil
}

type SMSNotifier struct{}

func (s SMSNotifier) Send(message string) error {
    fmt.Println("Sending SMS:", message)
    return nil
}

// High-level module
type OrderService struct {
    notifier Notifier
}

func (o OrderService) NotifyCustomer(message string) {
    o.notifier.Send(message)
}
```

---

# Summary Table 📋

| Principle | Meaning in 5 Words | Key Focus |
|:----------|:-------------------|:----------|
| SRP | One Class = One Responsibility | Separation of concerns |
| OCP | Extend, Don’t Modify | Easy to add features |
| LSP | Child Classes Work Properly | No unexpected behavior |
| ISP | Small, Focused Interfaces | No unnecessary dependencies |
| DIP | Depend on Abstractions | Loose coupling, flexible code |

---

# Final Tip 🎯

- Think of **SOLID** as rules that make your code like **Lego blocks**:  
  **Easy to connect, easy to replace, easy to grow!**
- They help you **write better**, **cleaner**, and **more professional** code, especially in large systems.

---