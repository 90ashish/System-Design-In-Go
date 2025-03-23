Amazon System Design Interview Question: Single Responsibility Principle (SRP)

Question:
Design a logging system for an e-commerce platform that tracks order processing. The system should adhere to the Single Responsibility Principle (SRP). The logging system should:

    Log information related to Order Processing events.

    Allow flexible logging to different outputs such as console, file, or external monitoring services.

    Ensure that the Order Processing logic remains separate from the Logging logic.

Solution in Golang Implementing SRP

Here's a well-structured solution in Golang that follows SRP:

    OrderProcessor: Responsible for processing orders.

    Logger: Responsible for handling logs independently.


Explanation

✅ Single Responsibility Principle (SRP):

    The OrderProcessor class is only responsible for processing orders.

    The Logger (with implementations like ConsoleLogger and FileLogger) is solely responsible for logging messages.

✅ Scalability & Flexibility:

    New log destinations (e.g., external services like Amazon CloudWatch or ElasticSearch) can be added without modifying OrderProcessor.

    This design leverages Dependency Injection, ensuring loose coupling.

✅ Real-World Readiness:

    The solution is clean, modular, and adheres to SRP, which is a crucial aspect of scalable system design.

Follow-Up Questions Amazon May Ask

    How would you extend this to support multiple log levels (INFO, WARN, ERROR, DEBUG)?

    What changes are needed to implement a retry mechanism for logging failures?

    How would you modify the solution to batch log entries for performance improvements?