package main

import "fmt"

// Notification defines the interface for sending notifications.
type Notification interface {
	Send(message string) error
}

// EmailNotification implements the Notification interface for emails.
type EmailNotification struct{}

func (e EmailNotification) Send(message string) error {
	fmt.Println("Sending Email:", message)
	return nil
}

// SMSNotification implements the Notification interface for SMS.
type SMSNotification struct{}

func (s SMSNotification) Send(message string) error {
	fmt.Println("Sending SMS:", message)
	return nil
}

// PushNotification implements the Notification interface for push notifications.
type PushNotification struct{}

func (p PushNotification) Send(message string) error {
	fmt.Println("Sending Push Notification:", message)
	return nil
}

// NotificationService uses a Notification interface to send messages.
// It is open for extension (new channels) but closed for modification.
type NotificationService struct {
	notification Notification
}

func NewNotificationService(n Notification) *NotificationService {
	return &NotificationService{notification: n}
}

func (ns *NotificationService) Notify(message string) error {
	return ns.notification.Send(message)
}

func main() {
	// Use Email Notification
	emailService := NewNotificationService(EmailNotification{})
	emailService.Notify("Hello via Email!")

	// Use SMS Notification
	smsService := NewNotificationService(SMSNotification{})
	smsService.Notify("Hello via SMS!")

	// Use Push Notification
	pushService := NewNotificationService(PushNotification{})
	pushService.Notify("Hello via Push Notification!")
}
