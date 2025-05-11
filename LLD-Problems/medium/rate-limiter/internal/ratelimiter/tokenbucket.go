package ratelimiter

import (
	"sync"
	"time"
)

// TokenBucket implements the token‑bucket algorithm.
// ‣ Pattern: Strategy (encapsulates the “allow or not” logic)
// ‣ Factory: NewTokenBucket hides setup details
type TokenBucket struct {
	capacity   float64    // max tokens
	refillRate float64    // tokens added per second
	tokens     float64    // current token count
	last       time.Time  // last refill timestamp
	mu         sync.Mutex // guards all fields
}

// NewTokenBucket creates a bucket of given capacity & refill rate.
// (Factory Pattern)
func NewTokenBucket(capacity int, refillRatePerSec float64) *TokenBucket {
	return &TokenBucket{
		capacity:   float64(capacity),
		refillRate: refillRatePerSec,
		tokens:     float64(capacity),
		last:       time.Now(),
	}
}

// Allow tries to consume one token.
// Returns true if successful (i.e. rate‑limit not exceeded).
// O(1) time, thread‑safe.
func (b *TokenBucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	// refill tokens
	elapsed := now.Sub(b.last).Seconds() // Calculate how much time has passed since the last request.
	b.tokens += elapsed * b.refillRate
	if b.tokens > b.capacity {
		b.tokens = b.capacity
	}
	b.last = now

	if b.tokens >= 1 {
		b.tokens--
		return true
	}
	return false
}
