# 8. Logging, Monitoring & Observability

Ensuring you can see and understand what’s happening inside your system is critical for reliability and rapid issue resolution. Below are the core pillars of observability, each explained with simple terminology and examples.

---

## 8.1 Distributed Tracing

Tracks a request as it flows through multiple services, showing timing and relationships.

### OpenTelemetry

- **What it is**: A vendor-neutral set of SDKs, APIs, and agents for instrumenting code.
- **How it works**: You add OpenTelemetry libraries to each service; they automatically collect spans (units of work) and export them.

**Example (Go)**:
```go
// Go example: create a tracer and instrument an HTTP handler
tracer := otel.Tracer("order-service")
http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
  ctx, span := tracer.Start(r.Context(), "CreateOrder")
  defer span.End()
  // ... business logic ...
})
```

### Jaeger

- **What it is**: A popular open-source tracing backend.
- **Use Case**: Visualize trace timelines, find slow spans, see service-to-service dependencies.

### Zipkin

- **What it is**: Another open-source tracing system with similar features to Jaeger.
- **Key Difference**: Slightly simpler architecture; some prefer its UI or language support.

### Flow

- Service A receives a request → creates span A1.
- Service A calls B → creates child span B1.
- Service B calls C → child span C1.
- All spans are reported to Jaeger/Zipkin and stitched into a single trace.

---

## 8.2 Log Aggregation

Collects logs from all services into a central store, making search and analysis easy.

### ELK Stack (Elasticsearch, Logstash, Kibana)

- **Logstash**: Agents or pipelines ingest logs from files, syslog, or HTTP and transform them (e.g., parse JSON).
- **Elasticsearch**: Stores log entries in a distributed, searchable index.
- **Kibana**: Web UI for querying, filtering, and visualizing logs (dashboards, charts).

**Example Workflow**:
- Each service writes structured logs to a file or stdout.
- Filebeat (lightweight shipper) tails these logs and forwards to Logstash.
- Logstash parses fields (timestamp, level, service, message) and sends to Elasticsearch.
- You use Kibana to search `"ERROR in payment-service"` and see all matching entries.

---

## 8.3 Metrics & Dashboards

Captures numerical data over time to alert on abnormal behavior and track performance.

### Prometheus

- **How it works**:
  - Services expose an HTTP `/metrics` endpoint with counters and gauges.
  - Prometheus server “scrapes” these endpoints at regular intervals (e.g., every 15s).
  - Stores time-series data locally.

**Example Metric (Go)**:
```go
httpRequests := prometheus.NewCounterVec(
  prometheus.CounterOpts{Name: "http_requests_total", Help: "Count of HTTP requests"},
  []string{"path", "method"},
)
prometheus.MustRegister(httpRequests)

// In handler:
httpRequests.WithLabelValues("/create", "POST").Inc()
```

### Grafana

- **What it is**: Dashboard and graph editor that connects to Prometheus (and other data sources).
- **Use Case**: Build dashboards showing request rate, error rate, latency percentiles over time.

---

## 8.4 Error Tracking

Captures exceptions or errors in your applications, with context and stack traces, and surfaces them to developers.

### Sentry

- **Features**:
  - SDKs for many languages capture uncaught exceptions automatically.
  - Groups similar errors, shows affected users, and provides stack traces.
  - Alerts via email, Slack, or webhooks when new or regressed errors occur.

### Datadog

- **Features**:
  - Full observability platform: logs, metrics, traces, and error tracking in one place.
  - APM (Application Performance Monitoring) captures errors alongside traces.
  - Customizable alerts and dashboards integrated with Datadog monitors.

**Example (Python + Sentry)**:
```python
import sentry_sdk
sentry_sdk.init("https://<PUBLIC_KEY>@sentry.io/<PROJECT_ID>")

def process_order(order):
    try:
        # ... processing logic ...
    except Exception as e:
        sentry_sdk.capture_exception(e)
        raise
```

---

## Bringing It All Together

- Instrument your code for traces (OpenTelemetry), metrics (Prometheus), and structured logs.
- Export and centralize data: send traces to Jaeger/Zipkin, metrics to Prometheus, logs to ELK, errors to Sentry/Datadog.
- Build dashboards and alerts in Grafana and Kibana so you know immediately when something goes wrong.
- Use traces to drill down from a high-level alert (e.g., increased latency in Grafana) into service-to-service call details in Jaeger.

By combining distributed tracing, log aggregation, real-time metrics, and error tracking, you gain full visibility into your system’s behavior, making debugging faster and ensuring reliability at scale.
