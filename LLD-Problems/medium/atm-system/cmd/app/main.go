package main

import (
	"atm-system/internal/atmsystem"
)

// Entry point: composes all components via Dependency Injection (DIP)
// and starts the ATM (Facade pattern).
func main() {
	// Factory methods to create concrete implementations (Factory Pattern)
	ui := atmsystem.NewConsoleUI()
	cr := atmsystem.NewConsoleCardReader(ui)
	disp := atmsystem.NewConsoleDispenser(ui)
	backend := atmsystem.NewMockBackend()

	// Inject dependencies into ATM (DIP)
	atm := atmsystem.NewATM(cr, disp, backend, ui)
	atm.Run() // Start the main loop
}
