package main

import (
	"fmt"
)

// PaymentMethod interface (Common Interface)
type PaymentMethod interface {
	ProcessPayment(amount float64)
}

// CreditCard struct
type CreditCard struct{}

func (c *CreditCard) ProcessPayment(amount float64) {
	fmt.Printf("[Credit Card] Payment of $%.2f processed successfully.\n", amount)
}

// PayPal struct
type PayPal struct{}

func (p *PayPal) ProcessPayment(amount float64) {
	fmt.Printf("[PayPal] Payment of $%.2f processed successfully.\n", amount)
}

// AmazonPay struct
type AmazonPay struct{}

func (a *AmazonPay) ProcessPayment(amount float64) {
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

// ProcessOrder - Simulates processing a payment using the factory
func ProcessOrder(factory *PaymentFactory, method string, amount float64) {
	paymentMethod, err := factory.CreatePaymentMethod(method)
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		return
	}
	paymentMethod.ProcessPayment(amount)
}

func main() {
	factory := &PaymentFactory{}

	// Test different payment methods
	ProcessOrder(factory, "creditcard", 100.50)
	ProcessOrder(factory, "paypal", 250.75)
	ProcessOrder(factory, "amazonpay", 150.00)
	ProcessOrder(factory, "bitcoin", 500.00) // Unsupported payment method
}
