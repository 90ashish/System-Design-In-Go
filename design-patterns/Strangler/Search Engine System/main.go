package main

import (
	"fmt"
	"math/rand"
	"time"
)

// LegacySearch - Represents the legacy search system
type LegacySearch struct{}

func (l *LegacySearch) Search(query string) string {
	return fmt.Sprintf("[Legacy Search] Results for: '%s' (old system)", query)
}

// ModernSearch - Represents the modernized search system
type ModernSearch struct{}

func (m *ModernSearch) Search(query string) string {
	return fmt.Sprintf("[Modern Search] Results for: '%s' (new system)", query)
}

// SearchFacade - Acts as a proxy to gradually switch between systems
type SearchFacade struct {
	legacySystem *LegacySearch
	modernSystem *ModernSearch
	trafficSplit int // Percentage of traffic diverted to the modern system
}

// NewSearchFacade - Initializes the SearchFacade with defined traffic split
func NewSearchFacade(trafficSplit int) *SearchFacade {
	return &SearchFacade{
		legacySystem: &LegacySearch{},
		modernSystem: &ModernSearch{},
		trafficSplit: trafficSplit,
	}
}

// Search - Routes requests based on the defined traffic split
func (f *SearchFacade) Search(query string) string {
	if rand.Intn(100) < f.trafficSplit {
		return f.modernSystem.Search(query) // Divert to Modern System
	}
	return f.legacySystem.Search(query) // Default to Legacy System
}

// GradualTrafficShift - Simulates gradual migration of traffic
func GradualTrafficShift(facade *SearchFacade) {
	steps := []int{10, 30, 60, 90, 100} // Gradual migration percentages

	for _, split := range steps {
		fmt.Printf("\n🔄 Shifting %d%% of traffic to the modern system...\n", split)
		facade.trafficSplit = split

		// Simulate multiple search requests
		for i := 0; i < 5; i++ {
			fmt.Println(facade.Search(fmt.Sprintf("Search Query %d", i+1)))
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Random seed for simulating traffic split
	searchFacade := NewSearchFacade(0)

	// Gradual Traffic Shift Simulation
	GradualTrafficShift(searchFacade)

	fmt.Println("✅ Migration to the new system completed successfully.")
}
