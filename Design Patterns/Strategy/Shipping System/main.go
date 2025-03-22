package main

import (
	"fmt"
)

// ShippingStrategy Interface
type ShippingStrategy interface {
	CalculateCost(weight float64) float64
}

// StandardShipping Strategy
type StandardShipping struct{}

func (s *StandardShipping) CalculateCost(weight float64) float64 {
	return weight * 5.0 // $5 per kg
}

// ExpressShipping Strategy
type ExpressShipping struct{}

func (e *ExpressShipping) CalculateCost(weight float64) float64 {
	return weight * 8.0 // $8 per kg
}

// SameDayShipping Strategy
type SameDayShipping struct{}

func (s *SameDayShipping) CalculateCost(weight float64) float64 {
	return weight * 15.0 // $15 per kg
}

// Context - Shipping Cost Calculator
type ShippingCostCalculator struct {
	strategy ShippingStrategy
}

// SetStrategy - Assigns the chosen strategy
func (c *ShippingCostCalculator) SetStrategy(strategy ShippingStrategy) {
	c.strategy = strategy
}

// CalculateShippingCost - Executes the chosen strategy
func (c *ShippingCostCalculator) CalculateShippingCost(weight float64) {
	if c.strategy == nil {
		fmt.Println("Error: No shipping strategy selected!")
		return
	}
	cost := c.strategy.CalculateCost(weight)
	fmt.Printf("🚚 Shipping Cost for %.2f kg: $%.2f\n", weight, cost)
}

func main() {
	// Initialize the context (Shipping Cost Calculator)
	calculator := &ShippingCostCalculator{}

	// Example 1: Standard Shipping
	fmt.Println("🔹 Standard Shipping:")
	calculator.SetStrategy(&StandardShipping{})
	calculator.CalculateShippingCost(10.5) // Example weight: 10.5 kg

	// Example 2: Express Shipping
	fmt.Println("🔹 Express Shipping:")
	calculator.SetStrategy(&ExpressShipping{})
	calculator.CalculateShippingCost(5.0) // Example weight: 5 kg

	// Example 3: Same-Day Shipping
	fmt.Println("🔹 Same-Day Shipping:")
	calculator.SetStrategy(&SameDayShipping{})
	calculator.CalculateShippingCost(2.0) // Example weight: 2 kg
}
