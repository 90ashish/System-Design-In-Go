package atmsystem

import "sync"

// Account represents a bank account. Thread-safe via embedded mutex.
type Account struct {
	ID      string
	PIN     string
	Balance int
	mu      sync.RWMutex
}

// NewAccount is a simple factory for creating accounts.
func NewAccount(id, pin string, balance int) *Account {
	return &Account{ID: id, PIN: pin, Balance: balance}
}
