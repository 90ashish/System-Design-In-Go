package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Service represents a generic service with its own bulkhead (semaphore).
type Service struct {
	name          string
	workerPool    chan struct{} // buffered channel to limit concurrency
	fallbackFunc  func(request string) string
	processFunc   func(request string) (string, error)
}

// NewService creates a new Service with a given name, max concurrency, processing function, and fallback.
func NewService(name string, maxConcurrent int, processFunc func(request string) (string, error), fallbackFunc func(request string) string) *Service {
	return &Service{
		name:         name,
		workerPool:   make(chan struct{}, maxConcurrent),
		processFunc:  processFunc,
		fallbackFunc: fallbackFunc,
	}
}

// HandleRequest attempts to process a request. If the worker pool is full, it uses the fallback.
func (s *Service) HandleRequest(request string) string {
	select {
	case s.workerPool <- struct{}{}:
		// Got a slot to process the request.
		defer func() { <-s.workerPool }() // release the slot after processing
		result, err := s.processFunc(request)
		if err != nil {
			log.Printf("[%s] Error processing request '%s': %v. Executing fallback.", s.name, request, err)
			return s.fallbackFunc(request)
		}
		return result
	default:
		// Bulkhead is full: immediately fallback.
		log.Printf("[%s] Bulkhead full for request '%s'. Executing fallback.", s.name, request)
		return s.fallbackFunc(request)
	}
}

// Dummy process functions simulate work with random delays and occasional failures.
func documentProcess(request string) (string, error) {
	// simulate processing time
	time.Sleep(time.Duration(rand.Intn(100)+50) * time.Millisecond)
	// simulate random failure
	if rand.Float32() < 0.2 {
		return "", fmt.Errorf("document processing failed")
	}
	return fmt.Sprintf("Document processed: %s", request), nil
}

func emailNotification(request string) (string, error) {
	time.Sleep(time.Duration(rand.Intn(100)+50) * time.Millisecond)
	if rand.Float32() < 0.15 {
		return "", fmt.Errorf("email notification failed")
	}
	return fmt.Sprintf("Email sent: %s", request), nil
}

func dataAnalytics(request string) (string, error) {
	time.Sleep(time.Duration(rand.Intn(100)+50) * time.Millisecond)
	if rand.Float32() < 0.25 {
		return "", fmt.Errorf("data analytics failed")
	}
	return fmt.Sprintf("Analytics complete for: %s", request), nil
}

// Fallback functions provide alternate behavior when the service is overloaded or fails.
func fallbackDocument(request string) string {
	return fmt.Sprintf("Fallback: Document queued for later processing: %s", request)
}

func fallbackEmail(request string) string {
	return fmt.Sprintf("Fallback: Email will be retried later: %s", request)
}

func fallbackAnalytics(request string) string {
	return fmt.Sprintf("Fallback: Analytics deferred for: %s", request)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create services with a concurrency limit
	docService := NewService("DocumentProcessing", 3, documentProcess, fallbackDocument)
	emailService := NewService("EmailNotification", 2, emailNotification, fallbackEmail)
	analyticsService := NewService("DataAnalytics", 2, dataAnalytics, fallbackAnalytics)

	// Simulate a burst of concurrent requests
	requests := []string{"Req1", "Req2", "Req3", "Req4", "Req5", "Req6", "Req7", "Req8"}

	// Channel to wait for all processing to finish.
	done := make(chan bool)

	// Process document requests concurrently.
	go func() {
		for _, req := range requests {
			go func(r string) {
				result := docService.HandleRequest(r)
				log.Println(result)
			}(req)
		}
		done <- true
	}()

	// Process email requests concurrently.
	go func() {
		for _, req := range requests {
			go func(r string) {
				result := emailService.HandleRequest(r)
				log.Println(result)
			}(req)
		}
		done <- true
	}()

	// Process analytics requests concurrently.
	go func() {
		for _, req := range requests {
			go func(r string) {
				result := analyticsService.HandleRequest(r)
				log.Println(result)
			}(req)
		}
		done <- true
	}()

	// Wait for all goroutines to be initiated.
	<-done
	<-done
	<-done

	// Allow some time for all requests to finish processing.
	time.Sleep(2 * time.Second)
	log.Println("All requests processed.")
}
