# LLD/System Design Interview Question – Rate Limiter

**Question:**

> Design a distributed API rate‑limiter service that enforces per‑user quotas (e.g. 100 requests/minute) with burst support.

Your design should cover:

- **Class/interface definitions** for the in‑process limiter (e.g. `TokenBucket` or `LeakyBucket`).
- **Data structures and algorithms** (sliding window vs. token bucket).
- **Synchronization/thread‑safety** in a single JVM/process.
- **Distributed coordination** so that multiple service instances share counters (e.g. Redis, in‑memory caches, or consistent hashing).
- **Handling** clock skew, retries, and fault tolerance.
- **Extensibility** to support multiple rate‑limits per API key or endpoint.


## How It Works

1. **Manager.Allow(key)**  
   - **Read-lock** the map of buckets and try to fetch `TokenBucket` for `key`.  
   - If missing, **upgrade to write‑lock**, create a new `TokenBucket` via `NewTokenBucket(capacity, refillRate)`, store it, then unlock.  
   - Delegate to that bucket’s `Allow()` method.

2. **TokenBucket.Allow()**  
   - **Lock** the bucket.  
   - Compute time since last refill (`elapsed := now – last`).  
   - **Refill** tokens: `tokens += elapsed * refillRate`, capped at `capacity`.  
   - Update `last = now`.  
   - If `tokens ≥ 1`, **consume** one token (`tokens–`) and return `true` (allowed).  
   - Otherwise return `false` (rate‑limited).  

All operations run in **O(1)** time and use **mutexes** (`sync.RWMutex` in Manager, `sync.Mutex` in TokenBucket) to ensure thread safety under concurrent access.
