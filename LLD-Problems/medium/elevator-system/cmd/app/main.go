package main

import (
	"elevator-system/internal/elevator"
	"time"
)

func main() {
	// Passing num of elevators and capacity
	controller := elevator.NewElevatorController(2, 2)
	defer controller.Stop()

	// Simulate elevator requests source to destination
	controller.RequestElevator(2, 5)
	controller.RequestElevator(3, 6)
	controller.RequestElevator(3, 2)
	controller.RequestElevator(1, 3)
	controller.RequestElevator(2, 1)

	// Wait for elevators to process requests
	time.Sleep(30 * time.Second)
}
