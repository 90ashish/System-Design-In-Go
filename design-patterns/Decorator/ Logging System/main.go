package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Logger Interface
type Logger interface {
	Log(message string)
}

// BasicLogger - Core Logger Implementation
type BasicLogger struct{}

func (b *BasicLogger) Log(message string) {
	fmt.Printf("[Basic Log] %s\n", message)
}

// TimestampDecorator - Adds Timestamps to Logs
type TimestampDecorator struct {
	wrapped Logger
}

func (t *TimestampDecorator) Log(message string) {
	timestampedMessage := fmt.Sprintf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), message)
	t.wrapped.Log(timestampedMessage)
}

// ErrorLevelDecorator - Adds Error Level to Logs
type ErrorLevelDecorator struct {
	wrapped Logger
	level   string
}

func (e *ErrorLevelDecorator) Log(message string) {
	levelledMessage := fmt.Sprintf("[%s] %s", e.level, message)
	e.wrapped.Log(levelledMessage)
}

// FileOutputDecorator - Writes Logs to a File
type FileOutputDecorator struct {
	wrapped Logger
	filePath string
}

func (f *FileOutputDecorator) Log(message string) {
	// Append the log entry to a file
	file, err := os.OpenFile(f.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("[ERROR] Failed to open log file: %v\n", err)
		return
	}
	defer file.Close()

	logEntry := fmt.Sprintf("%s\n", message)
	file.WriteString(logEntry)
	fmt.Println("[File Log] Entry written to file:", f.filePath)
}

// ConcurrentLogger - Runs logging concurrently with Goroutines
func ConcurrentLogger(wg *sync.WaitGroup, logger Logger, message string) {
	defer wg.Done()
	logger.Log(message)
}

func main() {
	var wg sync.WaitGroup

	// Base Logger
	baseLogger := &BasicLogger{}

	// Decorators
	timestampedLogger := &TimestampDecorator{wrapped: baseLogger}
	errorLogger := &ErrorLevelDecorator{wrapped: timestampedLogger, level: "ERROR"}
	fileLogger := &FileOutputDecorator{wrapped: errorLogger, filePath: "app_logs.txt"}

	// Simulating Concurrent Log Entries
	logMessages := []string{
		"User login successful",
		"Database connection failed",
		"Server started successfully",
		"Error while processing request",
		"File uploaded successfully",
	}

	// Concurrent Log Processing
	for _, msg := range logMessages {
		wg.Add(1)
		go ConcurrentLogger(&wg, fileLogger, msg)
	}

	wg.Wait()

	fmt.Println("✅ All log entries have been processed successfully.")
}
