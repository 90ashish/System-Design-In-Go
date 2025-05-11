package main

import (
	"fmt"
	"sync"
	"time"

	"rate-limiter/internal/ratelimiter"
)

func main() {
	// 100 req/minute → ~1.667 tokens/sec
	mgr := ratelimiter.NewManager(100, 100.0/60.0)

	var wg sync.WaitGroup
	users := []string{"alice", "bob"}

	for _, u := range users {
		wg.Add(1)
		go func(user string) {
			defer wg.Done()
			for i := 1; i <= 20; i++ {
				if mgr.Allow(user) {
					fmt.Printf("[%s] request %d: ✅ allowed\n", user, i)
				} else {
					fmt.Printf("[%s] request %d: 🚫 rate‑limited\n", user, i)
				}
				time.Sleep(100 * time.Millisecond)
			}
		}(u)
	}

	wg.Wait()
}
