package main

import (
	"fmt"
	"vending-machine/internal/vendingmachine"
)

func main() {
	Run()
}

// Run demonstrates the Vending Machine functionality.
func Run() {
	vm := vendingmachine.NewVendingMachine()

	coke := vendingmachine.NewProduct("Coke", 1.5)
	pepsi := vendingmachine.NewProduct("Pepsi", 1.5)
	water := vendingmachine.NewProduct("Water", 1.0)

	vm.Inventory.AddProduct(coke, 5)
	vm.Inventory.AddProduct(pepsi, 3)
	vm.Inventory.AddProduct(water, 2)

	fmt.Println("Starting Vending Machine...")

	fmt.Println("\nSelecting Coke")
	vm.SelectProduct(coke)

	fmt.Println("Inserting coins")
	vm.InsertCoin(vendingmachine.QUARTER)
	vm.InsertCoin(vendingmachine.QUARTER)
	vm.InsertCoin(vendingmachine.QUARTER)
	vm.InsertCoin(vendingmachine.QUARTER)

	fmt.Println("\nDispensing product")
	vm.DispenseProduct()

	fmt.Println("\nReturning change")
	vm.ReturnChange()

	fmt.Println("\nSelecting Pepsi with insufficient funds")
	vm.SelectProduct(pepsi)
	vm.InsertCoin(vendingmachine.QUARTER)

	fmt.Println("\nTrying to dispense Pepsi")
	vm.DispenseProduct()

	fmt.Println("\nAdding more coins for Pepsi")
	vm.InsertCoin(vendingmachine.HALF)
	vm.InsertCoin(vendingmachine.QUARTER)
	vm.InsertCoin(vendingmachine.QUARTER)
	vm.InsertCoin(vendingmachine.QUARTER)

	fmt.Println("\nDispensing product")
	vm.DispenseProduct()

	fmt.Println("\nReturning change")
	vm.ReturnChange()
}
