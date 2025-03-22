package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Circuit Breaker States
const (
	Closed    = "Closed"
	Open      = "Open"
	HalfOpen  = "Half-Open"
)

// CircuitBreaker struct
type CircuitBreaker struct {
	mu           sync.Mutex
	state        string
	failureCount int
	resetTimeout time.Duration
	lastAttempt  time.Time
}

// NewCircuitBreaker - Initializes a new CircuitBreaker
func NewCircuitBreaker(resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:        Closed,
		resetTimeout: resetTimeout,
	}
}

// CallService - Executes a request based on the Circuit Breaker state
func (cb *CircuitBreaker) CallService(service func() error) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	switch cb.state {
	case Open:
		// Cooldown Period Check
		if time.Since(cb.lastAttempt) > cb.resetTimeout {
			fmt.Println("🔄 Transitioning to Half-Open state...")
			cb.state = HalfOpen
		} else {
			return errors.New("🚨 Circuit Breaker OPEN - Request Blocked")
		}

	case HalfOpen:
		err := service()
		if err != nil {
			fmt.Println("❌ Service failed in Half-Open state. Returning to Open state.")
			cb.state = Open
			cb.lastAttempt = time.Now()
			return err
		}

		fmt.Println("✅ Service restored successfully. Switching back to Closed state.")
		cb.state = Closed
		cb.failureCount = 0
		return nil

	case Closed:
		err := service()
		if err != nil {
			cb.failureCount++
			fmt.Printf("⚠️ Service failed. Failure Count: %d\n", cb.failureCount)

			if cb.failureCount >= 3 {
				fmt.Println("🚨 Too many failures! Switching to Open state.")
				cb.state = Open
				cb.lastAttempt = time.Now()
			}
			return err
		}

		cb.failureCount = 0
		return nil
	}

	return nil
}

// Simulate External Service
func MockService() error {
	if rand.Intn(100) < 60 { // 60% chance of failure
		return errors.New("Service Unavailable")
	}
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	cb := NewCircuitBreaker(5 * time.Second) // Cooldown period: 5 seconds
	var wg sync.WaitGroup

	for i := 1; i <= 15; i++ {
		wg.Add(1)
		go func(requestID int) {
			defer wg.Done()
			err := cb.CallService(MockService)
			if err != nil {
				fmt.Printf("❌ Request %d failed: %v\n", requestID, err)
			} else {
				fmt.Printf("✅ Request %d succeeded!\n", requestID)
			}
			time.Sleep(1 * time.Second)
		}(i)
	}

	wg.Wait()
	fmt.Println("🚀 All requests processed successfully.")
}
