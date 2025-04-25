7. Distributed Systems & Microservices

Below is a detailed, beginner-friendly overview of key concepts in distributed systems and microservices, complete with simple examples.
7.1 Architectural Styles
Style	Description	Pros	Cons	Example Use Case
Monolith	Single, unified codebase and deployment artifact	Simple development, single transaction span	Hard to scale pieces independently; large deployments	Small teams, simple apps
Microservices	Many small services, each with its own code/deployment	Independently deployable; scale per service	Operational overhead; cross-service communication	Large e-commerce with separate order, catalog, user services
Serverless	Functions deployed to managed FaaS platform (e.g. AWS Lambda)	No server management; scales automatically	Cold starts; limited execution time; vendor lock-in	Lightweight event-driven tasks (image thumbnailing)
7.2 Service Discovery

When services multiply, hard-coding addresses breaks. Service discovery automates finding service endpoints.

    Client-Side Discovery

        Service registers itself with registry (e.g., Consul, Eureka).

        Client queries registry for available instances.

        Client load-balances among returned addresses.

    [ UserService ] → Register → [ Consul ]
                       ↖ Query ↗
                   [ OrderService ]

    Server-Side Discovery

        Client sends request to a load-balancer (API gateway or service mesh sidecar).

        Load-balancer queries registry and forwards the request.

    Tools

        Consul: Key–value store + health checks + DNS interface.

        Eureka: Netflix open-source registry; Java-centric but language-agnostic clients.

        Kubernetes DNS: Every Service gets a DNS name (ordersvc.default.svc.cluster.local).

        Service Mesh (Istio, Linkerd): Sidecar proxies automatically register, discover, and load-balance, plus provide telemetry and mTLS.

7.3 API Gateway & BFF

    API Gateway

        Role: Single entry point for external clients → routes to internal services, handles auth, rate-limiting, metrics, SSL termination.

        Example:

            Client calls https://api.shop.com/orders.

            Gateway verifies JWT, applies rate limit, logs request, then forwards to OrderService.

    BFF (Backend for Frontend)

        Role: A thin backend layer tailored per client type (web, mobile), aggregating multiple services into one optimized response.

        Example:

            Web BFF: Combines UserService + CatalogService calls into one JSON.

            Mobile BFF: Only returns minimal fields needed to reduce payload size.

7.4 Inter-service Communication
REST

    Format: HTTP + JSON

    Pros: Ubiquitous; human-readable; easy to debug with cURL.

    Cons: Verbose; higher overhead; no built-in streaming.

GET /users/123 HTTP/1.1
Host: api.shop.com
Accept: application/json

gRPC

    Format: HTTP/2 + Protocol Buffers (binary)

    Pros: High performance; bi-directional streaming; strict contracts via .proto files.

    Cons: Requires code generation; less human-readable without tooling.

// user.proto
service UserService {
  rpc GetUser (GetUserRequest) returns (GetUserResponse);
}
message GetUserRequest { int64 id = 1; }
message GetUserResponse { string name = 1; string email = 2; }

    Streaming Example:

        Server Streaming: Client sends one request, server streams many responses.

        Client Streaming: Client streams many requests, server replies once.

        Bi-Directional: Both stream simultaneously.

7.5 Transactions
Two-Phase Commit (2PC)

    Prepare phase

        Coordinator asks all participants “Can you commit?”

        Each replies “Yes” or “No”.

    Commit phase

        If all said “Yes,” coordinator sends “Commit.”

        Otherwise sends “Rollback.”

Drawback: Blocks resources if coordinator or participants fail.
Saga Pattern

    Concept: Split a distributed transaction into a sequence of local transactions, each with a corresponding compensating action if something fails.

    Example (Order Workflow):

        Create Order (local DB)

        Reserve Inventory (InventoryService)

        Charge Payment (PaymentService)

    On failure at step 3:

        Run compensating transactions:

            Refund if charged

            Release inventory

            Cancel order

CreateOrder → ReserveInventory → ChargePayment
     ↑               ↑             ↑
CancelOrder ← ReleaseInventory ← RefundPayment

7.6 Data Consistency
Pattern	How It Works	Pros	Cons
2PC	Blocking commit protocol across services	Strong consistency	Slow; blocking; single coordinator risk
Eventual Consistency	Services publish events; other services react asynchronously	Highly available; scales easily	Temporary stale reads; complex recovery

    2PC: Use when you must have atomic cross-service commits (rare in microservices).

    Eventual Consistency: Publish domain events (e.g., “OrderCreated”); subscribers update their own data stores when they receive the event. Systems converge over time.

Bringing It Together

    Choose your style: start monolithic, break into microservices as complexity grows, or use serverless for event-driven pieces.

    Automate discovery with a registry or service mesh so services find each other reliably.

    Front your services with an API gateway and tailor BFFs for different clients.

    Pick the right RPC: REST for simplicity, gRPC for performance and streaming.

    Handle multi-service workflows with Saga for partial failures; reserve 2PC for rare truly atomic needs.

    Balance consistency needs—often eventual consistency suffices and lets you scale.

By applying these patterns thoughtfully, you’ll build distributed systems that are modular, scalable, and maintainable.