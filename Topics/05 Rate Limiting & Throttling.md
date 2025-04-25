5. Rate Limiting & Throttling

Rate limiting and throttling protect services from overload or abuse by controlling how many requests a client can make in a given time period.
5.1 Algorithms
Token Bucket

    Concept: Tokens accumulate in a “bucket” at a steady rate (e.g., 5 tokens/sec). Each request consumes one token.

    Behavior:

        Burstable: If the bucket has saved tokens, a client can send requests in a quick burst up to the bucket’s capacity.

        Refill: Tokens are added at the configured rate until capacity is reached.

    Example:

    capacity = 10 tokens
    refill_rate = 5 tokens/sec

    At t=0: bucket=10
    Client sends 8 requests → bucket=2
    Wait 1s → bucket refills to min(10, 2+5)=7

    Use Case: APIs that allow occasional bursts but limit sustained traffic.

Leaky Bucket

    Concept: Requests enter a queue (the “bucket”) and depart at a constant rate. If the bucket overflows, excess requests are dropped or delayed.

    Behavior:

        Smoothing: Even if arrivals are bursty, departures remain steady.

        Overflow: Once capacity is exceeded, new requests are rejected until space frees up.

    Example:

    capacity = 5 requests
    drain_rate = 1 request/100ms

    If 10 requests arrive instantly:
      - 5 are queued
      - 5 are dropped
    Then every 100ms 1 request leaves the bucket.

    Use Case: Traffic shaping on network routers or ingress controllers.

Sliding Window

    Concept: Counts requests in a moving time window (e.g., last 60 seconds).

    Behavior:

        Fixed Window: Simplest: reset counter at each interval (may allow bursts at boundaries).

        Sliding Window Log: Records a timestamp for each request, evicts those older than window.

        Sliding Window Counter: Maintains counts in two adjacent intervals to approximate a sliding effect without per-request logs.

    Example (Sliding Window Log):

    window = 60s, limit = 100 requests
    On each request:
      purge timestamps older than now-60s
      if count < 100: allow and record timestamp
      else: reject

    Use Case: Fairly tight control where you want no burst spikes even at boundaries.

5.2 Distributed Rate Limiting

When you have multiple application instances, you need a shared store to coordinate usage.
Redis-Based Counters

    How it works:

        Use Redis INCR and EXPIRE to count requests per key (e.g., per user/IP).

        Atomic operations ensure accurate counts across instances.

    Example (pseudo-code):

    key = "rate:123"               // user or API key
    count = redis.INCR(key)        // increment
    if count == 1:
        redis.EXPIRE(key, 60)      // set 60s window
    if count > limit:
        reject request
    else:
        allow request

    Pros: Strong consistency, simple to implement.

    Cons: Redis is a single point of failure unless clustered.

API-Gateway Built-Ins

    What they offer: Many gateways (e.g., AWS API Gateway, Kong, NGINX) provide native rate-limit plugins.

    Features:

        Pre-built policies (per-IP, per-user, per-API-key).

        Support for global or per-route limits.

        Dashboards and metrics out of the box.

    Example (AWS API Gateway):

        Configure Stage throttle: 1,000 RPS steady, 2,000 RPS burst.

        Gateway automatically enforces limits and returns HTTP 429.