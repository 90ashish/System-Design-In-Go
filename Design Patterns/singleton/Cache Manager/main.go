package main

import (
	"container/list"
	"fmt"
	"sync"
)

// CacheEntry holds the key-value pair for the cache
type CacheEntry struct {
	key   string
	value string
}

// CacheManager represents the singleton cache with LRU eviction
type CacheManager struct {
	cache     map[string]*list.Element
	evictList *list.List
	capacity  int
	mutex     sync.Mutex
}

// Singleton instance and sync object
var (
	instance *CacheManager
	once     sync.Once
)

// GetInstance - Provides the singleton instance of the CacheManager
func GetInstance(capacity int) *CacheManager {
	once.Do(func() {
		instance = &CacheManager{
			cache:     make(map[string]*list.Element),
			evictList: list.New(),
			capacity:  capacity,
		}
	})
	return instance
}

// Get - Retrieves a value from the cache
func (cm *CacheManager) Get(key string) (string, bool) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if elem, exists := cm.cache[key]; exists {
		cm.evictList.MoveToFront(elem) // Mark item as recently used
		return elem.Value.(*CacheEntry).value, true
	}
	return "", false
}

// Put - Adds a key-value pair to the cache
func (cm *CacheManager) Put(key string, value string) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	// If key exists, update value and move it to the front
	if elem, exists := cm.cache[key]; exists {
		elem.Value.(*CacheEntry).value = value
		cm.evictList.MoveToFront(elem)
		return
	}

	// Evict LRU if capacity exceeded
	if len(cm.cache) >= cm.capacity {
		cm.evict()
	}

	// Insert the new entry
	entry := &CacheEntry{key, value}
	element := cm.evictList.PushFront(entry)
	cm.cache[key] = element
}

// evict - Removes the least recently used item from the cache
func (cm *CacheManager) evict() {
	if elem := cm.evictList.Back(); elem != nil {
		delete(cm.cache, elem.Value.(*CacheEntry).key)
		cm.evictList.Remove(elem)
	}
}

// DisplayCache - Displays cache content for visualization
func (cm *CacheManager) DisplayCache() {
	fmt.Println("Current Cache:")
	for elem := cm.evictList.Front(); elem != nil; elem = elem.Next() {
		entry := elem.Value.(*CacheEntry)
		fmt.Printf("[%s: %s] ", entry.key, entry.value)
	}
	fmt.Println()
}

func main() {
	cache := GetInstance(3)

	cache.Put("A", "Alpha")
	cache.Put("B", "Bravo")
	cache.Put("C", "Charlie")

	cache.DisplayCache()

	cache.Get("A")          // Access 'A' (moves it to front)
	cache.Put("D", "Delta") // 'B' should be evicted (LRU Policy)

	cache.DisplayCache()

	cache.Get("C")         // Access 'C' (moves it to front)
	cache.Put("E", "Echo") // 'A' should be evicted (LRU Policy)

	cache.DisplayCache()
}
