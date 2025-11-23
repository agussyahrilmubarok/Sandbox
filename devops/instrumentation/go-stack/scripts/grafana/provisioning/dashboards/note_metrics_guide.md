# Selecting the Right Metric Type

## **1. Counter**

**Definition:** Its value only increases (cannot go down).
Used for: **request count**, **error count**, **tasks processed**, etc.
**Example (Go):**

```go
var RequestTotal = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "http_requests_total",
        Help: "Total number of HTTP requests",
    },
    []string{"method", "status"},
)
```

### **PromQL Query Examples**

* Total requests:

  ```
  http_requests_total
  ```
* Requests per second (RPS):

  ```
  rate(http_requests_total[1m])
  ```
* Error rate:

  ```
  rate(http_requests_total{status=~"5.."}[5m])
  ```

---

## **2. Gauge**

**Definition:** Can go up or down.
Used for: **goroutine count**, **queue size**, **seat stock**, **memory usage**, etc.
**Example (Go):**

```go
var QueueSize = prometheus.NewGaugeVec(
    prometheus.GaugeOpts{
        Name: "queue_size",
        Help: "Current queue size",
    },
    []string{"queue"},
)
```

### **PromQL Query Examples**

* Current queue size:

  ```
  queue_size
  ```
* Queue size grouped by queue name:

  ```
  queue_size by (queue)
  ```
* Free seats example:

  ```
  seat_stock
  ```

---

## **3. Histogram**

**Definition:** Measures distribution and latency using buckets.
Used for: **request latency**, **processing time**, **payload size**, etc.
**Example (Go):**

```go
var RequestDuration = prometheus.NewHistogramVec(
    prometheus.HistogramOpts{
        Name:    "http_request_duration_seconds",
        Help:    "Request latency distribution",
        Buckets: prometheus.DefBuckets, // 0.005 → 10s
    },
    []string{"method", "path"},
)
```

### **PromQL Query Examples**

* Average latency:

  ```
  rate(http_request_duration_seconds_sum[5m])
  /
  rate(http_request_duration_seconds_count[5m])
  ```
* 95th percentile latency:

  ```
  histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m]))
  ```
* Latency per method:

  ```
  histogram_quantile(0.90, rate(http_request_duration_seconds_bucket[5m])) by (method, le)
  ```

---

## **4. Summary**

**Definition:** Similar to Histogram but calculates percentiles locally (client-side).
Used for: **latency when you need client-side quantiles**, **request sizes**, etc.
**Example (Go):**

```go
var RequestSummary = prometheus.NewSummaryVec(
    prometheus.SummaryOpts{
        Name:       "http_request_summary_seconds",
        Help:       "Summary of request durations",
        Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
    },
    []string{"method"},
)
```

### **PromQL Query Examples**

* 90th percentile (from summary):

  ```
  http_request_summary_seconds{quantile="0.9"}
  ```
* 99th percentile:

  ```
  http_request_summary_seconds{quantile="0.99"}
  ```
* Average latency:

  ```
  rate(http_request_summary_seconds_sum[5m])
  /
  rate(http_request_summary_seconds_count[5m])
  ```

---
