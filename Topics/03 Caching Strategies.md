# 3. Caching Strategies

## 3.1 Cache Types

| Type | Where It Lives | Typical Use |
|:---|:---|:---|
| In-Process Cache | Inside your application’s memory (e.g., a map in Go, a dictionary in Python) | Very low-latency reads for data you’ve just fetched |
| Distributed Cache | Separate service (e.g., Redis, Memcached) | Share cached data across multiple application servers |
| CDN / Edge Cache | Globally distributed network of servers (e.g., Cloudflare, Akamai) | Cache static assets (images, JS, CSS) close to users |

### In-Process Cache

- **What it is**: A simple key–value store living inside your app.

**Pros**:
- Fastest access (no network hop).

**Cons**:
- Data not shared between instances; limited by app memory.

**Example (Go pseudocode)**:
```go
var cache = map[string]string{}  // in-process

func GetUserName(id string) string {
  if name, ok := cache[id]; ok {
    return name            // cache hit
  }
  name := fetchFromDB(id)  // cache miss
  cache[id] = name
  return name
}
```

---

### Distributed Cache

- **What it is**: A separate caching service your apps talk to over the network.

**Pros**:
- Shared by all app instances; can grow beyond one machine.

**Cons**:
- Slightly higher latency; network-dependent.

**Example (Python-like pseudocode)**:
```python
import redis
client = redis.Redis(host="cache.example.com")

def get_user(id):
    name = client.get(id)
    if name:                    # cache hit
        return name
    name = db.query_user(id)    # cache miss
    client.set(id, name)
    return name
```

---

### CDN / Edge Cache

- **What it is**: Servers around the world storing copies of static content.

**Pros**:
- Users fetch assets from a server “edge” close to them → minimal latency.

**Cons**:
- Primarily for static content; requires proper cache-control headers.

**Example**:
- You deploy `logo.png` and `style.css` to Cloudflare.
- When a user in Brazil visits your site, they download those files from a nearby edge node instead of your origin server in the US.

---

## 3.2 Invalidation Strategies

How and when your cache gets updated or purged to stay consistent with the source of truth.

| Strategy | Flow | Pros & Cons |
|:---|:---|:---|
| Write-Through | Write → Cache and DB at the same time | ✅ Always up-to-date <br> ❌ Write latency doubles |
| Write-Back | Write → Cache only; asynchronously flush to DB later | ✅ Fast writes <br> ❌ Risk of data loss on crash |
| Write-Around | Write → DB only; cache on the next read | ✅ Avoids caching rarely read data <br> ❌ First read after write is slow |

**Example Scenario**: Updating a user’s email

- **Write-Through**:
  - App writes new email to Redis cache.
  - App writes new email to the database.
  - Both always match.

- **Write-Back**:
  - App writes new email only to Redis.
  - Background job later persists it to the database.

- **Write-Around**:
  - App writes new email directly to the database.
  - Cache is unchanged until someone requests that user’s data, then it’s re-populated.

---

## 3.3 Eviction Policies

When a cache is full, these rules decide which items to remove.

- **LRU (Least Recently Used)**:
  - Evict the item that hasn’t been accessed for the longest time.
  - Good for workloads where recent data is more likely to be reused.

- **LFU (Least Frequently Used)**:
  - Evict the item with the fewest total accesses.
  - Good for workloads where some items are “hot” for long periods.

- **FIFO (First In, First Out)**:
  - Evict the oldest item in the cache, regardless of use.
  - Simple, but may remove “hot” items prematurely.

**Illustration**:  
Imagine a cache holding 3 items: A, B, C. You access A, then D arrives:
- **LRU**: Remove whichever of B or C you haven’t used recently.
- **LFU**: Remove whichever of B or C has been accessed fewer times.
- **FIFO**: Remove A (it was added first), even though you just accessed it.

---

## 3.4 CDN & Edge

### Key Players

- **Cloudflare**
  - Easy setup via DNS.
  - Global network, DDoS protection, automatic asset minification.

- **Akamai**
  - One of the oldest and largest CDNs.
  - Fine-grained control, very broad global presence.

### How It Works

- User requests `https://example.com/image.jpg`.
- DNS points to the CDN.
- CDN edge node checks its cache:
  - **Hit**: Returns the image immediately.
  - **Miss**: Fetches from origin, returns to user, and caches it for future requests.

---

## 3.5 Bloom Filters

A probabilistic data structure that tests set membership very quickly using minimal memory, with a small chance of false positives (but zero false negatives).

**Use Case**: Before querying the cache or database, check “might this key exist?”

- If Bloom filter says “no”, you skip both cache and DB reads.
- If “yes”, you proceed—but you may still find it isn’t in the cache/DB.

**Simple Example (Python-like pseudocode)**:
```python
# Initialize Bloom filter for user IDs
bloom = BloomFilter(size=10000, num_hashes=4)
bloom.add("user:123")

def get_user(id):
    if not bloom.contains(f"user:{id}"):
        return None     # definitely not in DB or cache
    return query_cache_or_db(id)
```

**Benefit**: Reduces wasted lookups for keys you know aren’t there.
