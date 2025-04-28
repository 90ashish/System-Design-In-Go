package parking

import (
	"errors"
	"fmt"
	"parking-lot/internal/models"
	"sync"
	"sync/atomic"
	"time"
)

var ticketCount uint64

// Level holds spots for one floor
type Level struct {
	Number    int
	spots     []*ParkingSpot
	mu        sync.Mutex
	available map[models.VehicleType]int
}

func NewLevel(levelNumber, spotsPerType int) *Level {
	l := &Level{
		Number:    levelNumber,
		available: map[models.VehicleType]int{},
	}
	types := []models.VehicleType{
		models.CarType,
		models.MotorcycleType,
		models.TruckType,
	}
	for _, t := range types {
		for i := 0; i < spotsPerType; i++ {
			id := fmt.Sprintf("L%d-%s-%d", levelNumber, t.String(), i+1)
			l.spots = append(l.spots, &ParkingSpot{ID: id, SpotType: t})
		}
		l.available[t] = spotsPerType
	}
	return l
}

func (l *Level) ParkVehicle(v models.Vehicle) (*models.Ticket, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, s := range l.spots {
		if s.SpotType == v.Type() && s.IsAvailable() {
			if err := s.Park(v); err != nil {
				continue
			}
			l.available[v.Type()]--
			tid := fmt.Sprintf("%s-%d", s.ID, atomic.AddUint64(&ticketCount, 1))
			return &models.Ticket{
				ID:        tid,
				Vehicle:   v,
				Level:     l.Number,
				SpotID:    s.ID,
				EntryTime: time.Now(),
			}, nil
		}
	}
	return nil, errors.New("no available spots on level")
}

func (l *Level) UnparkVehicle(t *models.Ticket) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, s := range l.spots {
		if s.ID == t.SpotID {
			if _, err := s.Unpark(); err != nil {
				return err
			}
			l.available[s.SpotType]++
			return nil
		}
	}
	return errors.New("spot not found")
}
