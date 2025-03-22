package main

import (
	"fmt"
	"sync"
	"time"
)

// Observer Interface (Subscriber Services)
type Observer interface {
	Update(news string, wg *sync.WaitGroup)
}

// Subject Interface (News Publisher)
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(news string)
}

// NewsPublisher (Subject Implementation)
type NewsPublisher struct {
	observers sync.Map // Thread-safe observer list
}

// RegisterObserver - Adds a new observer
func (p *NewsPublisher) RegisterObserver(observer Observer) {
	p.observers.Store(observer, true)
}

// RemoveObserver - Removes an existing observer
func (p *NewsPublisher) RemoveObserver(observer Observer) {
	p.observers.Delete(observer)
}

// NotifyObservers - Notifies all observers about breaking news concurrently
func (p *NewsPublisher) NotifyObservers(news string) {
	var wg sync.WaitGroup

	p.observers.Range(func(key, value interface{}) bool {
		observer := key.(Observer)
		wg.Add(1)
		go observer.Update(news, &wg) // Notify concurrently
		return true
	})

	wg.Wait() // Ensures all Goroutines finish before exiting
	fmt.Println("✅ All subscribers have been notified successfully.\n")
}

// MobileAppService (Observer Implementation)
type MobileAppService struct{}

func (m *MobileAppService) Update(news string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second) // Simulate notification delay
	fmt.Printf("[Mobile App] Breaking News Alert: %s\n", news)
}

// EmailService (Observer Implementation)
type EmailService struct{}

func (e *EmailService) Update(news string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second) // Simulate email delay
	fmt.Printf("[Email Service] Breaking News Email Sent: %s\n", news)
}

// AggregatorService (Observer Implementation)
type AggregatorService struct{}

func (a *AggregatorService) Update(news string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second) // Simulate aggregator delay
	fmt.Printf("[News Aggregator] News Data Updated: %s\n", news)
}

func main() {
	// Create News Publisher
	newsPublisher := &NewsPublisher{}

	// Create Observer Services
	mobileApp := &MobileAppService{}
	emailService := &EmailService{}
	aggregatorService := &AggregatorService{}

	// Register Observers
	newsPublisher.RegisterObserver(mobileApp)
	newsPublisher.RegisterObserver(emailService)
	newsPublisher.RegisterObserver(aggregatorService)

	// Publish Breaking News Alerts
	newsPublisher.NotifyObservers("🔥 Google launches new AI-powered search engine!")
	newsPublisher.NotifyObservers("🚀 Google Cloud announces new security features!")
}
