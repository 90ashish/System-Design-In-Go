package parking

import (
	"errors"
	"parking-lot/internal/models"
	"sync"
)

// ParkingLot is a singleton managing all levels
type ParkingLot struct {
	levels []*Level
	mu     sync.Mutex
}

var (
	instance *ParkingLot
	once     sync.Once
)

// GetParkingLot initializes on first call
func GetParkingLot(numLevels, spotsPerType int) *ParkingLot {
	once.Do(func() {
		pl := &ParkingLot{}
		for i := 1; i <= numLevels; i++ {
			pl.levels = append(pl.levels, NewLevel(i, spotsPerType))
		}
		instance = pl
	})
	return instance
}

// ParkVehicle tries each level in order
func (pl *ParkingLot) ParkVehicle(v models.Vehicle) (*models.Ticket, error) {
	pl.mu.Lock()
	defer pl.mu.Unlock()
	for _, lvl := range pl.levels {
		if t, err := lvl.ParkVehicle(v); err == nil {
			return t, nil
		}
	}
	return nil, errors.New("parking lot full")
}

// UnparkVehicle releases based on ticket
func (pl *ParkingLot) UnparkVehicle(t *models.Ticket) error {
	pl.mu.Lock()
	defer pl.mu.Unlock()
	idx := t.Level - 1
	if idx < 0 || idx >= len(pl.levels) {
		return errors.New("invalid level")
	}
	return pl.levels[idx].UnparkVehicle(t)
}

// GetAvailability returns free‐spot counts per level/type
func (pl *ParkingLot) GetAvailability() map[int]map[models.VehicleType]int {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	out := make(map[int]map[models.VehicleType]int, len(pl.levels))
	for _, lvl := range pl.levels {
		lvl.mu.Lock()
		counts := make(map[models.VehicleType]int, len(lvl.available))
		for vt, c := range lvl.available {
			counts[vt] = c
		}
		lvl.mu.Unlock()
		out[lvl.Number] = counts
	}
	return out
}
