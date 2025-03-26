package main

import "fmt"

// --------------------------------------------------------------------
// Monolithic Interface (violates ISP)

type RestaurantEmployee1 interface {
	washDishes() string
	serveCustomers() string
	cookFood() string
}

type Waiter1 struct {
	food string
}

func (w *Waiter1) washDishes() string {
	// Not my job, but forced to implement because of monolithic interface
	return "Waiter1: I do not wash dishes."
}

func (w *Waiter1) serveCustomers() string {
	// My job: serving customers
	return "Waiter1 serves: " + w.food
}

func (w *Waiter1) cookFood() string {
	// Not my job
	return "Waiter1: I do not cook food."
}

// --------------------------------------------------------------------
// Segregated Interfaces (following ISP)

type Waiter interface {
	serveCustomers() string
	takeOrder() string
}

type Chef interface {
	cookFood() string
}

type frontdesk struct {
	food string
}

func (f *frontdesk) serveCustomers() string {
	// This is my job: serving customers.
	return "Frontdesk serves: " + f.food
}

func (f *frontdesk) takeOrder() string {
	// This is my job: taking orders.
	return "Frontdesk takes order for: " + f.food
}

type backdesk struct {
	food string
}

func (b *backdesk) cookFood() string {
	// This is my job: cooking food.
	return "Backdesk cooks: " + b.food
}

// --------------------------------------------------------------------
// Main function to demonstrate the usage

func main() {
	// Demonstrate the monolithic interface usage (not ideal)
	fmt.Println("----- RestaurantEmployee1 demonstration -----")
	waiter1 := &Waiter1{food: "Pasta"}
	fmt.Println(waiter1.washDishes())
	fmt.Println(waiter1.serveCustomers())
	fmt.Println(waiter1.cookFood())

	// Demonstrate the segregated interfaces usage (ideal design)
	fmt.Println("\n----- Interface Segregation demonstration -----")
	
	// Frontdesk acts as a Waiter
	var w Waiter = &frontdesk{food: "Salad"}
	// Backdesk acts as a Chef
	var c Chef = &backdesk{food: "Steak"}
	
	fmt.Println(w.serveCustomers())
	fmt.Println(w.takeOrder())
	fmt.Println(c.cookFood())
}
