package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

// Logger struct (Singleton)
type Logger struct {
	file     *os.File
	logCount int
}

// Singleton instance and sync object
var (
	instance *Logger
	once     sync.Once
)

// GetInstance - Provides the singleton instance
func GetInstance() *Logger {
	once.Do(func() {
		file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal("Failed to create log file:", err)
		}

		instance = &Logger{
			file: file,
		}
	})
	return instance
}

// Log - Logs a message and updates log count
func (l *Logger) Log(message string) {
	log.SetOutput(l.file)
	log.Println(message)
	fmt.Println(message)
	l.logCount++
}

// GetLogCount - Returns the total number of logs recorded
func (l *Logger) GetLogCount() int {
	return l.logCount
}

// Simulate multiple Goroutines logging
func simulateLogs(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	logger := GetInstance()
	logger.Log(fmt.Sprintf("Log entry from Goroutine %d", id))
}

func main() {
	var wg sync.WaitGroup
	logger := GetInstance()

	// Simulate concurrent logging
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go simulateLogs(&wg, i)
	}

	wg.Wait()

	// Final log count
	fmt.Printf("Total logs recorded: %d\n", logger.GetLogCount())
}
