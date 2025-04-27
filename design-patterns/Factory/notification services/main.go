package main

import (
	"fmt"
	"sync"
	"time"
)

// Notification interface (Common Interface)
type Notification interface {
	SendNotification(message string)
}

// EmailNotification struct
type EmailNotification struct{}

func (e *EmailNotification) SendNotification(message string) {
	time.Sleep(1 * time.Second) // Simulate processing delay
	fmt.Printf("[Email] Notification Sent: %s\n", message)
}

// SMSNotification struct
type SMSNotification struct{}

func (s *SMSNotification) SendNotification(message string) {
	time.Sleep(1 * time.Second) // Simulate processing delay
	fmt.Printf("[SMS] Notification Sent: %s\n", message)
}

// PushNotification struct
type PushNotification struct{}

func (p *PushNotification) SendNotification(message string) {
	time.Sleep(1 * time.Second) // Simulate processing delay
	fmt.Printf("[Push Notification] Notification Sent: %s\n", message)
}

// NotificationFactory - Factory for creating notification methods
type NotificationFactory struct{}

// CreateNotification - Factory method to generate the correct notification method
func (f *NotificationFactory) CreateNotification(method string) (Notification, error) {
	switch method {
	case "email":
		return &EmailNotification{}, nil
	case "sms":
		return &SMSNotification{}, nil
	case "push":
		return &PushNotification{}, nil
	default:
		return nil, fmt.Errorf("unsupported notification method: %s", method)
	}
}

// ProcessNotification - Uses Goroutines for concurrent notification processing
func ProcessNotification(wg *sync.WaitGroup, factory *NotificationFactory, method string, message string) {
	defer wg.Done()

	notificationMethod, err := factory.CreateNotification(method)
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		return
	}
	notificationMethod.SendNotification(message)
}

func main() {
	factory := &NotificationFactory{}
	var wg sync.WaitGroup

	// Bulk Notifications Data
	notifications := []struct {
		method  string
		message string
	}{
		{"email", "Your order has been shipped!"},
		{"sms", "Your package is arriving today."},
		{"push", "Flash Sale Alert! Don't miss out."},
		{"email", "Your payment has been processed."},
		{"sms", "We have received your order."},
		{"whatsapp", "New offer on electronics!"}, // Unsupported method
	}

	// Concurrently send notifications
	for _, note := range notifications {
		wg.Add(1)
		go ProcessNotification(&wg, factory, note.method, note.message)
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	fmt.Println("✅ All notifications have been processed successfully.")
}
