package main

import (
	"fmt"
	"parking-lot/internal/models"
	"parking-lot/internal/parking"
)

func main() {
	// initialize 3 levels, 5 spots of each type per level
	lot := parking.GetParkingLot(3, 5)

	vehicles := []models.Vehicle{
		&models.Car{License: "CAR-001"},
		&models.Car{License: "CAR-002"},
		&models.Car{License: "CAR-003"},
		&models.Motorcycle{License: "BIKE-001"},
		&models.Truck{License: "TRUCK-001"},
		&models.Truck{License: "TRUCK-002"},
	}

	var tickets []*models.Ticket
	for _, v := range vehicles {
		t, err := lot.ParkVehicle(v)
		if err != nil {
			fmt.Println("Park error:", err)
			continue
		}
		fmt.Printf("Parked %s [%s] @ L%d/%s (Ticket %s)\n",
			v.Type(), v.LicensePlate(), t.Level, t.SpotID, t.ID)
		tickets = append(tickets, t)
	}

	fmt.Println("Availability:", lot.GetAvailability())

	// unpark first vehicle
	if len(tickets) > 0 {
		if err := lot.UnparkVehicle(tickets[0]); err != nil {
			fmt.Println("Unpark error:", err)
		} else {
			fmt.Println("Unparked:", tickets[0].Vehicle.LicensePlate())
		}
		fmt.Println("After unpark:", lot.GetAvailability())
	}
}
