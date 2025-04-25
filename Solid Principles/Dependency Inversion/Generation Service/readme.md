# Google System Design Interview Question: Report Generation Service Using the Dependency Inversion Principle (DIP)

---

## Question:

You’re tasked with designing a **Report Generation Service** that can output business reports in multiple formats (**JSON**, **CSV**, **XML**).  
Following the **Dependency Inversion Principle (DIP)**, ensure that:

- **High-level modules** (the report generation orchestration) do **not depend directly** on low-level modules (formatters for JSON, CSV, XML).
- Both high-level and low-level modules **depend on abstractions**.
- **Abstractions do not depend on details**; details (the concrete formatters) depend on abstractions.

**Design Goal**:  
Adding a new format (e.g., **PDF**) should require **no changes** to the core report-generation logic — only a new implementation of the formatter interface.

---

## Why This Demonstrates the Dependency Inversion Principle

### High-Level Module (`ReportService`)
- Does **not import or instantiate** any concrete formatter.
- Depends **only** on the abstraction `ReportFormatter`.

### Low-Level Modules (`JSONFormatter`, `CSVFormatter`, `XMLFormatter`)
- Depend on the `ReportFormatter` interface, **not** on `ReportService`.
- Implement the `Format` method required by the interface.

### Abstraction Independence
- The `ReportFormatter` interface stands **between** high- and low-level code.
- Neither side depends on the other’s details — changes to formatters **never touch** `ReportService`, and vice versa.

### Extensibility
- To support a new format (e.g., **PDF**, **HTML**), you only need:
  ```go
  type PDFFormatter struct{ /* ...fields... */ }
  
  func (p PDFFormatter) Format(r Report) ([]byte, error) {
      /* ... */
  }
