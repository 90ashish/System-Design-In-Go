package elevator

import (
	"math"
	"sync"
)

type ElevatorController struct {
	elevators []*Elevator
	mu        sync.RWMutex
}

func NewElevatorController(numElevators, capacity int) *ElevatorController {
	ec := &ElevatorController{
		elevators: make([]*Elevator, numElevators),
	}
	for i := range ec.elevators {
		ev := NewElevator(i+1, capacity)
		ec.elevators[i] = ev
		ev.Run()
	}
	return ec
}

func (ec *ElevatorController) RequestElevator(source, dest int) {
	ev := ec.findOptimalElevator(source, dest)
	req := NewRequest(source, dest)
	ev.AddRequest(req)
}

func (ec *ElevatorController) findOptimalElevator(source, dest int) *Elevator {
	ec.mu.RLock()
	defer ec.mu.RUnlock()

	reqDir := getDirection(source, dest)
	var best *Elevator
	bestScore := math.MaxInt32

	for _, ev := range ec.elevators {
		floor := ev.getCurrentFloor()
		dir := ev.GetCurrentDirection()
		qlen := len(ev.requests)
		var score int

		switch {
		// best: moving in same dir and will pass through source
		case dir == reqDir && ((dir == DirectionUp && floor <= source) || (dir == DirectionDown && floor >= source)):
			score = abs(floor - source)
		// idle elevator
		case qlen == 0:
			score = abs(floor - source)
		// moving same direction but won't pass
		case dir == reqDir:
			score = 1000 + abs(floor-source)
		// moving opposite or busy
		default:
			score = 2000 + abs(floor-source)
		}

		if score < bestScore {
			bestScore = score
			best = ev
		}
	}
	return best
}

func (ec *ElevatorController) Stop() {
	ec.mu.Lock()
	defer ec.mu.Unlock()
	for _, ev := range ec.elevators {
		ev.Stop()
	}
}

// utils

func getDirection(from, to int) Direction {
	if to > from {
		return DirectionUp
	}
	return DirectionDown
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
