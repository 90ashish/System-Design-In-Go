package parking

import (
	"errors"
	"parking-lot/internal/models"
	"sync"
)

// ParkingSpot represents one spot
type ParkingSpot struct {
	ID       string
	SpotType models.VehicleType

	mu       sync.Mutex
	occupied bool
	vehicle  models.Vehicle
}

func (ps *ParkingSpot) IsAvailable() bool {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	return !ps.occupied
}

func (ps *ParkingSpot) Park(v models.Vehicle) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	if ps.occupied {
		return errors.New("spot occupied")
	}
	if v.Type() != ps.SpotType {
		return errors.New("vehicle-type mismatch")
	}
	ps.vehicle = v
	ps.occupied = true
	return nil
}

func (ps *ParkingSpot) Unpark() (models.Vehicle, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	if !ps.occupied {
		return nil, errors.New("spot empty")
	}
	v := ps.vehicle
	ps.vehicle = nil
	ps.occupied = false
	return v, nil
}
