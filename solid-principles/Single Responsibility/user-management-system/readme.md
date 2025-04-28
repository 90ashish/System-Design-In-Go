# Refactoring a Bloated User Management Class Using SRP

## Question
You have inherited a legacy user management system where a single class handles multiple responsibilities such as:
- **User Authentication**
- **Email Notifications**
- **Database Operations**

This monolithic class has grown too complex, making it difficult to maintain and extend.

### Your Task
Refactor this system by applying the **Single Responsibility Principle (SRP)** in Golang by:
- Implementing separate components (structs) for:
  - **User Authentication:** Handles user login and credential verification.
  - **Email Notifications:** Sends emails to users (e.g., welcome emails, password resets).
  - **Database Operations:** Performs CRUD operations for user data.
- Refactoring the bloated class into modular units where each component is responsible for a single aspect of the user management process.

---

## Explanation

### Modular Design
- **Authentication:** The `AuthService` is solely responsible for verifying user credentials, interacting with the database through an interface.
- **Email Notifications:** The `EmailService` handles sending emails, with its responsibility limited to email-related tasks.
- **Database Operations:** The `DBService` manages user data storage and retrieval.
- **UserManagement:** Acts as a coordinator that utilizes the above services for user registration and login.

### Benefits of SRP Refactoring
- **Maintainability:** Each component can be maintained or extended independently without impacting others.
- **Testability:** Smaller, focused units are easier to test in isolation.
- **Scalability:** The system can be scaled horizontally by improving or replacing individual modules.
- **Loose Coupling:** Dependency injection ensures that the components remain loosely coupled, facilitating easier integration and replacement.
