# Google System Design Interview Question: Media Player System Using Liskov Substitution Principle (LSP)

## Question

Design a media player system that supports multiple media types (e.g., Audio and Video) following the **Liskov Substitution Principle (LSP)**. The media player should be able to work with any media type without altering its functionality. Ensure that any subclass (or concrete type) of the media abstraction can replace the base type without affecting the correctness of the system. Explain your design and provide a sample implementation in Golang.

---

## Solution in Golang

The following implementation demonstrates the Liskov Substitution Principle by:

### Defining a Common Abstraction

- A `Media` interface is defined with a method `Play()`.  
- Both audio and video types implement this interface.

### Ensuring Substitutability

- The `MediaPlayer` struct depends on the `Media` interface rather than concrete types.  
- This allows any type that implements `Media` to be substituted seamlessly.

### Consistent Behavior

- Each media type implements its own `Play()` method.  
- This ensures that the client code (in this case, `MediaPlayer`) can operate on them without knowing their internal details.

---

## Explanation

### Common Abstraction

- The `Media` interface defines a contract with the `Play()` method.  
- Both `Audio` and `Video` implement this interface, ensuring that they can be used interchangeably.

### Liskov Substitution Principle (LSP)

- The `MediaPlayer` struct is designed to work with the `Media` interface.  
- Whether an `Audio` or `Video` object is provided, `MediaPlayer` can call the `Play()` method without concern for the specific type.  
- This ensures that any object that satisfies the `Media` interface can substitute another, adhering to LSP.

### Extensibility

- New media types (e.g., `Podcast`, `LiveStream`) can be added by implementing the `Media` interface.  
- This does not require changes to the `MediaPlayer`, enabling open-ended extensibility.

### Consistent Behavior

- The `PlayMedia()` method in `MediaPlayer` delegates the call to the `Play()` method of the concrete media type.  
- This ensures consistent behavior regardless of the media type being used.

---

## Conclusion

This design and implementation demonstrate how adhering to the **Liskov Substitution Principle** allows for building flexible, extensible systems that can easily incorporate new functionalities without modifying the core logic.
