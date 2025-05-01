package atmsystem

import "fmt"

// CashDispenser abstracts how cash is dispensed (Strategy Pattern).
type CashDispenser interface {
	Dispense(amount int) error
}

// ConsoleDispenser prints to console instead of real hardware.
type ConsoleDispenser struct {
	ui UI
}

// NewConsoleDispenser builds a dispenser (Factory Pattern).
func NewConsoleDispenser(ui UI) CashDispenser {
	return &ConsoleDispenser{ui}
}

// Dispense simulates giving out cash.
func (d *ConsoleDispenser) Dispense(amount int) error {
	d.ui.Display(fmt.Sprintf("Dispensing cash: %d", amount))
	return nil
}
