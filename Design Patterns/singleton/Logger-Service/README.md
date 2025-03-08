Problem Statement

Imagine you're building a Logger Service for a backend application. The logger should be designed such that:

    Only one instance of the logger should exist throughout the application lifecycle.
    The logger should support thread-safe operations since multiple Goroutines may need to log data concurrently.
    The logger should track how many logs have been recorded.

Question

    Design a Logger Service using the Singleton Design Pattern that ensures only one instance is created. Add functionality to record logs and count the total logs recorded.
