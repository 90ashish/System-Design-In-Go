2. Databases & Storage
2.1 SQL vs. NoSQL – When to Choose Each
SQL (Relational Databases)

    Structure: Tables with fixed schemas (columns and types).

    Transactions & ACID: Guarantees Atomicity, Consistency, Isolation, Durability.

    Use Cases:

        Strong data relationships (e.g. joins between Users ↔ Orders).

        When you need strict correctness (e.g. banking, inventory).

    Example:

    CREATE TABLE users (
      id SERIAL PRIMARY KEY,
      name VARCHAR(100),
      email VARCHAR(100) UNIQUE
    );

NoSQL (Non-Relational Databases)

    Structure: Flexible schemas—documents, key–value pairs, wide-columns, or graphs.

    Consistency Models: Often Eventual Consistency (BASE: Basically Available, Soft state, Eventual consistency).

    Use Cases:

        Rapid iteration on changing data shapes (e.g. user profiles with varying fields).

        Massive horizontal scale with simple access patterns (e.g. caching, session stores).

    Example (Document Store):

    // In MongoDB
    db.users.insertOne({
      _id: ObjectId(),
      name: "Alice",
      preferences: { theme: "dark", notifications: true }
    });

When to Choose
Scenario	Pick SQL	Pick NoSQL
Complex joins & constraints	✔️	❌
Evolving schema	❌	✔️
Ultra-high write throughput	❌ (harder to shard)	✔️ (easy to partition/shard)
Strict transactions required	✔️	❌ (or limited)
2.2 Indexing

Proper indexes speed up lookups; without them, a query may scan every row.
B-Tree Indexes

    How it works: Balanced tree of key pointers; good for range scans (WHERE age BETWEEN 20 AND 30).

    Example:

    CREATE INDEX idx_users_email ON users(email);

Hash Indexes

    How it works: Hash table under the hood; excellent for exact matches (WHERE id = 123).

    Limitation: Not suitable for range queries.

Covering Indexes

    How it works: Index that includes all columns a query needs, so the database never touches the table itself.

    Example:

    -- This index “covers” the query on users(name)
    CREATE INDEX idx_users_name_email 
      ON users(name, email);

    -- Fast lookup: SELECT email FROM users WHERE name = 'Bob';

2.3 Sharding & Partitioning

Splitting your data so no single node holds the entire dataset.
Key-Based (Hash) Sharding

    Strategy: Compute hash(key) % N to pick one of N shards.

    Example:

    shard_id = hash(user_id) % 4

Range-Based Sharding

    Strategy: Assign contiguous key ranges to each shard.

    Example:

        Shard 1: user_id 1–1,000,000

        Shard 2: user_id 1,000,001–2,000,000

Geo-Based Partitioning

    Strategy: Separate data by geography.

    Example:

        EU shard for European users

        US shard for North American users

2.4 Replication Strategies

Keeping multiple copies of data for reliability and scale.
Leader–Follower (Master–Slave)

    Flow: All writes go to Leader; Followers replicate asynchronously and serve reads.

    Pros: Simple.

    Cons: Leader is a bottleneck for writes.

Multi-Leader (Active–Active)

    Flow: Multiple nodes accept writes and replicate to each other.

    Pros: High write availability across datacenters.

    Cons: Conflict resolution required.

Quorum-Based

    Flow: A write must be acknowledged by a majority (“write quorum”), and reads may require a majority (“read quorum”).

    Pros: Tunable consistency vs. latency.

    Cons: More network chatter.

2.5 Consistency Models
Strong Consistency

    Definition: After a write completes, all reads see that write.

    Use Case: Bank account balances.

Eventual Consistency

    Definition: Writes propagate in background; reads might see stale data briefly.

    Use Case: Social-media likes/counters.

2.6 Workloads
OLTP (Online Transaction Processing)

    Characteristics: Many short, transactional queries (inserts/updates).

    Example: E-commerce checkout, banking transactions.

OLAP (Online Analytical Processing)

    Characteristics: Fewer, complex, read-heavy queries over large datasets.

    Example: Data-warehouse analytics, BI dashboards.

2.7 Specialized Stores
Time-Series Databases (e.g. InfluxDB)

    Optimized for: High-volume timestamped data, time-range queries, downsampling.

    Example Use: Sensor readings, application metrics.

Columnar (Wide-Column) Stores (e.g. Apache Cassandra)

    Optimized for: Massive scale, high write throughput, row-distributed by key.

    Data Model: Tables with dynamic columns per row.

    Example Use: User activity logs, IoT event streams.

2.8 Data Storage & Management
Relational DBs

    Normalization: Organize data to reduce duplication (1NF, 2NF, 3NF).

    Transactions & 2PC: Two-Phase Commit for cross-node consistency in distributed transactions.

    Indexing Recap: B-Trees, hashes, covering indexes.

NoSQL Categories

    Key–Value: Simple get/set (e.g. Redis).

    Document: JSON/BSON objects (e.g. MongoDB).

    Wide-Column: Rows with dynamic columns (e.g. Cassandra).

    Graph: Nodes and edges (e.g. Neo4j).

Distributed Storage

    Replication & Sharding: Mix leader–follower + hash/range sharding for both high availability and capacity.

    Consistency Choices: Pick strong or eventual based on your SLAs.

Search & Indexing

    Inverted Indexes (Elasticsearch): Map terms → document IDs for full-text search.

    Autocomplete & Ranking: Prefix trees + scoring algorithms to return best matches as you type.