package ratelimiter

import "sync"

// RateLimiter defines the interface for any limiter.
// (Pattern: Strategy)
type RateLimiter interface {
	Allow(key string) bool
}

// Manager orchestrates per‑key token buckets.
// ‣ Pattern: Singleton/Manager (central registry)
// ‣ DIP: Depends on the RateLimiter interface
type Manager struct {
	mu         sync.RWMutex
	buckets    map[string]*TokenBucket
	capacity   int
	refillRate float64
}

// NewManager constructs a Manager with global capacity/refill settings.
// (Factory Pattern)
func NewManager(capacity int, refillRatePerSec float64) *Manager {
	return &Manager{
		buckets:    make(map[string]*TokenBucket),
		capacity:   capacity,
		refillRate: refillRatePerSec,
	}
}

// Allow returns whether the request for `key` is permitted.
// Lazy‑inits a TokenBucket per key.
// Thread‑safe.
func (m *Manager) Allow(key string) bool {
	// fast path: read‑lock
	m.mu.RLock()
	tb, exists := m.buckets[key]
	m.mu.RUnlock()

	if !exists {
		// upgrade to write‑lock
		m.mu.Lock()
		// double‑check
		if tb, exists = m.buckets[key]; !exists {
			tb = NewTokenBucket(m.capacity, m.refillRate)
			m.buckets[key] = tb
		}
		m.mu.Unlock()
	}
	return tb.Allow()
}
