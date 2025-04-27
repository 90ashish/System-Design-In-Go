package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// PaymentService - Simulates a payment service with limited capacity
func PaymentService(requestID int, paymentChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case paymentChannel <- requestID:
		// Simulate payment processing time
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Printf("[✅ Payment Service] Processed Payment Request %d\n", requestID)
	default:
		fmt.Printf("[❌ Payment Service] Request %d rejected due to overload\n", requestID)
	}
}

// InventoryService - Simulates an inventory service with limited capacity
func InventoryService(requestID int, inventoryChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case inventoryChannel <- requestID:
		time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
		fmt.Printf("[✅ Inventory Service] Updated stock for Request %d\n", requestID)
	default:
		fmt.Printf("[❌ Inventory Service] Request %d rejected due to overload\n", requestID)
	}
}

// NotificationService - Simulates a notification service with limited capacity
func NotificationService(requestID int, notificationChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case notificationChannel <- requestID:
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Printf("[✅ Notification Service] Sent confirmation for Request %d\n", requestID)
	default:
		fmt.Printf("[❌ Notification Service] Request %d rejected due to overload\n", requestID)
	}
}

func main() {
	var wg sync.WaitGroup

	// Bulkhead pattern implementation using buffered channels
	paymentChannel := make(chan int, 3)         // Payment Service can handle 3 concurrent requests
	inventoryChannel := make(chan int, 2)       // Inventory Service can handle 2 concurrent requests
	notificationChannel := make(chan int, 5)    // Notification Service can handle 5 concurrent requests

	requests := 5 // Total incoming requests

	for i := 1; i <= requests; i++ {
		wg.Add(3)
		go PaymentService(i, paymentChannel, &wg)
		go InventoryService(i, inventoryChannel, &wg)
		go NotificationService(i, notificationChannel, &wg)
	}

	wg.Wait()
	close(paymentChannel)
	close(inventoryChannel)
	close(notificationChannel)

	fmt.Println("✅ All requests have been processed with Bulkhead Isolation.")
}
