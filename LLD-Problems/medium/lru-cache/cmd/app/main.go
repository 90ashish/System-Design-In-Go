package main

import (
	"fmt"
	"sync"
	"time"

	"lru-cache/internal/lrucache"
)

func main() {
	// Create an LRU cache of capacity 2 (Factory Pattern)
	cache := lrucache.New(2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		cache.Put("a", 1) // Cache: a
		fmt.Println("Put a=1")
		cache.Put("b", 2) // Cache: b, a
		fmt.Println("Put b=2")

		// Get "a" → exists, move to front (MRU)
		if v, ok := cache.Get("a"); ok {
			fmt.Println("Get a:", v.(int))
		} else {
			fmt.Println("Get a: -1")
		}

		cache.Put("c", 3) // Evicts LRU "b"; Cache: c, a
		fmt.Println("Put c=3")

		// Get "b" → evicted
		if v, ok := cache.Get("b"); ok {
			fmt.Println("Get b:", v.(int))
		} else {
			fmt.Println("Get b: -1")
		}
	}()

	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)

		// Get "a" → still in cache
		if v, ok := cache.Get("a"); ok {
			fmt.Println("Get a:", v.(int))
		} else {
			fmt.Println("Get a: -1")
		}

		cache.Put("d", 4) // Evicts LRU "c"; Cache: d, a
		fmt.Println("Put d=4")

		// Get "c" → evicted
		if v, ok := cache.Get("c"); ok {
			fmt.Println("Get c:", v.(int))
		} else {
			fmt.Println("Get c: -1")
		}

		// Get "d" → exists
		if v, ok := cache.Get("d"); ok {
			fmt.Println("Get d:", v.(int))
		} else {
			fmt.Println("Get d: -1")
		}
	}()

	wg.Wait()
}
