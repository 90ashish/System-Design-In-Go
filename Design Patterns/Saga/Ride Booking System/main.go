package main

import (
	"errors"
	"fmt"
)

// PaymentProcessor interface for payment services.
type PaymentProcessor interface {
	ProcessPayment(rideID string, amount float64) error
	RefundPayment(rideID string, amount float64)
}

// Step 1: Book Ride Service
type RideBookingService struct{}

func (r *RideBookingService) BookRide(rideID string) error {
	fmt.Printf("[RideBookingService] Ride %s booked successfully.\n", rideID)
	return nil
}

func (r *RideBookingService) CancelRide(rideID string) {
	fmt.Printf("[RideBookingService] Booking for Ride %s canceled.\n", rideID)
}

// Step 2: Driver Assignment Service
type DriverAssignmentService struct{}

func (d *DriverAssignmentService) AssignDriver(rideID string) error {
	fmt.Printf("[DriverAssignmentService] Driver assigned for Ride %s.\n", rideID)
	return nil
}

func (d *DriverAssignmentService) RevokeDriverAssignment(rideID string) {
	fmt.Printf("[DriverAssignmentService] Driver unassigned for Ride %s.\n", rideID)
}

// Step 3: Payment Service
type PaymentService struct{}

func (p *PaymentService) ProcessPayment(rideID string, amount float64) error {
	fmt.Printf("[PaymentService] Payment of $%.2f for Ride %s processed successfully.\n", amount, rideID)
	return nil
}

func (p *PaymentService) RefundPayment(rideID string, amount float64) {
	fmt.Printf("[PaymentService] Refund of $%.2f issued for Ride %s.\n", amount, rideID)
}

// Step 4: Notification Service
type NotificationService struct{}

func (n *NotificationService) SendNotification(rideID string, message string) {
	fmt.Printf("[NotificationService] Notification Sent: %s (Ride %s)\n", message, rideID)
}

// Saga Orchestrator - Manages the sequence and rollback on failure
type SagaOrchestrator struct {
	rideBookingService      *RideBookingService
	driverAssignmentService *DriverAssignmentService
	paymentService          PaymentProcessor // use the interface here
	notificationService     *NotificationService
}

func NewSagaOrchestrator() *SagaOrchestrator {
	return &SagaOrchestrator{
		rideBookingService:      &RideBookingService{},
		driverAssignmentService: &DriverAssignmentService{},
		paymentService:          &PaymentService{}, // initial concrete implementation
		notificationService:     &NotificationService{},
	}
}

// Execute Saga Process
func (s *SagaOrchestrator) BookRideSaga(rideID string, amount float64) {
	// Step 1: Book Ride
	if err := s.rideBookingService.BookRide(rideID); err != nil {
		fmt.Println("[ERROR] Booking ride failed.")
		return
	}

	// Step 2: Assign Driver (Simulate Failure)
	if err := s.driverAssignmentService.AssignDriver(rideID); err != nil {
		fmt.Println("[ERROR] Driver assignment failed. Rolling back...")
		s.rideBookingService.CancelRide(rideID)
		return
	}

	// Step 3: Process Payment (Simulate Failure)
	if err := s.paymentService.ProcessPayment(rideID, amount); err != nil {
		fmt.Println("[ERROR] Payment processing failed. Rolling back...")
		s.driverAssignmentService.RevokeDriverAssignment(rideID)
		s.rideBookingService.CancelRide(rideID)
		return
	}

	// Step 4: Notification
	s.notificationService.SendNotification(rideID, "Your ride has been confirmed!")

	fmt.Println("✅ Ride booking completed successfully.")
}

// Simulated Failure Service for Payment
type PaymentServiceWithError struct{}

func (p *PaymentServiceWithError) ProcessPayment(rideID string, amount float64) error {
	return errors.New("[PaymentService] Payment processing failed.")
}

func (p *PaymentServiceWithError) RefundPayment(rideID string, amount float64) {
	fmt.Printf("[PaymentService] Refund of $%.2f issued for Ride %s.\n", amount, rideID)
}

func main() {
	saga := NewSagaOrchestrator()

	// Successful Scenario
	fmt.Println("\n=== Successful Ride Booking ===")
	saga.BookRideSaga("RIDE001", 20.00)

	// Failed Scenario (Simulated Payment Failure)
	fmt.Println("\n=== Failed Scenario: Payment Service Failure ===")
	saga.paymentService = &PaymentServiceWithError{} // Now valid since both implement PaymentProcessor
	saga.BookRideSaga("RIDE002", 25.00)
}
