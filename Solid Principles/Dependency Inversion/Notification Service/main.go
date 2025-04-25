package main

import (
	"fmt"
)

// -----------------------------
// Domain Model
// -----------------------------
type Order struct {
	ID     string
	Status string
}

// -----------------------------
// Abstraction (Notifier interface)
// -----------------------------
type Notifier interface {
	Notify(order Order) error
}

// -----------------------------
// Low-Level Modules (Concrete Notifiers)
// -----------------------------

// EmailNotifier sends notifications via email.
type EmailNotifier struct{}

func (e EmailNotifier) Notify(order Order) error {
	fmt.Printf("[Email] Order %s status: %s\n", order.ID, order.Status)
	return nil
}

// SMSNotifier sends notifications via SMS.
type SMSNotifier struct{}

func (s SMSNotifier) Notify(order Order) error {
	fmt.Printf("[SMS] Order %s status: %s\n", order.ID, order.Status)
	return nil
}

// PushNotifier sends notifications via push messages.
type PushNotifier struct{}

func (p PushNotifier) Notify(order Order) error {
	fmt.Printf("[Push] Order %s status: %s\n", order.ID, order.Status)
	return nil
}

// -----------------------------
// High-Level Module (OrderService)
// -----------------------------
type OrderService struct {
	notifier Notifier
}

// NewOrderService injects a Notifier into OrderService.
func NewOrderService(n Notifier) *OrderService {
	return &OrderService{notifier: n}
}

// UpdateStatus updates the order's status and triggers a notification.
func (s *OrderService) UpdateStatus(order Order, newStatus string) {
	// (1) Update order status (e.g. in DB) — simulated here:
	order.Status = newStatus
	fmt.Printf("Order %s updated to status: %s\n", order.ID, order.Status)

	// (2) Notify customer via injected Notifier
	if err := s.notifier.Notify(order); err != nil {
		fmt.Println("Notification error:", err)
	}
}

func main() {
	order := Order{ID: "A123", Status: "Processing"}

	// Inject EmailNotifier
	emailService := NewOrderService(EmailNotifier{})
	emailService.UpdateStatus(order, "Shipped")

	// Inject SMSNotifier
	smsService := NewOrderService(SMSNotifier{})
	smsService.UpdateStatus(order, "Out for Delivery")

	// Inject PushNotifier
	pushService := NewOrderService(PushNotifier{})
	pushService.UpdateStatus(order, "Delivered")

	// ➤ To add a new channel (e.g., SlackNotifier), simply:
	//   1. Implement `Notifier`:
	//        type SlackNotifier struct { /*…*/ }
	//        func (s SlackNotifier) Notify(o Order) error { /*…*/ }
	//   2. Inject it: NewOrderService(SlackNotifier{}).UpdateStatus(…)
	// No changes to OrderService are required, satisfying DIP.
}
