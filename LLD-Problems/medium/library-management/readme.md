# Designing a Library Management System

## Requirements

- The system should allow librarians to manage **books**, **members**, and **borrowing activities**.
- Support for **adding**, **updating**, and **removing** books from the library catalog.
- Each book should include:
  - Title
  - Author
  - ISBN
  - Publication Year
  - Availability Status
- Members should be able to **borrow** and **return** books.
- Each member should include:
  - Name
  - Member ID
  - Contact Information
  - Borrowing History
- Enforce **borrowing rules**, such as:
  - Maximum number of books borrowed at a time
  - Loan duration
- Handle **concurrent access** to catalog and member records.
- System should be **extensible** for future enhancements and new features.

---

## Classes, Interfaces, and Enumerations

- **`Book` Class**  
  Represents a book in the library catalog with:
  - ISBN
  - Title
  - Author
  - Publication Year
  - Availability Status

- **`Member` Class**  
  Represents a library member with:
  - Member ID
  - Name
  - Contact Information
  - List of Borrowed Books

- **`LibraryManager` Class**  
  Core component of the system:
  - Follows the **Singleton Pattern** to ensure a single instance
  - Uses **ConcurrentHashMap** to manage:
    - Library Catalog
    - Member Records
  - Provides methods for:
    - Adding/Removing Books
    - Registering/Unregistering Members
    - Borrowing/Returning Books
    - Searching for Books by keywords

- **`LibraryManagementSystemDemo` Class**  
  Entry point of the application demonstrating system functionality.
