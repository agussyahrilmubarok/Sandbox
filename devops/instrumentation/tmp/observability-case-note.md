# 🧭 Booking Service Observability & Alerting

This document describes the **observability strategy**, **key metrics**, and **alerting topics** for the `booking` business service.  
The goal is to ensure **system reliability**, **business continuity**, and **early detection of booking-related issues** that may impact user experience.

---

## 📈 Overview

The **Booking Service** interacts with:
- The **Member Service** (to validate members)
- The **Course Service** (to reserve and release course seats)
- The **Database Store** (to persist booking data)

To maintain healthy operations, observability is built with:
- **OpenTelemetry** for tracing and distributed context
- **Prometheus metrics** for monitoring
- **Zerolog** for structured logging
- **Grafana** for visualization and alerting

---

## ⚙️ Key Metrics

| Metric Name | Type | Description |
|--------------|------|-------------|
| `failed_bookings_total` | Counter | Total number of failed booking attempts |
| `successful_bookings_total` | Counter | Total number of successful bookings |
| `booking_latency_seconds` | Histogram | Duration (in seconds) of the booking process |
| `booking_status` | CounterVec | Booking counts grouped by status (e.g., confirmed, failed) |

---

## 🚨 Primary Alert Topic — **Booking Failure Rate Spike**

### 🎯 Purpose
Detects a significant increase in booking failures, which may indicate downstream service issues, course reservation bottlenecks, or data consistency problems.

### 🔍 Description
This alert monitors the **failure rate** (failed bookings vs total bookings) and triggers if more than **20%** of all booking attempts fail within a 10-minute window.

### ⚙️ Prometheus Rule

```yaml
- alert: HighBookingFailureRate
  expr: |
    (
      rate(failed_bookings_total[5m])
      /
      (rate(successful_bookings_total[5m]) + rate(failed_bookings_total[5m]))
    ) > 0.2
  for: 10m
  labels:
    severity: critical
    topic: business-booking
  annotations:
    summary: "Booking Failure Rate Spike (>20%)"
    description: "More than 20% of booking attempts have failed in the last 10 minutes. Investigate member, course, or database dependencies."
````

---

## ⚠️ Additional Recommended Alerts

### 🕓 **High Booking Latency**

Detects degraded performance when booking API response times increase.

```yaml
- alert: BookingLatencyHigh
  expr: histogram_quantile(0.95, sum(rate(booking_latency_seconds_bucket[5m])) by (le)) > 2
  for: 5m
  labels:
    severity: warning
    topic: business-booking
  annotations:
    summary: "High Booking Latency"
    description: "95th percentile of booking latency exceeds 2 seconds over the last 5 minutes."
```

---

### 🧩 **Course Reservation Errors Spike**

Detects recurring failures when reserving or releasing course seats — often caused by capacity exhaustion or downstream service instability.

```yaml
- alert: CourseReservationErrors
  expr: increase(failed_bookings_total[10m]) > 10
  for: 10m
  labels:
    severity: warning
    topic: business-booking
  annotations:
    summary: "Course Reservation Errors Spike"
    description: "More than 10 booking failures occurred in the last 10 minutes, possibly due to full or unavailable courses."
```

---

## 🧠 Observability Best Practices

1. **Trace everything** — Ensure all booking flows (`Booking`, `BookingV2`) are traced using OpenTelemetry spans:

   * `service.Booking`
   * `service.BookingV2`
   * `client.FindMemberByCode`
   * `client.ReserveCourseByCode`
   * `store.Save`

2. **Use contextual logging** — Every log entry should include:

   * `booking_id`, `member_code`, `course_code`
   * Error context and trace ID

3. **Correlate metrics and traces** — Use the same trace ID to correlate Prometheus metrics with distributed traces in Grafana Tempo or Jaeger.

4. **Single Alert Topic** — All alerts related to the booking flow use:

   ```yaml
   topic: business-booking
   ```

   This allows unified monitoring and notification through a single alert channel (e.g., Slack `#booking-alerts`).

---

## 📊 Example Grafana Dashboard Sections

| Section                  | Description                                     |
| ------------------------ | ----------------------------------------------- |
| **Booking Overview**     | Shows total, successful, and failed bookings    |
| **Failure Rate (%)**     | Visualizes real-time error ratio                |
| **Latency Distribution** | P95, P99 latency from `booking_latency_seconds` |
| **Dependency Health**    | Member API and Course API success rates         |
| **Recent Alerts**        | Active alerts from `business-booking` topic     |

---

## ✅ Summary

| Area                  | Metric / Alert            | Purpose                                           |
| --------------------- | ------------------------- | ------------------------------------------------- |
| **Reliability**       | `HighBookingFailureRate`  | Detect user-impacting booking failures            |
| **Performance**       | `BookingLatencyHigh`      | Identify degraded booking response times          |
| **Dependency Health** | `CourseReservationErrors` | Detect issues with course capacity or service     |
| **Unified Alerting**  | Topic: `business-booking` | Simplify alert routing and observability tracking |

---

> **Maintainer Note:**
> Any new metric or alert related to the booking process should include the `topic="business-booking"` label for unified monitoring and alert correlation.

```

---

It visualizes:

* Booking success/failure rates
* Booking latency (p95)
* Course availability
* Active alerts from the topic `business-booking`

This dashboard assumes your Prometheus metrics follow names like:
`failed_bookings_total`, `successful_bookings_total`, `booking_latency_seconds_bucket`, etc.

---

```json
{
  "id": null,
  "uid": "business-booking-dashboard",
  "title": "📊 Business Booking Observability",
  "tags": ["booking", "observability", "business"],
  "timezone": "browser",
  "schemaVersion": 38,
  "version": 1,
  "refresh": "30s",
  "panels": [
    {
      "type": "row",
      "title": "Booking Overview",
      "collapsed": false,
      "gridPos": { "h": 1, "w": 24, "x": 0, "y": 0 }
    },
    {
      "id": 1,
      "title": "Total Bookings",
      "type": "stat",
      "targets": [
        {
          "expr": "increase(successful_bookings_total[1h]) + increase(failed_bookings_total[1h])",
          "legendFormat": "Bookings (1h)"
        }
      ],
      "gridPos": { "h": 6, "w": 6, "x": 0, "y": 1 },
      "fieldConfig": {
        "defaults": { "unit": "none", "color": { "mode": "palette-classic" } }
      }
    },
    {
      "id": 2,
      "title": "Booking Failure Rate (%)",
      "type": "timeseries",
      "targets": [
        {
          "expr": "(rate(failed_bookings_total[5m]) / (rate(successful_bookings_total[5m]) + rate(failed_bookings_total[5m]))) * 100",
          "legendFormat": "Failure Rate (%)"
        }
      ],
      "gridPos": { "h": 8, "w": 12, "x": 6, "y": 1 },
      "fieldConfig": {
        "defaults": {
          "unit": "percent",
          "min": 0,
          "max": 100,
          "thresholds": {
            "mode": "absolute",
            "steps": [
              { "color": "green", "value": null },
              { "color": "yellow", "value": 10 },
              { "color": "red", "value": 20 }
            ]
          }
        }
      }
    },
    {
      "id": 3,
      "title": "Booking Latency (P95)",
      "type": "timeseries",
      "targets": [
        {
          "expr": "histogram_quantile(0.95, sum(rate(booking_latency_seconds_bucket[5m])) by (le))",
          "legendFormat": "P95 Booking Latency"
        }
      ],
      "gridPos": { "h": 8, "w": 6, "x": 18, "y": 1 },
      "fieldConfig": {
        "defaults": {
          "unit": "s",
          "thresholds": {
            "mode": "absolute",
            "steps": [
              { "color": "green", "value": null },
              { "color": "yellow", "value": 1.5 },
              { "color": "red", "value": 2 }
            ]
          }
        }
      }
    },
    {
      "type": "row",
      "title": "Course & Member Health",
      "collapsed": false,
      "gridPos": { "h": 1, "w": 24, "x": 0, "y": 9 }
    },
    {
      "id": 4,
      "title": "Available vs Reserved Courses",
      "type": "timeseries",
      "targets": [
        { "expr": "available_courses", "legendFormat": "Available" },
        { "expr": "reserved_courses", "legendFormat": "Reserved" }
      ],
      "gridPos": { "h": 8, "w": 12, "x": 0, "y": 10 },
      "fieldConfig": { "defaults": { "unit": "none" } }
    },
    {
      "id": 5,
      "title": "Course Reservation Errors (10m)",
      "type": "stat",
      "targets": [
        {
          "expr": "increase(failed_bookings_total[10m])",
          "legendFormat": "Failed Bookings (10m)"
        }
      ],
      "gridPos": { "h": 8, "w": 6, "x": 12, "y": 10 },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "thresholds",
            "seriesBy": "last"
          },
          "thresholds": {
            "mode": "absolute",
            "steps": [
              { "color": "green", "value": null },
              { "color": "yellow", "value": 5 },
              { "color": "red", "value": 10 }
            ]
          }
        }
      }
    },
    {
      "type": "row",
      "title": "Alerting & Incident Overview",
      "collapsed": false,
      "gridPos": { "h": 1, "w": 24, "x": 0, "y": 18 }
    },
    {
      "id": 6,
      "title": "Active Alerts — business-booking",
      "type": "table",
      "targets": [
        {
          "expr": "ALERTS{topic=\"business-booking\", alertstate=\"firing\"}",
          "legendFormat": "{{alertname}}"
        }
      ],
      "gridPos": { "h": 8, "w": 24, "x": 0, "y": 19 },
      "fieldConfig": {
        "defaults": {
          "color": { "mode": "thresholds" },
          "thresholds": {
            "mode": "absolute",
            "steps": [
              { "color": "green", "value": null },
              { "color": "yellow", "value": 1 },
              { "color": "red", "value": 2 }
            ]
          }
        }
      },
      "options": { "showHeader": true }
    }
  ],
  "templating": { "list": [] },
  "time": { "from": "now-1h", "to": "now" },
  "annotations": {
    "list": [
      {
        "builtIn": 1,
        "datasource": "-- Grafana --",
        "enable": true,
        "hide": false,
        "iconColor": "rgba(255, 96, 96, 1)",
        "name": "Annotations & Alerts",
        "type": "dashboard"
      }
    ]
  }
}
```

---

## 💡 How to Use This Dashboard

### 1. **Import into Grafana**

1. Open your Grafana UI → **Dashboards → New → Import**
2. Paste the JSON above or upload it as a `.json` file.
3. Select your **Prometheus data source**.
4. Click **Import**.

### 2. **Adjust Metric Names (if different)**

If your metric names differ slightly (e.g. `booking_failed_total` instead of `failed_bookings_total`), edit each panel query accordingly.

### 3. **Alert Correlation**

This dashboard automatically surfaces alerts from Prometheus using:

```promql
ALERTS{topic="business-booking", alertstate="firing"}
```

So make sure your alert rules include:

```yaml
labels:
  topic: business-booking
```

### 4. **Recommended Slack/Webhook Integration**

Route `topic=business-booking` alerts to a single notification channel such as:

* Slack: `#booking-alerts`
* Email: `booking-alerts@yourcompany.com`
* PagerDuty or Opsgenie integration for critical severity

---

## ✅ Summary

This dashboard provides a **business-level overview** of the booking process, combining metrics, traces, and alerting into one visual system.

| Section          | Key Metric / Alert       | Purpose                   |
| ---------------- | ------------------------ | ------------------------- |
| Booking Overview | Success & Failure rates  | Business health indicator |
| Latency          | P95 latency              | Performance monitoring    |
| Course Health    | Available vs Reserved    | Dependency visibility     |
| Active Alerts    | `topic=business-booking` | Real-time issue tracking  |

---
