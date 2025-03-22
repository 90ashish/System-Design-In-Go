package main

import (
	"fmt"
	"time"
)

// Service Interface
type Service interface {
	Execute(data string) error
	Compensate(data string)
}

// PaymentService - Handles payments
type PaymentService struct{}

func (p *PaymentService) Execute(data string) error {
	fmt.Println("[Payment Service] Payment processed successfully for:", data)
	return nil
}

func (p *PaymentService) Compensate(data string) {
	fmt.Println("[Payment Service] Payment rollback initiated for:", data)
}

// InventoryService - Handles inventory reservation
type InventoryService struct {
	ExecuteFunc func(data string) error
}

func (i *InventoryService) Execute(data string) error {
	if i.ExecuteFunc != nil {
		return i.ExecuteFunc(data)
	}
	fmt.Println("[Inventory Service] Inventory reserved for:", data)
	return nil
}

func (i *InventoryService) Compensate(data string) {
	fmt.Println("[Inventory Service] Inventory release initiated for:", data)
}

// ShippingService - Handles shipping
type ShippingService struct{}

func (s *ShippingService) Execute(data string) error {
	fmt.Println("[Shipping Service] Shipment arranged for:", data)
	return nil
}

func (s *ShippingService) Compensate(data string) {
	fmt.Println("[Shipping Service] Shipment cancellation initiated for:", data)
}

// NotificationService - Handles customer notifications
type NotificationService struct{}

func (n *NotificationService) Execute(data string) error {
	fmt.Println("[Notification Service] Customer notified for:", data)
	return nil
}

func (n *NotificationService) Compensate(data string) {
	fmt.Println("[Notification Service] Notification rollback initiated for:", data)
}

// Saga Orchestrator - Manages the Saga Workflow
type SagaOrchestrator struct {
	steps         []Service
	compensations []Service
}

func (so *SagaOrchestrator) AddStep(service Service) {
	so.steps = append(so.steps, service)
}

func (so *SagaOrchestrator) ExecuteSaga(data string) {
	fmt.Println("\nStarting Order Fulfillment Process...")

	for _, step := range so.steps {
		err := step.Execute(data)
		if err != nil {
			fmt.Println("[ERROR] Transaction failed. Initiating rollback...")
			so.Rollback(data)
			return
		}
		so.compensations = append([]Service{step}, so.compensations...) // Collect compensations
	}

	fmt.Println("✅ Order Fulfillment Process Completed Successfully!")
}

func (so *SagaOrchestrator) Rollback(data string) {
	for _, step := range so.compensations {
		step.Compensate(data)
	}
	fmt.Println("❌ Order Fulfillment Rolled Back Successfully!")
}

func main() {
	// Initialize Services
	payment := &PaymentService{}
	// For InventoryService, you can optionally override ExecuteFunc
	inventory := &InventoryService{
		ExecuteFunc: nil, // Default behavior
	}
	shipping := &ShippingService{}
	notification := &NotificationService{}

	// Setup Saga Orchestrator
	orchestrator := &SagaOrchestrator{}
	orchestrator.AddStep(payment)
	orchestrator.AddStep(inventory)
	orchestrator.AddStep(shipping)
	orchestrator.AddStep(notification)

	// Simulate Successful Transaction
	fmt.Println(">>> Processing Successful Order")
	orchestrator.ExecuteSaga("Order#123")

	time.Sleep(2 * time.Second)

	// Simulate Failed Transaction by overriding InventoryService's ExecuteFunc
	inventory.ExecuteFunc = func(data string) error {
		fmt.Println("[Inventory Service] Inventory reservation FAILED for:", data)
		return fmt.Errorf("inventory error")
	}
	fmt.Println("\n>>> Processing Failed Order")
	orchestrator.ExecuteSaga("Order#456")
}
