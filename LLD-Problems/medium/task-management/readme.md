# Task Management System Design

## Task Entity

The `Task` entity contains the following properties:

- **Title** (`string`)
- **Description** (`string`)
- **Priority** (`LOW`, `MEDIUM`, `HIGH`)
- **Due Date** (`datetime`)
- **Created Date** (`datetime`)
- **Status** (`OPEN`, `IN_PROGRESS`, `COMPLETED`)

---

## Features to Implement

- Create a Task
- Update a Task
- Delete a Task
- Fetch Tasks:
  - All Tasks
  - Filtered by Status / Priority / Due Date
- Fetch the Most Priority Task (task with the highest priority and nearest due date)
- Sort Tasks by Due Date

---

## Expectations

- Proper modeling of classes using good OOP design.
- Apply **SOLID** principles where applicable.
- Design should be **extensible** (e.g., new priority types can be added easily).
- Use **in-memory** storage (no DB layer required).
- **Bonus**: Ensure thread-safety for concurrent task creation/updating.

---

## Key Design Patterns & Principles

- **Factory Pattern**: 
  - `NewTask`
  - `NewInMemoryTaskRepository`
  
- **Repository Pattern**:
  - `TaskRepository` interface
  - `InMemoryTaskRepository` implementation

- **Strategy Pattern**:
  - `TaskFilter` functions for flexible querying

- **Facade Pattern**:
  - `TaskService` provides a unified API over multiple operations

- **Dependency Injection (DIP)**:
  - `TaskService` depends on the `TaskRepository` interface

- **Single Responsibility Principle**:
  - Separate entities, storage layer, and business logic

- **Thread Safety**:
  - Use `sync.RWMutex` to guard in-memory store

- **Open/Closed Principle**:
  - Support adding new filters or storage backends without modifying existing logic
