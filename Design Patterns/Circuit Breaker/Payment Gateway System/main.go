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
	Closed     = "CLOSED"
	Open       = "OPEN"
	HalfOpen   = "HALF-OPEN"
	FailureThreshold = 3   // Failures before circuit opens
	RetryTimeout      = 5  // Time to wait before retrying (Half-Open state)
)

// CircuitBreaker - Circuit Breaker Struct
type CircuitBreaker struct {
	state          string
	failureCount   int
	mutex          sync.Mutex
	lastFailureTime time.Time
}

// NewCircuitBreaker - Initializes a new Circuit Breaker
func NewCircuitBreaker() *CircuitBreaker {
	return &CircuitBreaker{
		state:        Closed,
		failureCount: 0,
	}
}

// AllowRequest - Decides if the request should proceed based on circuit state
func (cb *CircuitBreaker) AllowRequest() bool {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	switch cb.state {
	case Closed:
		return true
	case Open:
		if time.Since(cb.lastFailureTime) > RetryTimeout*time.Second {
			fmt.Println("⏳ Circuit Breaker transitioning to HALF-OPEN for retry...")
			cb.state = HalfOpen
			return true
		}
		return false
	case HalfOpen:
		return true
	default:
		return false
	}
}

// RecordSuccess - Marks the request as successful and resets breaker
func (cb *CircuitBreaker) RecordSuccess() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	fmt.Println("✅ Circuit Breaker RESET to CLOSED after successful retry.")
	cb.state = Closed
	cb.failureCount = 0
}

// RecordFailure - Marks the request as failed
func (cb *CircuitBreaker) RecordFailure() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	cb.failureCount++
	if cb.failureCount >= FailureThreshold {
		cb.state = Open
		cb.lastFailureTime = time.Now()
		fmt.Println("🚨 Circuit Breaker OPEN: Blocking requests to prevent overload.")
	}
}

// SimulatedPaymentService - Simulates a payment request to an external API
func SimulatedPaymentService() error {
	// Simulating 70% chance of failure
	if rand.Intn(10) < 7 {
		return errors.New("Payment Service Failed")
	}
	return nil
}

// ProcessPayment - Handles payment requests with Circuit Breaker protection
func ProcessPayment(wg *sync.WaitGroup, cb *CircuitBreaker, transactionID int) {
	defer wg.Done()

	if !cb.AllowRequest() {
		fmt.Printf("[❌ Transaction %d] Request blocked due to Circuit Breaker OPEN state.\n", transactionID)
		return
	}

	err := SimulatedPaymentService()
	if err != nil {
		fmt.Printf("[❌ Transaction %d] Failed: %v\n", transactionID, err)
		cb.RecordFailure()
	} else {
		fmt.Printf("[✅ Transaction %d] Payment Processed Successfully!\n", transactionID)
		cb.RecordSuccess()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	circuitBreaker := NewCircuitBreaker()

	// Simulating Concurrent Transactions with Goroutines
	for i := 1; i <= 15; i++ {
		wg.Add(1)
		go ProcessPayment(&wg, circuitBreaker, i)
		time.Sleep(500 * time.Millisecond) // Delay between requests to observe retry timing
	}

	wg.Wait()

	fmt.Println("✅ All transactions processed successfully.")
}
