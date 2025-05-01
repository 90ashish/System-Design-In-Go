# Designing a LRU Cache

## Requirements

The **LRU (Least Recently Used) Cache** should support the following operations:

- **`put(key, value)`**:  
  Insert a key-value pair into the cache.  
  - If the cache is at capacity, remove the least recently used item before inserting the new item.

- **`get(key)`**:  
  Get the value associated with the given key.  
  - If the key exists in the cache, move it to the front of the cache (most recently used) and return its value.  
  - If the key does not exist, return `-1`.

---

## Additional Requirements

- The cache should have a **fixed capacity**, specified during initialization.
- The cache should be **thread-safe**, allowing concurrent access from multiple threads.
- The cache should be **efficient** in terms of time complexity for both `put` and `get` operations, ideally **O(1)**.

---

## Working Flow
1. **Initialization**: `New(capacity)` creates an LRU cache combining a hashmap and a doubly-linked list.  
2. **Put(key,value)**:  
   - If key exists, update its value and move its node to the front (most-recent).  
   - If new and at capacity, remove the tail node (least-recent) before inserting at the front.  
3. **Get(key)**:  
   - If found, move its node to the front and return its value.  
   - If missing, return “not found.”  
4. **Concurrency**: All operations lock a `sync.RWMutex` to ensure thread safety.

## Design Patterns
- **Factory**: `New(capacity)` hides setup details.  
- **Cache (LRU Eviction)**: Uses a map + linked list for O(1) access and eviction logic.  
- **Mutex Guard**: Ensures safe concurrent reads/writes.  
