# Google System Design Interview Question: Document Editor Using Interface Segregation Principle (ISP)

## Question

Design a document editor system that supports multiple types of documents such as `TextDocument`, `PrintableDocument`, and `ScannableDocument`. Each document type may support different capabilities—editing, printing, or scanning.

Using the **Interface Segregation Principle (ISP)**, ensure that the system avoids forcing document types to implement methods they do not need. Create appropriate interfaces and demonstrate how different document types implement only the interfaces relevant to their functionality. Provide a working implementation in Golang and explain how your design follows ISP.

---

## Solution in Golang

The **Interface Segregation Principle (ISP)** states:

> “Clients should not be forced to depend on methods they do not use.”

Instead of defining one big `Document` interface with all possible operations (edit, print, scan), we break it into smaller, role-specific interfaces:

- `Editable` — For document types that can be edited.
- `Printable` — For documents that can be printed.
- `Scannable` — For scanned inputs or documents that support scanning.

---

## ✅ Explanation

### Interfaces are Segregated by Responsibility

- `Editable` → For document types that support editing.
- `Printable` → For document types that support printing.
- `Scannable` → For document types that support scanning.

### Document Types Implement Only the Interfaces They Need

- `TextDocument` implements only `Editable`.
- `PrintableDocument` implements only `Printable`.
- `ScannableDocument` implements only `Scannable`.
- `MultiFunctionalDocument` implements all three interfaces.

---

## ✅ Why This Follows ISP

- No document type is forced to implement irrelevant methods.
- Client code can rely on small, focused interfaces tailored to its needs.
- This improves modularity and testability.

---

## ✅ Benefits of Using ISP in This Design

- **Reduces Coupling**: Each module or class depends only on the functionality it uses.
- **Increases Flexibility**: New capabilities (e.g., `Shareable`) can be added without impacting existing types.
- **Improves Readability and Maintainability**: Smaller interfaces are easier to understand and evolve independently.

---

## ✅ Interview Follow-Up Questions Google Might Ask

- **How would you modify this design to support a plugin system for future document operations?**
- **What are the trade-offs between creating many small interfaces versus one larger interface?**
- **Can you apply this principle in microservices or REST APIs?**

---

By applying the **Interface Segregation Principle**, this design ensures that document types are cleanly separated by functionality, promoting scalability, clarity, and robust interface design.
