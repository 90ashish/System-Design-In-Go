# Fundamentals of System Design

## 1. Scalability

### Vertical Scaling (Scale Up)
- **What it is**: Increasing the power of a single machine by adding more CPU, RAM, or disk.
- **Example**: Upgrading a database server from 8 cores & 32 GB RAM to 16 cores & 64 GB RAM so it can handle more queries.

**Pros**:
- Easy: often just a hardware or cloud-instance resize.
- No application changes needed.

**Cons**:
- There’s an upper limit (you can’t add infinite RAM).
- Single point of failure remains.

### Horizontal Scaling (Scale Out)
- **What it is**: Adding more machines/instances to your pool of servers.
- **Example**: Running your web application on three identical servers behind a load balancer instead of one.

**Pros**:
- Virtually unlimited growth by adding more nodes.
- Better fault tolerance: if one server fails, others continue serving.

**Cons**:
- Requires coordination (e.g., sharing session data or caching).
- More complex networking and deployment.

---

## 2. Latency vs. Throughput

### Latency
- **Definition**: Time taken for a single request to travel from client to server, be processed, and return.
- **Example**: A search query on a website returns results in 50 ms.

### Throughput
- **Definition**: Number of requests the system can handle per second.
- **Example**: A service handling 2,000 requests per second (RPS).

### Trade-offs & Percentiles
- **p50 (Median)**: Half of requests are faster than this time.
- **p99 (Tail)**: 99% of requests are faster than this time; captures “worst” 1%.

**Why it matters**:  
You may optimize average latency (p50) but ignore the slowest 1% (p99), which impacts user experience.

**Illustration**:
- If p50 = 20 ms but p99 = 500 ms, most users see snappy responses, but 1 in 100 users waits half a second.

---

## 3. Load Balancing

Distributes incoming traffic among multiple servers to improve utilization and reliability.

| Algorithm | How It Works | Use Case |
|:---|:---|:---|
| Round-Robin | Sends each new request to the next server in order | Evenly distributes simple, stateless traffic |
| Least Connections | Chooses the server with the fewest active connections | Best when requests vary in duration |
| Consistent Hashing | Hashes client ID or request key to pick a server | Useful for sticky sessions or distributed caches |

### Example:
A pool of three servers A, B, C receiving requests R1, R2, …

- Round-Robin: R1→A, R2→B, R3→C, R4→A,…
- Least Connections: directs new requests to whichever server currently has the lightest load.

---

## 4. Caching

Storing frequently accessed data in fast storage to reduce load on primary data sources.

### Technologies
- **Redis**: In-memory data store supporting advanced data types (lists, sets, sorted sets).
- **Memcached**: Simple, high-performance key–value store.

### Invalidation Strategies
- **Write-Through**:
  - **Flow**: Write to cache and data store simultaneously.
  - **Pro**: Cache always up-to-date.
  - **Con**: Latency on writes doubles.

- **Write-Back (Write-Behind)**:
  - **Flow**: Write to cache first; asynchronously persist to data store.
  - **Pro**: Fast writes.
  - **Con**: Risk of data loss if cache crashes before persisting.

- **Write-Around**:
  - **Flow**: Write only to data store, not cache; cache is populated on next read.
  - **Pro**: Prevents cache pollution from rarely read data.
  - **Con**: First read after write experiences full latency.

### Eviction Policies
- **LRU (Least Recently Used)**: Evicts items that haven’t been accessed for the longest time.
- **LFU (Least Frequently Used)**: Evicts items accessed the fewest times.
- **FIFO (First In, First Out)**: Evicts the oldest items in cache.

### Example:
If a cache max size is 100 items, when the 101st item arrives:
- **LRU** removes the one not used since longest ago.
- **LFU** removes the item with the smallest access count.

---

## 5. Concurrency & Parallelism

### Multi-threading vs. Async
- **Multi-threading**: Multiple threads share memory and run in parallel on different CPU cores.
- **Async (Event-Driven)**: Single thread handles many tasks by switching contexts on I/O waits (e.g., Node.js callbacks) or via lightweight goroutines in Go.

### Synchronization Primitives
- **Mutex (Mutual Exclusion)**:
  - Locks a section of code so only one thread/goroutine can execute it at a time.
  - Example in Go:
    ```go
    var mu sync.Mutex
    mu.Lock()
    counter++
    mu.Unlock()
    ```

- **Locks / Read-Write Locks**: Allow multiple readers or one writer.
- **Atomic Operations**: Hardware-level instructions that update values without full locks (e.g., `atomic.AddInt64(&counter, 1)`).

---

## 6. CAP Theorem

In the presence of a network Partition, you can choose only **Consistency** or **Availability**:

- **Consistency (C)**: All clients see the same data at the same time.
- **Availability (A)**: Every request gets a response (even if it’s stale).
- **Partition Tolerance (P)**: System continues to operate despite network failures between nodes.

### Example Choices:
- **CP System**: Prioritizes consistency (e.g., a banking ledger); may reject requests if partitioned.
- **AP System**: Prioritizes availability (e.g., DNS); may serve stale data during partitions.

---

## 7. Database Scaling

### Replication
- **What it is**: Copying data from one (master) to multiple (slave) nodes.
- **Example**: One master handles writes; three replicas handle reads.

### Sharding (Horizontal Partitioning)
- **What it is**: Splitting data across separate databases based on a key.
- **Example**: Orders with IDs 1–1,000,000 go to shard 1; 1,000,001–2,000,000 go to shard 2.

### Partitioning (Vertical vs. Horizontal)
- **Vertical Partitioning**: Splitting a table into multiple tables by columns.
  - **Example**: A User table split into UserProfile and UserCredentials.

- **Horizontal Partitioning**: Same as sharding—splitting rows across tables.

---

## Putting It All Together

When designing a system, you’ll mix and match these techniques:

- Scale your web tier horizontally behind a load balancer (round-robin).
- Cache frequent reads in Redis with LRU eviction and write-through for critical data.
- Optimize your code with goroutines for async processing, using mutexes only where needed.
- Ensure your database replicates for reads, and shard when a single instance can’t handle the write load.
- Balance latency (aim for low p50 and p99) against throughput (high RPS) by profiling under load.
- Decide on a CAP trade-off—choose CP for financial data, AP for global content delivery.
