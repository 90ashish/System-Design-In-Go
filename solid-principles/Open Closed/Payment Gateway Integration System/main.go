package main

import (
	"errors"
	"fmt"
)

// PaymentRequest represents a payment request
type PaymentRequest struct {
	Amount   float64
	Currency string
}

// PaymentResponse represents a payment response
type PaymentResponse struct {
	Provider string
	Status   string
	Message  string
}

// PaymentProvider defines an abstraction for payment providers.
type PaymentProvider interface {
	ProcessPayment(request PaymentRequest) (PaymentResponse, error)
}

// PayPalProvider implements PaymentProvider interface.
type PayPalProvider struct{}

func (p PayPalProvider) ProcessPayment(request PaymentRequest) (PaymentResponse, error) {
	// Simulate processing payment with PayPal
	fmt.Println("Processing payment with PayPal...")
	// In real-world code, call PayPal API here.
	return PaymentResponse{
		Provider: "PayPal",
		Status:   "Success",
		Message:  fmt.Sprintf("Processed $%.2f %s via PayPal", request.Amount, request.Currency),
	}, nil
}

// StripeProvider implements PaymentProvider interface.
type StripeProvider struct{}

func (s StripeProvider) ProcessPayment(request PaymentRequest) (PaymentResponse, error) {
	// Simulate processing payment with Stripe
	fmt.Println("Processing payment with Stripe...")
	// In real-world code, call Stripe API here.
	return PaymentResponse{
		Provider: "Stripe",
		Status:   "Success",
		Message:  fmt.Sprintf("Processed $%.2f %s via Stripe", request.Amount, request.Currency),
	}, nil
}

// PaymentGateway orchestrates the payment processing using a PaymentProvider.
type PaymentGateway struct {
	provider PaymentProvider
}

// NewPaymentGateway initializes a new PaymentGateway with the provided PaymentProvider.
func NewPaymentGateway(provider PaymentProvider) *PaymentGateway {
	return &PaymentGateway{provider: provider}
}

// Process handles the payment request using the injected PaymentProvider.
func (pg *PaymentGateway) Process(request PaymentRequest) (PaymentResponse, error) {
	if request.Amount <= 0 {
		return PaymentResponse{}, errors.New("invalid payment amount")
	}
	return pg.provider.ProcessPayment(request)
}

// Main function demonstrating usage
func main() {
	// Define a payment request
	payment := PaymentRequest{
		Amount:   100.0,
		Currency: "USD",
	}

	// Process payment with PayPal
	payPalGateway := NewPaymentGateway(PayPalProvider{})
	response, err := payPalGateway.Process(payment)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Response:", response.Message)
	}

	// Process payment with Stripe
	stripeGateway := NewPaymentGateway(StripeProvider{})
	response, err = stripeGateway.Process(payment)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Response:", response.Message)
	}

	// Future providers can be added by implementing PaymentProvider interface
}
