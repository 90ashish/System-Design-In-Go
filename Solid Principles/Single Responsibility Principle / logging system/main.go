package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Logger Interface — Provides flexibility for different log outputs
type Logger interface {
	LogInfo(message string)
	LogError(message string)
}

// ConsoleLogger — Logs to console
type ConsoleLogger struct{}

func (c ConsoleLogger) LogInfo(message string) {
	log.Printf("[INFO] %s\n", message)
}

func (c ConsoleLogger) LogError(message string) {
	log.Printf("[ERROR] %s\n", message)
}

// FileLogger — Logs to a file
type FileLogger struct {
	file *os.File
}

func NewFileLogger(filePath string) *FileLogger {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	return &FileLogger{file: file}
}

func (f FileLogger) LogInfo(message string) {
	fmt.Fprintf(f.file, "[INFO] %s\n", message)
}

func (f FileLogger) LogError(message string) {
	fmt.Fprintf(f.file, "[ERROR] %s\n", message)
}

// OrderProcessor — Responsible for handling order-related operations
type OrderProcessor struct {
	logger Logger
}

func NewOrderProcessor(logger Logger) *OrderProcessor {
	return &OrderProcessor{logger: logger}
}

func (op *OrderProcessor) ProcessOrder(orderID string) {
	op.logger.LogInfo(fmt.Sprintf("Starting order processing: %s", orderID))
	
	// Simulated order processing logic
	time.Sleep(2 * time.Second)  // Simulate some processing delay
	op.logger.LogInfo(fmt.Sprintf("Order %s processed successfully!", orderID))
}

func main() {
	// Console logger example
	consoleLogger := ConsoleLogger{}
	orderProcessorConsole := NewOrderProcessor(consoleLogger)
	orderProcessorConsole.ProcessOrder("ORD-123")

	// File logger example
	fileLogger := NewFileLogger("order_logs.txt")
	orderProcessorFile := NewOrderProcessor(fileLogger)
	orderProcessorFile.ProcessOrder("ORD-456")
}
