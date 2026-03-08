# Backend Engineer Technical Test

## Overview
This test evaluates your ability to design and implement a microservice using **Go**, the **Fiber** framework, and **Redis Streams** for event-driven communication.  

You will build a **flight search system** consisting of:
- **Main Service**: Exposes REST and SSE endpoints.
- **Provider Service**: Integrates with a mocked 3rd-party flight API.
- **Event-driven communication** using Redis Streams.

---

## Architecture Diagram
*(Insert your architecture diagram here)*

---

## Task Summary
You are expected to:
- Build two services: `main-service` and `provider-service`
- Integrate Redis Streams for inter-service communication
- Handle real-time data delivery using **Server-Sent Events (SSE)**
- Mock integration with an external flight API

---

## Tech Stack
- **Language**: Go
- **Web Framework**: Fiber
- **Messaging Queue**: Redis Streams
- **API Integration**: Mock 3rd Party Flight API
- **Streaming**: Server-Sent Events (SSE)

---

## Requirements

### Main Service
- Expose REST API endpoint to initiate flight searches
- Publish request to `flight.search.requested` Redis stream
- Subscribe to `flight.search.results` stream
- Stream real-time results to the client via SSE
- Handle concurrent SSE clients, connection lifecycle, and cleanup
- Log all requests and responses (structured logging preferred)

### Provider Service
- Consume messages from `flight.search.requested`
- Simulate calling a mock flight API to fetch search results
- Publish formatted results to `flight.search.results`
- Use Redis Stream consumer groups with acknowledgments (`XREADGROUP`, `XACK`)

---

## API Specification

### `POST /api/flights/search`
**Request Body:**
```json
{
  "from": "CGK",
  "to": "DPS",
  "date": "2025-07-10",
  "passengers": 2
}
````

**Response:**

```json
{
  "success": true,
  "message": "Search request submitted",
  "data": {
    "search_id": "uuid",
    "status": "processing"
  }
}
```

---

### `GET /api/flights/search/{search_id}/stream`

**Headers:**

```
Content-Type: text/event-stream
Cache-Control: no-cache
Connection: keep-alive
```

**Example SSE Events:**

```
data: {"search_id": "uuid", "status": "processing", "results": []}
data: {"search_id": "uuid", "status": "completed", "results": [...]}
data: {"search_id": "uuid", "status": "completed", "total_results": 1}
```

---

## Message Structure

### `flight.search.requested`

```json
{
  "search_id": "uuid",
  "from": "CGK",
  "to": "DPS",
  "date": "2025-07-10",
  "passengers": 2
}
```

### `flight.search.results`

```json
{
  "search_id": "uuid",
  "status": "completed",
  "results": [
    { "flight result object" }
  ]
}
```

---

## Deliverables

* Working Go code for both services (`main-service`, `provider-service`)
* Docker Compose file for easy setup
* **README.md** including:

  * Setup instructions
  * API usage (e.g., Postman or curl)
  * Explanation of architecture, design decisions, and trade-offs
* Handle error cases (e.g., invalid input, Redis connection failure)
* Structured logging for observability
* Use UUIDs for all `search_id`

---

## Bonus Points

* Unit/integration tests
* Clean architecture (separation of handler, service, repository)
* Use Redis consumer groups with `XREADGROUP`
* Graceful shutdown handling (cleanup resources)
* Metrics or tracing (OpenTelemetry, Prometheus, etc.)

---

## Duration

Estimated time to complete: **24 hours**

---

**Good Luck!**
Document any assumptions you make.
