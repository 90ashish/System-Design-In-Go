# 6. Load Balancing & Traffic Management

Distributing incoming traffic effectively ensures your services stay responsive, available, and resilient. Below are the key techniques and patterns, explained in simple terms with examples.

---

## 6.1 DNS-Level Load Balancing

### 1. DNS Load Balancing

- **How it works**: The DNS server returns multiple IP addresses (A/AAAA records) for the same domain. Clients pick one (often the first) to connect.

**Example**:
```text
example.com → 192.0.2.1, 192.0.2.2, 192.0.2.3
```
A client query might return `[192.0.2.2, 192.0.2.3, 192.0.2.1]`, so it connects to `192.0.2.2`.

**Pros & Cons**:
- ✅ Very simple to set up
- ❌ No awareness of server health or load
- ❌ Client or resolver caching can bypass intended distribution

---

### 2. Geo-DNS

- **How it works**: DNS resolution varies based on the client’s geographic location.

**Example**:
- A user in Europe resolves to `eu.example.com → 203.0.113.10`.
- A user in Asia resolves to `asia.example.com → 198.51.100.20`.

**Use Case**: Direct users to the nearest regional data center to minimize latency.

---

### 3. Anycast

- **How it works**: Advertise the same IP address from multiple points of presence (PoPs) in different locations.

**Behavior**: Internet routing (BGP) sends each user’s packets to the closest PoP advertising that IP.

**Example**:
- `203.0.113.5` is announced from PoP in New York, London, and Tokyo.
- A Tokyo user’s packets go to the Tokyo PoP; a London user’s packets go to the London PoP.

**Pros & Cons**:
- ✅ Low latency by network proximity
- ✅ Built-in failover (if a PoP goes down, BGP reroutes)
- ❌ Requires BGP support and multiple data centers

---

## 6.2 Reverse Proxy & API Gateway

A reverse proxy or API gateway sits between clients and your services, routing requests, enforcing policies, and often providing observability.

| Tool | Type | Key Features | Simple Use Case |
|:---|:---|:---|:---|
| Nginx | Reverse Proxy | High performance, load balancing, SSL termination | Serve web traffic to app servers |
| Traefik | Dynamic Proxy | Auto-discover services via Docker/Kubernetes | Microservices ingress controller |
| Kong | API Gateway | Plugin architecture, auth, rate-limiting | Secure public APIs |
| Ambassador | API Gateway | Kubernetes-native, Envoy-based | Kubernetes edge proxy |

**Example (Nginx config snippet)**:
```nginx
http {
  upstream backend {
    server app1.internal:8080;
    server app2.internal:8080;
  }
  server {
    listen 80;
    location / {
      proxy_pass http://backend;
    }
  }
}
```
Clients hit Nginx on port 80, and it forwards to `app1` or `app2`.

---

## 6.3 Global Traffic Routing

Combines DNS-level and proxy-level techniques for worldwide traffic control.

- **Anycast (again)**: Routes at the IP layer to the nearest PoP.
- **Geo-DNS (again)**: Routes by region at the DNS layer.

**Multi-Cloud Failover**:
- **Health Checks**: DNS provider pings endpoints; if one fails, its record is removed.

**Example**:
- Use AWS Route 53 to monitor a primary site in `us-east-1`; on failure, Route 53 shifts DNS to secondary site in `eu-west-1`.

---

## 6.4 Resiliency Patterns

### 1. Circuit Breaker

- **Purpose**: Prevent cascading failures by “opening” a circuit when downstream calls start failing.

**States**:
- **Closed**: Normal operation; failures increment an error counter.
- **Open**: Immediately fail calls without invoking downstream for a cooldown period.
- **Half-Open**: Allow a limited number of test calls; if they succeed, close the circuit.

**Example Pseudocode (Resilience4j style)**:
```java
CircuitBreaker cb = CircuitBreaker.ofDefaults("serviceA");
Supplier<String> decorated = CircuitBreaker
    .decorateSupplier(cb, () -> callServiceA());
Try.ofSupplier(decorated)
    .onFailure(throwable -> fallback());
```

---

### 2. Failover

- **Active–Passive**: One primary instance handles traffic; on failure, a secondary takes over.
  - **Example**: Database replica is promoted to primary if the master fails.

- **Active–Active**: Multiple regions or clusters serve traffic simultaneously; health-checked at the load-balancer or DNS layer.
  - **Example**: Two data centers each run active services; a global load balancer directs traffic away from any that fail health checks.

---

## Putting It All Together

- Start at DNS level with geo-DNS or anycast to guide users globally.
- Front your clusters with reverse proxies/API gateways (Nginx, Kong) to handle SSL, auth, and load balancing.
- Implement circuit breakers in your client libraries to avoid hammering failing services.
- Plan failover strategies (active–passive or active–active) to maintain uptime during outages.

By layering these techniques—global routing, smart proxies, and resiliency patterns—you ensure that your system can handle large volumes of traffic, remain performant worldwide, and recover gracefully from failures.
