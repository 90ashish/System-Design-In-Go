# Designing an Elevator System

## Requirements

- The elevator system should consist of multiple elevators serving multiple floors.
- Each elevator should have a capacity limit and should not exceed it.
- Users should be able to request an elevator from any floor and select a destination floor.
- The elevator system should efficiently handle user requests and optimize the movement of elevators to minimize waiting time.
- The system should prioritize requests based on the direction of travel and the proximity of the elevators to the requested floor.
- The elevators should be able to handle multiple requests concurrently and process them in an optimal order.
- The system should ensure thread safety and prevent race conditions when multiple threads interact with the elevators.

## Classes, Interfaces, and Enumerations

- **Direction (enum)**: Represents the possible directions of elevator movement (`UP` or `DOWN`).
- **Request (class)**: Represents a user request for an elevator, containing the source floor and destination floor.
- **Elevator (class)**: Represents an individual elevator in the system. It has a capacity limit and maintains a list of requests. The elevator processes requests concurrently and moves between floors based on the requests.
- **ElevatorController (class)**: Manages multiple elevators and handles user requests. It finds the optimal elevator to serve a request based on the proximity of the elevators to the requested floor.
- **ElevatorSystem (class)**: Entry point of the application and demonstrates the usage of the elevator system.

## How It Works

1. **Initialization**  
   The `main` function creates an `ElevatorController` with a fixed number of elevators and capacity, then simulates requests by calling `RequestElevator(source, destination)`.

2. **Elevator Dispatch**  
   Each new request is sent to `findOptimalElevator`, which scores elevators based on direction match, idle status and distance, favoring those already moving toward the request or idle and nearby.

3. **Asynchronous Processing**  
   Each `Elevator` runs in its own goroutine, reading from a buffered `requests` channel and moving floor-by-floor (simulated with `time.Sleep(1s)`), printing its progress.

4. **Direction Handling**  
   Elevators update their `currentDirection` (up or down) before each request and serve only those that best match their state, ensuring realistic traffic handling.

5. **Graceful Shutdown**  
   Calling `Stop()` closes each elevator’s `stopChan`, terminating its processing loop cleanly when the program exits.
