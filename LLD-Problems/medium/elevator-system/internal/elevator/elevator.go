package elevator

import (
	"fmt"
	"sync"
	"time"
)

type Elevator struct {
	id               int
	capacity         int
	currentFloor     int
	currentDirection Direction
	requests         chan *Request
	stopChan         chan struct{}
	mu               sync.RWMutex
}

func NewElevator(id, capacity int) *Elevator {
	return &Elevator{
		id:               id,
		capacity:         capacity,
		currentFloor:     1,
		currentDirection: DirectionUp,
		requests:         make(chan *Request, capacity),
		stopChan:         make(chan struct{}),
	}
}

func (e *Elevator) AddRequest(request *Request) bool {
	select {
	case e.requests <- request:
		fmt.Printf("Elevator %d added request: Floor %d to %d\n", e.id, request.SourceFloor, request.DestinationFloor)
		return true
	default:
		return false
	}
}

func (e *Elevator) getCurrentFloor() int {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.currentFloor
}

func (e *Elevator) GetCurrentDirection() Direction {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.currentDirection
}

func (e *Elevator) setCurrentFloor(floor int) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.currentFloor = floor
}

func (e *Elevator) Run() {
	go func() {
		for {
			select {
			case req := <-e.requests:
				e.processRequest(req)
			case <-e.stopChan:
				return
			}
		}
	}()
}

func (e *Elevator) Stop() {
	close(e.stopChan)
}

func (e *Elevator) processRequest(request *Request) {
	start := e.getCurrentFloor()
	end := request.DestinationFloor

	if start < end {
		e.currentDirection = DirectionUp
		for f := start; f <= end; f++ {
			e.setCurrentFloor(f)
			fmt.Printf("Elevator %d reached floor %d\n", e.id, f)
			time.Sleep(time.Second)
		}
	} else {
		e.currentDirection = DirectionDown
		for f := start; f >= end; f-- {
			e.setCurrentFloor(f)
			fmt.Printf("Elevator %d reached floor %d\n", e.id, f)
			time.Sleep(time.Second)
		}
	}
}
