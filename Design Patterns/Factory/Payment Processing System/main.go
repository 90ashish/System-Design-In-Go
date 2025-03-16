package main

import (
	"fmt"
	"sync"
	"time"
)

// PaymentMethod interface (Common Interface)
type PaymentMethod interface {
	ProcessPayment(amount float64)
}

// CreditCard struct
type CreditCard struct{}

func (c *CreditCard) ProcessPayment(amount float64) {
	time.Sleep(1 * time.Second) // Simulate processing delay
	fmt.Printf("[Credit Card] Payment of $%.2f processed successfully.\n", amount)
}

// PayPal struct
type PayPal struct{}

func (p *PayPal) ProcessPayment(amount float64) {
	time.Sleep(1 * time.Second) // Simulate processing delay
	fmt.Printf("[PayPal] Payment of $%.2f processed successfully.\n", amount)
}

// AmazonPay struct
type AmazonPay struct{}

func (a *AmazonPay) ProcessPayment(amount float64) {
	time.Sleep(1 * time.Second) // Simulate processing delay
	fmt.Printf("[Amazon Pay] Payment of $%.2f processed successfully.\n", amount)
}

// PaymentFactory - Factory for creating payment methods
type PaymentFactory struct{}

// CreatePaymentMethod - Factory method to generate the correct payment method
func (f *PaymentFactory) CreatePaymentMethod(method string) (PaymentMethod, error) {
	switch method {
	case "creditcard":
		return &CreditCard{}, nil
	case "paypal":
		return &PayPal{}, nil
	case "amazonpay":
		return &AmazonPay{}, nil
	default:
		return nil, fmt.Errorf("unsupported payment method: %s", method)
	}
}

// ProcessOrder - Uses Goroutines for concurrent payment processing
func ProcessOrder(wg *sync.WaitGroup, factory *PaymentFactory, method string, amount float64) {
	defer wg.Done()

	paymentMethod, err := factory.CreatePaymentMethod(method)
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		return
	}
	paymentMethod.ProcessPayment(amount)
}

func main() {
	factory := &PaymentFactory{}
	var wg sync.WaitGroup

	// Concurrent Payment Processing Using Goroutines
	transactions := []struct {
		method string
		amount float64
	}{
		{"creditcard", 100.50},
		{"paypal", 250.75},
		{"amazonpay", 150.00},
		{"bitcoin", 500.00}, // Unsupported payment method
		{"creditcard", 200.00},
	}

	for _, txn := range transactions {
		wg.Add(1)
		go ProcessOrder(&wg, factory, txn.method, txn.amount)
	}

	// Wait for all Goroutines to complete
	wg.Wait()

	fmt.Println("✅ All transactions have been processed successfully.")
}
