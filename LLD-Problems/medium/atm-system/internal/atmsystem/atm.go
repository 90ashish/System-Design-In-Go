package atmsystem

import (
	"fmt"
	"strconv"
)

// ATM orchestrates all components (Facade Pattern).
// Performs authentication, sessions, and operations.
type ATM struct {
	cardReader CardReader
	dispenser  CashDispenser
	backend    BankBackend
	ui         UI
	account    *Account
}

// NewATM wires up the ATM (DIP + Factory).
func NewATM(cr CardReader, d CashDispenser, bb BankBackend, ui UI) *ATM {
	return &ATM{cardReader: cr, dispenser: d, backend: bb, ui: ui}
}

// Run loops forever, handling card insert/authentication.
func (atm *ATM) Run() {
	for {
		atm.ui.Display("\n--- Welcome to GoATM ---")
		cardNum, err := atm.cardReader.ReadCard()
		if err != nil {
			atm.ui.Display("Failed to read card")
			continue
		}
		pin, err := atm.cardReader.ReadPIN()
		if err != nil {
			atm.ui.Display("Failed to read PIN")
			continue
		}
		acct, err := atm.backend.Validate(cardNum, pin)
		if err != nil {
			atm.ui.Display("Authentication failed")
			continue
		}
		atm.account = acct
		atm.session() // enter session loop
	}
}

// session handles user operations until exit.
func (atm *ATM) session() {
	for {
		atm.ui.Display("\nSelect operation:\n1) Balance Inquiry\n2) Cash Withdrawal\n3) Cash Deposit\n4) Exit")
		choice, err := atm.ui.Prompt("Choice: ")
		if err != nil {
			atm.ui.Display("Invalid input")
			continue
		}
		switch choice {
		case "1":
			atm.balanceInquiry()
		case "2":
			atm.cashWithdrawal()
		case "3":
			atm.cashDeposit()
		case "4":
			atm.ui.Display("Thank you. Goodbye!")
			return
		default:
			atm.ui.Display("Invalid choice")
		}
	}
}

// balanceInquiry fetches and displays the current balance.
func (atm *ATM) balanceInquiry() {
	bal, err := atm.backend.GetBalance(atm.account.ID)
	if err != nil {
		atm.ui.Display("Error retrieving balance")
		return
	}
	atm.ui.Display(fmt.Sprintf("Your balance: %d", bal))
}

// cashWithdrawal prompts for amount, validates, performs withdrawal, and dispenses cash.
func (atm *ATM) cashWithdrawal() {
	input, _ := atm.ui.Prompt("Enter amount to withdraw: ")
	amt, err := strconv.Atoi(input)
	if err != nil || amt <= 0 {
		atm.ui.Display("Invalid amount")
		return
	}
	if err := atm.backend.Withdraw(atm.account.ID, amt); err != nil {
		atm.ui.Display("Withdrawal failed: " + err.Error())
		return
	}
	atm.dispenser.Dispense(amt)                // Dispense hardware call
	atm.ui.Display("Please collect your cash") // Confirmation
}

// cashDeposit prompts for amount and updates backend.
func (atm *ATM) cashDeposit() {
	input, _ := atm.ui.Prompt("Enter amount to deposit: ")
	amt, err := strconv.Atoi(input)
	if err != nil || amt <= 0 {
		atm.ui.Display("Invalid amount")
		return
	}
	if err := atm.backend.Deposit(atm.account.ID, amt); err != nil {
		atm.ui.Display("Deposit failed: " + err.Error())
		return
	}
	atm.ui.Display("Deposit successful")
}
