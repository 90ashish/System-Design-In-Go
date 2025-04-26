# 1. Requirements & Constraints

## 1.1 Functional vs. Non-Functional

- **Functional**: What the system must do (e.g., “Allow users to upload photos”).
- **Non-Functional**: How well it does it (e.g., “Handle 10,000 uploads per second,” “99.9% uptime,” “data encrypted at rest”).

## 1.2 Use-Cases, Actors, User Stories

- **Actor**: Who interacts (user, admin, other systems).
- **Use-Case**: A scenario (“User logs in,” “System sends email notification”).
- **User Story**: "As a Customer, I want to reset my password so that I can regain account access."

## 1.3 Non-Functional Requirements (NFRs)

- **Scalability**
  - Vertical: Bigger machine (more CPU/RAM).
  - Horizontal: More machines in parallel.

- **Availability & Reliability**
  - Availability: Percentage of time the system is up (e.g., 99.9%).
  - MTBF/MTTR: Mean time between failures vs. mean time to recover.

- **Performance**
  - Latency: How long one operation takes (p50 = median, p99 = 99th percentile).
  - Throughput: How many operations per second (RPS/QPS).

- **Consistency & Durability**
  - Consistency: Will all users see the same data at the same time?
  - Durability: Once data is stored, it won’t be lost.

- **Security & Compliance**
  - Encryption, authentication, and legal rules (GDPR, PCI-DSS).

## 1.4 SLIs / SLOs / SLAs

- **SLI (Indicator)**: Measurable metric (latency, error rate).
- **SLO (Objective)**: Target for the SLI (e.g., 99% requests < 200 ms).
- **SLA (Agreement)**: Promise to customers; violation means penalty.
- **Error Budget**: Allowed margin of failure (100% – SLO).

---

# 2. Traffic & Capacity Planning

## 2.1 Load Estimation

- Estimate peak and average requests per second (RPS).
- Factor in daily/weekly spikes and future growth.

## 2.2 Resource Sizing

- Estimate CPU cycles, memory per request, disk I/O, and network bandwidth needed for given RPS.

## 2.3 Cost Modeling

- Cloud vs. On-Prem: Pay-as-you-go vs. up-front hardware cost.
- CAPEX vs. OPEX: Capital vs. operational expenses.
- Balance performance requirements against budget.

---

# 3. Core Fundamentals

## 3.1 Scalability

- Vertical: Upgrade single server.
- Horizontal: Add more servers; needs coordination (load-balancer, data partitioning).

## 3.2 Latency vs. Throughput

- Latency: Time for one request.
- Throughput: Number of requests per time unit.
- Trade-off: Batching can increase throughput but raise latency.

## 3.3 Load Balancing

- Round-Robin: Distribute evenly in turn.
- Least Connections: Send to server with fewest active connections.
- Consistent Hashing: Route same keys to same servers (good for caching).

## 3.4 Caching

- In-process: Data stored in app memory.
- Distributed: Redis or Memcached cluster.
- CDN: Edge caching close to users.

**Invalidation**:
- Write-through: Write cache then database.
- Write-back: Write to cache, later flush to database.
- Write-around: Write directly to database; cache on read.

**Eviction**:
- LRU (Least Recently Used), LFU (Least Frequently Used), FIFO.

## 3.5 Concurrency & Parallelism

- Threads vs. Async: OS threads vs. lightweight tasks (goroutines).
- Mutex/Locks/Atomics: Prevent two threads from corrupting shared data.

## 3.6 CAP Theorem

- You can only have two of: Consistency, Availability, Partition tolerance.

## 3.7 Database Scaling

- Replication: Copy data across multiple servers for availability.
- Sharding: Split data by key range or hash among many servers.
- Partitioning: Similar to sharding, but can be within one database engine.

---

# 4. High-Level Architecture Patterns

- **Monolith**: All features in one deployable unit. Simple, but harder to scale small parts independently.
- **Microservices**: Each service owns its data & logic. Scales independently, but adds network complexity.
- **Serverless**: Functions triggered by events. No server management, but cold-start and vendor lock-in concerns.
- **Layered (n-tier)**: Presentation, business logic, data layers. Clear separation but can introduce latency.
- **Hexagonal/Clean**: Core logic independent of frameworks; port-adapter pattern eases testing.
- **Event-Driven & Reactive**: Services communicate via events; great for decoupling and scalability.
- **Domain-Driven Design (DDD)**: Model software around business domains and bounded contexts.

---

# 5. API & Interface Design

## 5.1 RESTful APIs

- Resources: Nouns in URLs (e.g., `/users/123`).
- Versioning: `/v1/users` vs. header-based.
- HATEOAS: Include links to related actions in responses.

## 5.2 gRPC / Protobuf

- Define services and messages in `.proto`.
- Supports streaming (server-stream, client-stream, bi-directional).

## 5.3 GraphQL

- Single endpoint; clients request exactly the fields they need.
- Watch out for N+1 query problems—use batching or caching.

## 5.4 API Gateways & BFF

- Gateway: Single entry point; can handle auth, rate limiting, routing.
- BFF (Backend for Frontend): Separate gateway tailored for each client type (mobile vs. web).

---

# 6. Data Storage & Management

## 6.1 Relational Databases

- Normalization: Eliminate data redundancy.
- Indexing: B-Trees, Hash Indexes, covering indexes speed up queries.
- Transactions & 2PC: Ensure all-or-nothing updates across multiple resources.

## 6.2 NoSQL Databases

- Key–Value: Very fast lookups (Redis, DynamoDB).
- Document: JSON-style records (MongoDB).
- Wide-Column: Huge datasets with flexible schemas (Cassandra).
- Graph: Relationships first (Neo4j).

## 6.3 Distributed Storage

- Replication: Sync or async copies for failover.
- Sharding: Partition data by key (range, hash, geo).
- Consistency: Choose strong (every read sees latest write) vs. eventual.

## 6.4 Caching Strategies

- (Covered above in Core Fundamentals.)

## 6.5 Search & Indexing

- Inverted Index: Map terms → document IDs (Elasticsearch).
- Autocomplete & Relevance: Specialized data structures, scoring algorithms.

---

# 7. Networking & Traffic Management

- DNS & Anycast: Route users to nearest datacenter automatically.

- L4 vs. L7 Load Balancers:
  - L4 (TCP-level), L7 (HTTP-level; can inspect headers & cookies).

- Reverse Proxies: Nginx or Traefik handle TLS termination, routing.

- CDNs: Cache static assets at edge nodes (images, JS, CSS).

- Service Discovery: Keep a registry of healthy service instances (Consul, Kubernetes DNS, Istio).

---

# 8. Messaging & Event Systems

## 8.1 Message Queues vs. Pub/Sub

- Queues: Point-to-point; one consumer per message (RabbitMQ, SQS).
- Pub/Sub: Many subscribers get copies of each message (Kafka, Google Pub/Sub).

## 8.2 Stream Processing

- Stateless: Each event processed independently.
- Stateful: Keep sliding windows, aggregates.
- Exactly-Once: Complex; requires idempotency or special protocols (Kafka transactions).

## 8.3 Event Sourcing & CQRS

- Event Sourcing: Persist every change as an immutable event.
- CQRS: Separate read and write models for optimization.

## 8.4 Reliability

- Idempotency: Safe to retry without side-effects.
- Dead-Letter Queues: Store messages that failed repeatedly.

---

# 9. Distributed Systems Fundamentals

- Consensus: Protocols for agreeing on a leader or shared state (Paxos, Raft).

- Distributed Transactions:
  - Two-Phase Commit: Coordinated commit across services.
  - Saga Pattern: Break transaction into compensating steps.

- Coordination Services: Zookeeper, etcd for config and leader election.

- Time & Ordering:
  - Lamport Clocks: Logical timestamps.
  - Vector Clocks: Track causality in distributed events.

---

# 10. Reliability & Resilience

- Fault Isolation:
  - Bulkheads: Limit failure blast radius by isolating components.
  - Circuit Breakers: Stop calling failing services temporarily.

- Retries & Back-off: Exponential back-off to avoid thundering herd.

- Graceful Degradation: Offer reduced functionality under load.

- Graceful Shutdown: Let in-flight requests finish before stopping.

- Chaos Engineering: Intentionally inject failures to test resilience.

- Disaster Recovery:
  - RPO: How much data you can afford to lose.
  - RTO: How quickly you must recover.
  - Multi-region failover strategies.

---

# 11. Observability & Monitoring

- Logging: Structured JSON logs, centralized via ELK or Splunk.
- Metrics: Time-series databases (Prometheus), create dashboards and alerts.
- Distributed Tracing: Propagate trace IDs across services (OpenTelemetry, Jaeger).
- Health Checks & Heartbeats: Automatically remove unhealthy instances from load-balancer.

---

# 12. Security & Compliance

- Authentication & Authorization:
  - OAuth2, OpenID Connect, JWT tokens.
  - mTLS for service-to-service.

- Encryption:
  - TLS for data in-transit.
  - KMS/HSM for data at-rest.

- Rate Limiting & DDoS Protection: Token/leaky bucket, WAF rules.

- Auditing & Compliance: Keep logs, implement controls for PCI, SOC2, GDPR.

---

# 13. Scalability & Performance Tuning

- Sharding Patterns: Range, hash, directory (central routing table).

- Hot-Key Handling: Detect and split heavy partitions.

- Backpressure: Signal senders to slow down when overloaded.

- CDN & Edge: Push static and cacheable content closer to users.

- Database Optimizations: Denormalize for reads, materialized views, read-replicas.

---

# 14. Deployment & Operations

- Infrastructure as Code: Define servers, networks in Terraform or CloudFormation.

- Containers & Kubernetes:
  - Namespaces: Isolate teams/apps.
  - DaemonSets: Run one pod per node.
  - StatefulSets: Stable network IDs and storage for stateful apps.

- CI/CD Pipelines: Automate build, test, deploy. Use blue/green or canary releases.

- Autoscaling: Horizontal Pod Autoscaler (HPA) based on CPU/memory or custom metrics.

---

# 15. Advanced & Emerging Topics

- Service Mesh: Side-car proxies for mTLS, traffic shifting, observability (Istio, Linkerd).
- Serverless: FaaS (Lambda), pros (no servers) vs. cons (cold start, vendor lock-in).
- ML in Systems: Feature stores, online vs. offline inference, model rollback.
- Blockchain & DLT: Basics of decentralization and consensus.
- Edge & IoT: Pushing compute and storage toward devices for low latency.

---

# 16. Case Studies & Mock Drills

## Design Exercises

- URL shortener: key generation, redirection, analytics.
- Chat service: real-time messaging, presence, history.
- Notification system: fan-out, deduplication, retries.

## Scalability Drills

- Build a distributed rate limiter.
- Generate unique IDs (Snowflake algorithm).
- Leaderboard with frequent score updates.

## Failure Scenarios

- Single region outage: automatic failover.
- Network partition: degraded mode, eventual consistency.
