# Go Logging Framework

A thread-safe, extensible logging framework in Go supporting multiple log levels and output destinations. Designed with clean separations and common design patterns to make it easy to add new log levels or appenders.

---

## Table of Contents

- [Features](#features)  
- [Getting Started](#getting-started)  
- [Project Structure](#project-structure)  
- [Usage](#usage)  
  - [Configuration (Builder)](#configuration-builder)  
  - [Logging](#logging)  
- [Built-In Appenders](#built-in-appenders)  
- [Design Patterns](#design-patterns)  
- [Concurrency & Thread-Safety](#concurrency--thread-safety)  
- [Extending the Framework](#extending-the-framework)  
- [Example](#example)  
- [License](#license)  

---

## Features

- **Log Levels**: `DEBUG`, `INFO`, `WARNING`, `ERROR`, `FATAL`  
- **Thread-Safe**: Safe for concurrent use from multiple goroutines  
- **Multiple Destinations**: Console, file, PostgreSQL database  
- **Configurable**: Set log level and appenders at runtime  
- **Pluggable**: Easily add new appenders (e.g. remote HTTP, Kafka)  
- **Singleton**: Single global `Logger` instance  
- **Builder**: Fluent API for configuring your logger  

---

## Getting Started

 **Clone the repo**  
   ```bash
   git clone https://github.com/your-org/go-logging-framework.git
   cd go-logging-framework
   go build ./cmd/logging-app
   ./logging-app
   ```

## Project Structure

```
project-root/
├── cmd/logging-app/
│   └── main.go              # example application
├── internal/
│   └── logging/
│       ├── level.go         # LogLevel enum
│       ├── message.go       # LogMessage struct
│       ├── appender.go      # Appender interface
│       ├── config.go        # Builder for LoggerConfig
│       ├── logger.go        # Singleton Logger
│       └── appenders/
│           ├── console.go   # ConsoleAppender
│           ├── file.go      # FileAppender
│           └── database.go  # DatabaseAppender
└── go.mod
```

- **cmd/logging-app**: Demonstrates usage and concurrent logging.
- **internal/logging**: Core framework (levels, messages, config, logger).
- **internal/logging/appenders**: Pre-built destinations.

---

## Usage

### Configuration (Builder)

Use the Builder pattern to create a `LoggerConfig`:

```go
cfg := logging.
    NewConfigBuilder().
    Level(logging.DEBUG).
    AddAppender(appenders.NewConsoleAppender()).
    AddAppender(fileAppender).
    AddAppender(dbAppender).
    Build()
```

- `Level(...)` sets the minimum level to output.
- `AddAppender(...)` registers any number of Appender implementations.

Apply it to the singleton logger:

```go
log := logging.GetLogger()
log.Configure(cfg)
```

###Logging

- Once configured, call convenience methods from anywhere:

```go
log.Info("Server starting up")
log.Debug("Cache miss for key X")
log.Warning("Disk space low")
log.Error("Failed to connect to service")
log.Fatal("Unrecoverable error, exiting")
```

## Built-In Appenders
- **ConsoleAppender**  
  Writes formatted messages to standard output.
- **FileAppender**  
  Appends logs to a configurable file path.
- **DatabaseAppender**  
  Inserts logs into a PostgreSQL table (`logs`).

All appenders implement the `Appender` interface:

```go
type Appender interface {
    Append(message LogMessage) error
}
```

## Design Patterns
- **Singleton** – One global `Logger` instance.
- **Builder** – Fluent `ConfigBuilder` to assemble `LoggerConfig`.
- **Strategy/Observer** – `Appender` interface lets you swap in or register new destinations.
- **Factory** (optional) – You can add a factory for dynamic `Appender` creation.

## Concurrency & Thread-Safety
- `sync.RWMutex` guards the logger’s config during reads/writes.
- Each appender uses `sync.Mutex` to serialize I/O (console, file, DB).
- Log messages are timestamped with `time.Now()` to preserve ordering.

## Extending the Framework

### To Add a New Output (e.g., HTTP, Kafka)
- Create a new file under `internal/logging/appenders/`.
- Implement the `Append(LogMessage) error` method.
- Register it in your configuration via `AddAppender(newMyAppender(...))`.

### To Introduce a New Log Level
- Extend the `LogLevel` enum in `level.go`.
- Add any convenience method to `logger.go` (e.g., `Verbose()`).
