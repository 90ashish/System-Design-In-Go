package main

import (
	"fmt"
	"sync"
)

// Observer Interface (Common Interface for Services)
type Observer interface {
	Update(product string, stock int)
}

// Subject Interface (Inventory System)
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(product string, stock int)
}

// Inventory (Subject Implementation)
type Inventory struct {
	observers sync.Map // Concurrent-safe map for observer management
}

// RegisterObserver - Adds a new observer
func (i *Inventory) RegisterObserver(observer Observer) {
	i.observers.Store(observer, true)
}

// RemoveObserver - Removes an existing observer
func (i *Inventory) RemoveObserver(observer Observer) {
	i.observers.Delete(observer)
}

// NotifyObservers - Notifies all observers when stock updates
func (i *Inventory) NotifyObservers(product string, stock int) {
	i.observers.Range(func(key, value interface{}) bool {
		observer := key.(Observer)
		observer.Update(product, stock)
		return true
	})
}

// UpdateStock - Updates stock levels and notifies observers
func (i *Inventory) UpdateStock(product string, stock int) {
	fmt.Printf("\n📦 Stock Update: Product '%s' updated to %d units.\n", product, stock)
	i.NotifyObservers(product, stock)
}

// WarehouseService (Observer Implementation)
type WarehouseService struct{}

func (w *WarehouseService) Update(product string, stock int) {
	fmt.Printf("[Warehouse Service] Updated inventory for '%s' to %d units.\n", product, stock)
}

// UserNotificationService (Observer Implementation)
type UserNotificationService struct{}

func (u *UserNotificationService) Update(product string, stock int) {
	if stock > 0 {
		fmt.Printf("[User Notification] Alert: '%s' is back in stock with %d units!\n", product, stock)
	}
}

// AnalyticsService (Observer Implementation)
type AnalyticsService struct{}

func (a *AnalyticsService) Update(product string, stock int) {
	fmt.Printf("[Analytics Service] Tracking sales trends for '%s'. Stock level is now %d units.\n", product, stock)
}

func main() {
	// Create Inventory Subject
	inventory := &Inventory{}

	// Create Observer Services
	warehouse := &WarehouseService{}
	notification := &UserNotificationService{}
	analytics := &AnalyticsService{}

	// Register Observers
	inventory.RegisterObserver(warehouse)
	inventory.RegisterObserver(notification)
	inventory.RegisterObserver(analytics)

	// Stock Updates
	inventory.UpdateStock("Echo Dot", 20)
	inventory.UpdateStock("Kindle Paperwhite", 0)
	inventory.UpdateStock("Fire TV Stick", 50)
}
