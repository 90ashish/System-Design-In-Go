package models

import "time"

// Ticket holds parking info
type Ticket struct {
	ID        string
	Vehicle   Vehicle
	Level     int
	SpotID    string
	EntryTime time.Time
}
