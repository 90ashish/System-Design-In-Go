package atmsystem

import (
	"errors"
	"sync"
)

// BankBackend defines account operations (Repository Pattern).
type BankBackend interface {
	Validate(cardNum, pin string) (*Account, error)
	GetBalance(accountID string) (int, error)
	Withdraw(accountID string, amount int) error
	Deposit(accountID string, amount int) error
}

// MockBackend is an in-memory bank backend (for testing/demo).
// Uses mutex for concurrent safety.
type MockBackend struct {
	accounts map[string]*Account
	mu       sync.RWMutex
}

// NewMockBackend constructs the mock with some demo accounts.
func NewMockBackend() BankBackend {
	return &MockBackend{
		accounts: map[string]*Account{
			"1234": NewAccount("1234", "4321", 1000),
			"5678": NewAccount("5678", "8765", 500),
		},
	}
}

// Validate checks card number and PIN.
func (mb *MockBackend) Validate(cardNum, pin string) (*Account, error) {
	mb.mu.RLock()
	defer mb.mu.RUnlock()
	acct, ok := mb.accounts[cardNum]
	if !ok || acct.PIN != pin {
		return nil, errors.New("authentication failed")
	}
	return acct, nil
}

// GetBalance returns the current balance.
func (mb *MockBackend) GetBalance(accountID string) (int, error) {
	mb.mu.RLock()
	acct, ok := mb.accounts[accountID]
	mb.mu.RUnlock()
	if !ok {
		return 0, errors.New("account not found")
	}
	acct.mu.RLock()
	defer acct.mu.RUnlock()
	return acct.Balance, nil
}

// Withdraw deducts amount if sufficient funds exist.
func (mb *MockBackend) Withdraw(accountID string, amount int) error {
	mb.mu.RLock()
	acct, ok := mb.accounts[accountID]
	mb.mu.RUnlock()
	if !ok {
		return errors.New("account not found")
	}
	acct.mu.Lock()
	defer acct.mu.Unlock()
	if acct.Balance < amount {
		return errors.New("insufficient funds")
	}
	acct.Balance -= amount
	return nil
}

// Deposit adds the given amount.
func (mb *MockBackend) Deposit(accountID string, amount int) error {
	mb.mu.RLock()
	acct, ok := mb.accounts[accountID]
	mb.mu.RUnlock()
	if !ok {
		return errors.New("account not found")
	}
	acct.mu.Lock()
	defer acct.mu.Unlock()
	acct.Balance += amount
	return nil
}
