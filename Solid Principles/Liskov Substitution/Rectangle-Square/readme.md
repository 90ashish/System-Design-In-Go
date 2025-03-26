# Interview Question: Designing a Rectangle-Square Hierarchy Using Liskov Substitution Principle (LSP)

## Question

In many object-oriented systems, a common example used to illustrate the challenges of inheritance is the relationship between rectangles and squares. Traditionally, one might be tempted to make `Square` a subclass of `Rectangle` because a square "is a" rectangle. However, this naive design can lead to violations of the **Liskov Substitution Principle (LSP)**.

Design a solution in Golang that models rectangles and squares while respecting the Liskov Substitution Principle. In your design, ensure that:

- A square is modeled appropriately without forcing it to inherit behavior that does not make sense for squares.
- The design supports computing the area of both shapes.
- The design allows using both shapes interchangeably where appropriate without violating LSP.

---

## Explanation

### Common Abstraction

- A `Shape` interface is defined with the method `Area() float64`, which abstracts the concept of area calculation.
- Both `Rectangle` and `Square` implement the `Area()` method, allowing them to be used interchangeably where a `Shape` is expected.

### Separate Implementations

- **Rectangle**:
  - A `Rectangle` is defined with `Width` and `Height`.
  - Its `Area()` method multiplies these two dimensions.
  
- **Square**:
  - A `Square` is defined with a single `Side` field.
  - Its `Area()` method calculates the area as `Side * Side`.

- By modeling `Square` as its own type (not as a subclass of `Rectangle`), the design avoids the common pitfalls that lead to violations of LSP—such as trying to set width and height independently on a square.

### Liskov Substitution Principle (LSP)

- The `DisplayArea` function (or any client function) takes a `Shape` and calls its `Area()` method without knowing whether it is a `Rectangle` or a `Square`.
- This design ensures that any object implementing `Shape` can be substituted seamlessly, maintaining expected behavior regardless of the specific type.

### Benefits of the Design

- **Flexibility**: New shapes can be added easily by implementing the `Shape` interface.
- **Correctness**: Each shape's logic is encapsulated in its own type, preserving its unique constraints (e.g., all sides equal in a square).
- **Interchangeability**: Functions that rely on the `Shape` abstraction (like `DisplayArea`) work with any shape without needing to change their logic.

---

## Conclusion

This design illustrates how to correctly model a rectangle-square hierarchy while adhering to the **Liskov Substitution Principle**. By using interfaces and avoiding inappropriate inheritance, the system maintains flexibility, correctness, and substitution safety—core goals of clean, maintainable design.
