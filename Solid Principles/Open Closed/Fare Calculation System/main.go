package main

import (
	"fmt"
)

// Ride represents details of a ride needed for fare calculation.
type Ride struct {
	Distance        float64 // in kilometers
	Duration        float64 // in minutes
	SurgeMultiplier float64 // e.g., 1.0 (normal), 1.5 (surge)
}

// ------------------------------
// Pricing Strategy Interfaces
// ------------------------------

// AdditivePricingStrategy defines an interface for fare components that are summed.
type AdditivePricingStrategy interface {
	Calculate(ride Ride) float64
}

// MultiplicativePricingStrategy defines an interface for fare modifiers applied multiplicatively.
type MultiplicativePricingStrategy interface {
	Multiply(ride Ride) float64
}

// -----------------------------------------
// Concrete Additive Pricing Strategies
// -----------------------------------------

// BaseFareStrategy calculates a fixed base fare.
type BaseFareStrategy struct {
	BaseFare float64
}

func (b BaseFareStrategy) Calculate(ride Ride) float64 {
	return b.BaseFare
}

// DistanceFareStrategy calculates fare based on the ride distance.
type DistanceFareStrategy struct {
	RatePerKm float64
}

func (d DistanceFareStrategy) Calculate(ride Ride) float64 {
	return ride.Distance * d.RatePerKm
}

// TimeFareStrategy calculates fare based on ride duration.
type TimeFareStrategy struct {
	RatePerMinute float64
}

func (t TimeFareStrategy) Calculate(ride Ride) float64 {
	return ride.Duration * t.RatePerMinute
}

// ---------------------------------------------
// Concrete Multiplicative Pricing Strategy
// ---------------------------------------------

// SurgePricingStrategy applies a surge multiplier to the fare.
type SurgePricingStrategy struct{}

func (s SurgePricingStrategy) Multiply(ride Ride) float64 {
	if ride.SurgeMultiplier > 1.0 {
		return ride.SurgeMultiplier
	}
	return 1.0
}

// ---------------------------------------------
// FareCalculator Composing Strategies
// ---------------------------------------------

// FareCalculator computes the final fare by combining additive and multiplicative strategies.
type FareCalculator struct {
	additiveStrategies     []AdditivePricingStrategy
	multiplicativeStrategies []MultiplicativePricingStrategy
}

func NewFareCalculator(additive []AdditivePricingStrategy, multiplier []MultiplicativePricingStrategy) *FareCalculator {
	return &FareCalculator{
		additiveStrategies:     additive,
		multiplicativeStrategies: multiplier,
	}
}

func (f *FareCalculator) CalculateFare(ride Ride) float64 {
	// Sum all additive pricing components.
	baseFare := 0.0
	for _, strategy := range f.additiveStrategies {
		baseFare += strategy.Calculate(ride)
	}

	// Apply all multiplicative modifiers.
	multiplier := 1.0
	for _, strategy := range f.multiplicativeStrategies {
		multiplier *= strategy.Multiply(ride)
	}

	return baseFare * multiplier
}

// ---------------------------------------------
// Main Function: Demonstration
// ---------------------------------------------
func main() {
	// Define additive pricing strategies.
	baseFare := BaseFareStrategy{BaseFare: 2.0}             // Flat base fare of $2.00
	distanceFare := DistanceFareStrategy{RatePerKm: 1.5}      // $1.50 per kilometer
	timeFare := TimeFareStrategy{RatePerMinute: 0.5}          // $0.50 per minute

	// Define multiplicative pricing strategy.
	surge := SurgePricingStrategy{}

	// Create a fare calculator with the current pricing strategies.
	fareCalculator := NewFareCalculator(
		[]AdditivePricingStrategy{baseFare, distanceFare, timeFare},
		[]MultiplicativePricingStrategy{surge},
	)

	// Example ride details.
	ride := Ride{
		Distance:        10,    // 10 km
		Duration:        15,    // 15 minutes
		SurgeMultiplier: 1.5,   // 50% surge pricing
	}

	// Calculate the final fare.
	finalFare := fareCalculator.CalculateFare(ride)
	fmt.Printf("Final Fare for the ride: $%.2f\n", finalFare)
}
