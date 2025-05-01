package lrucache

import (
	"container/list"
	"sync"
)

// entry holds a key/value pair for the doubly-linked list.
type entry struct {
	key   interface{}
	value interface{}
}

// LRUCache is a thread-safe LRU cache.
// Uses a hashmap + doubly-linked list to achieve O(1) Get/Put.
// Design Patterns:
//   - Factory (New)
//   - Cache (LRU eviction)
//   - Mutex for Thread Safety
type LRUCache struct {
	capacity int
	ll       *list.List                    // most-recent at Front
	cache    map[interface{}]*list.Element // key → *list.Element
	mu       sync.RWMutex                  // guards ll and cache
}

// New constructs an LRUCache with the given capacity.
// (Factory Pattern)
func New(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		ll:       list.New(),
		cache:    make(map[interface{}]*list.Element, capacity),
	}
}

// Get looks up a key’s value.
// If found, moves its element to front (MRU) and returns the value.
// If not found, returns nil, false.
// O(1) time.
// Thread-safe via RWMutex.
func (c *LRUCache) Get(key interface{}) (value interface{}, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, exists := c.cache[key]; exists {
		c.ll.MoveToFront(elem) // mark as most-recent
		return elem.Value.(*entry).value, true
	}
	return nil, false
}

// Put inserts or updates a key/value.
// If key exists, updates value and moves to front.
// If new key and at capacity, evicts Least-Recently-Used (tail).
// O(1) time.
// Thread-safe via RWMutex.
func (c *LRUCache) Put(key, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, exists := c.cache[key]; exists {
		// Update existing entry, move to front
		elem.Value.(*entry).value = value
		c.ll.MoveToFront(elem)
		return
	}

	// New entry
	if c.ll.Len() >= c.capacity {
		// Evict LRU at back
		lru := c.ll.Back()
		if lru != nil {
			ent := lru.Value.(*entry)
			delete(c.cache, ent.key)
			c.ll.Remove(lru)
		}
	}

	// Add to front as MRU
	e := &entry{key: key, value: value}
	elem := c.ll.PushFront(e)
	c.cache[key] = elem
}
