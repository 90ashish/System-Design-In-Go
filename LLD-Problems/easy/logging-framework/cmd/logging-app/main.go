package main

import (
	"fmt"
	"sync"
	"time"

	"logger/internal/logging"
	"logger/internal/logging/appenders"
)

func main() {
	// Build a config with Console, File and DB appenders
	console := appenders.NewConsoleAppender()

	fileApp, err := appenders.NewFileAppender("app.log")
	if err != nil {
		fmt.Println("FileAppender error:", err)
		return
	}

	dbApp, err := appenders.NewDatabaseAppender("postgres://ashishjaiswal@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		fmt.Println("DBAppender error:", err)
		return
	}

	cfg := logging.
		NewConfigBuilder().
		Level(logging.DEBUG).
		AddAppender(console).
		AddAppender(fileApp).
		AddAppender(dbApp).
		Build()

	log := logging.GetLogger()
	log.Configure(cfg)

	// Example usage:
	log.Info("Application starting up")
	log.Debug("This is a debug message")

	// Simulate concurrent logging
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				log.Warning(fmt.Sprintf("goroutine %d iteration %d", id, j))
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}
	wg.Wait()

	log.Error("Something went wrong")
	log.Fatal("Fatal error, exiting")
}
