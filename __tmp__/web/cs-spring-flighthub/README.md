# Flighthub

This project provides a Spring Boot API for managing user authentication, airport management, flight management, and
flight search functionalities. It includes advanced logging and security features.

## Feature

### 🔐 User Authentication

- **User Registration**: Register users with an email and password.
- **User Login**: Authenticate users with email and password, generating access and refresh tokens.
- **Token Refresh**: Refresh expired access tokens using refresh tokens.
- **User Logout**: Invalidate tokens to log users out.

### 🛫 Airport Management

- **Create Airports**: Accessible only to `ADMIN` users.
- **Retrieve Airports**: Paginated airport list retrieval, accessible to both `ADMIN` and `USER` roles.
- **Retrieve Specific Airport**: Search by airport ID, accessible to both `ADMIN` and `USER` roles.
- **Update Airports**: Modify an existing airport by ID, accessible only to `ADMIN` users.
- **Delete Airports**: Remove airports by ID, accessible only to `ADMIN` users.

### ✈️ Flight Management

- **Create Flights**: Accessible only to `ADMIN` users.
- **Retrieve Flights**: Paginated flight list retrieval, accessible to both `ADMIN` and `USER` roles.
- **Retrieve Specific Flight**: Search by flight ID, accessible to both `ADMIN` and `USER` roles.
- **Update Flights**: Modify an existing flight by ID, accessible only to `ADMIN` users.
- **Delete Flights**: Remove flights by ID, accessible only to `ADMIN` users.

### 🔍 Flight Search

- **Search Flights**: Search for flights based on departure, arrival airports, and departure date. Accessible to both
  `ADMIN` and `USER` roles. (One Way - Round Trip)

### 🔒 Security

- **User Roles**: `ADMIN` and `USER` roles implemented using Spring Security.
- **JWT Authentication**: Secure endpoints with JSON Web Tokens.

### 📋 Logging

- **Custom Logging Aspect**: Logs details of REST controller method calls and exceptions, including HTTP request and
  response details.

### 📘 Explore REST APIs

#### Endpoints Summary

| Method | URL                                         | Description                | Request Body             | Path Variable | Response                                                |
|--------|---------------------------------------------|----------------------------|--------------------------|---------------|---------------------------------------------------------|
| POST   | `/api/v1/authentication/user/register`      | Register for Admin or User | `RegisterRequest`        | -             | `CustomResponse<Void>`                                  |
| POST   | `/api/v1/authentication/user/login`         | Login for Admin or User    | `LoginRequest`           | -             | `CustomResponse<TokenResponse>`                         |
| POST   | `/api/v1/authentication/user/refresh-token` | Refresh Token              | `TokenRefreshRequest`    | -             | `CustomResponse<TokenResponse>`                         |
| POST   | `/api/v1/authentication/user/logout`        | Logout for Admin or User   | `TokenInvalidateRequest` | -             | `CustomResponse<Void>`                                  |
| POST   | `/api/v1/airports`                          | Create a new Airport       | `CreateAirportRequest`   | -             | `CustomResponse<String>`                                |
| GET    | `/api/v1/airports`                          | Get all Airports           | `AirportPagingRequest`   | -             | `CustomResponse<CustomPagingResponse<AirportResponse>>` |
| GET    | `/api/v1/airports/{id}`                     | Get Airport by ID          | -                        | `UUID`        | `CustomResponse<AirportResponse>`                       |
| PUT    | `/api/v1/airports/{id}`                     | Update an Airport          | `UpdateAirportRequest`   | `UUID`        | `CustomResponse<AirportResponse>`                       |
| DELETE | `/api/v1/airports/{id}`                     | Delete an Airport          | -                        | `UUID`        | `CustomResponse<String>`                                |
| POST   | `/api/v1/flights`                           | Create a new Flight        | `CreateFlightRequest`    | -             | `CustomResponse<String>`                                |
| GET    | `/api/v1/flights`                           | Get all Flights            | `AirportPagingRequest`   | -             | `CustomResponse<CustomPagingResponse<FlightResponse>>`  |
| GET    | `/api/v1/flights/{id}`                      | Get Flight by ID           | -                        | `UUID`        | `CustomResponse<FlightResponse>`                        |
| PUT    | `/api/v1/flights/{id}`                      | Update a Flight            | `UpdateFlightRequest`    | `UUID`        | `CustomResponse<FlightResponse>`                        |
| DELETE | `/api/v1/flights/{id}`                      | Delete a Flight            | -                        | `UUID`        | `CustomResponse<String>`                                |
| POST   | `/api/v1/flights/search`                    | Search for Flights         | `SearchFlightRequest`    | -             | `CustomResponse<CustomPagingResponse<FlightResponse>>`  |

## 🚀 Technologies Used

This project leverages a modern and robust tech stack:

- **Java 21**
- **Spring Boot 3.0**
- **RESTful API Design**
- **MapStruct** – Object mapping made easy
- **OpenAPI (Swagger)** – API documentation and testing
- **Maven** – Build and dependency management
- **JUnit 5** – Unit testing framework
- **Mockito** – Mocking framework for unit tests
- **Integration Tests**
- **Docker** – Containerization
- **Docker Compose** – Multi-container orchestration
- **CI/CD** – GitHub Actions & Jenkins pipelines
- **Postman** – API testing
- **TestContainers** – Integration testing with real dependencies
- **MongoDB** – NoSQL database
- **Prometheus** – Monitoring
- **Grafana** – Metrics visualization
- **Kubernetes** – Container orchestration
- **JaCoCo** – Code coverage reporting

## 🚀 Explore

### 🧪 Postman

Import the Postman collection located in the `postman_collection` folder to test the API endpoints.

### 📄 Prerequisites

Create a `.env` file in the root directory and define the following environment variables:

```env
MONGO_DB_NAME=flightdatabase
MONGO_DB_HOST=localhost
MONGO_DB_PORT=27017
```

### 🔍 Open API (Swagger)

Access API docs at:

```bash
http://localhost:1133/swagger-ui/index.html
```

### 📊 JaCoCo (Test Report)

Run this command:

```bash
mvn clean install
```

Then open:

```bash
target/site/jacoco/index.html
```

### 🏁 How to Run

#### 🔧 With Maven

```bash
mvn clean install
mvn spring-boot:run
```

#### 🐳 With Docker

```bash
docker-compose up -d --build
```

### 📈 Monitoring

- Prometheus: http://localhost:9090
- Grafana: http://localhost:3000
  Username: admin / Password: admin

### ☸️ Kubernetes Run

To build and run the application with `Maven`, please follow the directions shown below;

- Start Minikube

```sh
$ minikube start
```

- Open Minikube Dashboard

```sh
$ minikube dashboard
```

- To deploy the application on Kubernetes, apply the Kubernetes configuration file underneath k8s folder

```sh
$ kubectl apply -f k8s
```

- To open Prometheus, click tunnel url link provided by the command shown below to reach out Prometheus

```sh
minikube service prometheus-service
```

- To open Grafana, click tunnel url link provided by the command shown below to reach out Prometheus

```sh
minikube service grafana-service
```

- Define prometheus data source url, use this link shown below

```bash
http://prometheus-service.default.svc.cluster.local:9090
```

- 

## References

