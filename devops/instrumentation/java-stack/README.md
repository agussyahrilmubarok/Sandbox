# Booking API Example

The Booking Course API is a distributed microservices system designed for managing course availability and participant bookings. The system includes telemetry, monitoring, logging, tracing, and service discovery for resilient production-ready architecture.

## Services

| Service         | Port         | Description                                     |
| --------------- | ------------ | ----------------------------------------------- |
| **booking-app** | 8080         | Handles booking operations for courses          |
| **member-app**  | 8081         | Manages member accounts and user profiles       |
| **course-app**  | 8082         | Manages course data, schedules, categories      |
| **postgres**    | 5432         | Primary relational database for all services    |
| **consul**      | 8500 / 8600  | Service discovery + DNS-based resolution        |
| **grafana**     | 3000         | Monitoring dashboards and metrics visualization |
| **loki**        | 3100         | Centralized log aggregation                     |
| **fluentbit**   | 24224        | Log forwarder to Loki                           |
| **prometheus**  | 9090         | Metrics collection for microservices            |
| **jaeger**      | 16686 / 4317 | Distributed tracing and performance analysis    |


## Run

1. Run all services:

```bash
docker compose up -d --build
```

2. Test all endpoint:

```bash
cd k6 && \ 
k6 run *
```

3. Explore

```bash
# Logs
http://localhost:3000

# Metrics
http://localhost:9090

# Traces
http://localhost:16686
```

## Technology Stack

- Grafana
- Loki
- FluentBit
- Prometheus
- Prometheus AlertManager
- Open Telemetry
- Jaeger
- Go
- Echo
- GORM
- Postgres
- Swagger

## Features

- Microservice-based architecture
- Centralized logging
- Distributed tracing
- Service monitoring and alerting
- API documentation available via Swagger UI