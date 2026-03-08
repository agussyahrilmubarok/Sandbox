# Demo Apache Kafka Spring Boot

## Overview

This repository demonstrates how to integrate **Apache Kafka** with **Spring Boot**, covering key Kafka concepts, configurations, and advanced use cases, including **transactions, idempotency, error handling, integration testing, and the Saga design pattern**.

## Insight

1. RESTful API on product `:8081`, order `:8082`, payment `:8083` service.

2. Implement Event-Driven Architecture SAGA Pattern.

3. How to test it
    1. Create Order
        - Case Product Not Found Or Less then Order Quantity
        - Case Product Available
    2. Payment Order
        - Case Success Payment
        - Case Fail Payment

## Features

- Kafka Producer & Consumer with Spring Boot
- Transactional Messaging in Kafka
- Handling Serialization and Deserialization Errors
- Dead Letter Topic (DLT) Handling
- Consumer Retries & Exception Handling
- Idempotency in Producer and Consumer
- Kafka Integration Testing
- Saga Pattern Implementation with Kafka

## Contents

1. **Introduction to Apache Kafka**  
   Overview of Apache Kafka, its architecture, and core components.
2. **Apache Kafka Broker's**  
   Understanding Kafka Brokers, Zookeeper, and Cluster Setup.
3. **Apache Kafka CLI Topic's**  
   Managing Kafka Topics using Command-Line Interface.
4. **Apache Kafka CLI Producer**  
   Producing messages to Kafka topics using Kafka CLI.
5. **Apache Kafka CLI Consumer**  
   Consuming messages from Kafka topics using Kafka CLI.
6. **Apache Kafka Producer in Spring Boot**  
   Implementing Kafka Producers in Spring Boot applications.
7. **Apache Kafka Producer Acknowledgement**  
   Configuring Producer Acknowledgements for reliability.
8. **Idempotent Producer in Kafka**  
   Ensuring exactly-once message delivery using Idempotent Producers.
9. **Apache Kafka Consumer in Spring Boot**  
   Implementing Kafka Consumers in Spring Boot applications.
10. **Apache Kafka Consumer Handle Deserializer Errors**  
    Handling serialization and deserialization exceptions.
11. **Apache Kafka Consumer Dead Letter Topic (DLT)**  
    Using DLT for handling poisoned messages.
12. **Apache Kafka Consumer Exceptions and Retries**  
    Implementing retry mechanisms for transient failures.
13. **Apache Kafka Consumer Multiple Consumers in Group**  
    Understanding consumer groups and load balancing.
14. **Idempotent Consumer in Kafka**  
    Implementing idempotency to prevent duplicate processing.
15. **Apache Kafka Transactions**  
    Implementing transactional messaging with Kafka.
16. **Apache Kafka Database Transactions**  
    Synchronizing Kafka messages with database transactions.
17. **Integration Testing in Apache Kafka Producer**  
    Writing and executing integration tests for Kafka Producers.
18. **Integration Testing in Apache Kafka Consumer**  
    Writing and executing integration tests for Kafka Consumers.
19. **Saga Design Pattern with Apache Kafka**  
    Implementing distributed transactions using Saga (Choreography & Orchestration).
20. **Saga Design Pattern with Apache Kafka Compensating Transactions**  
    Handling rollback scenarios using compensating transactions.

## Getting Started

### Prerequisites

- Java 11 or later
- Apache Kafka & Zookeeper
- Docker (Optional)
- Spring Boot Framework

### Installation & Running Kafka

1. **Start Zookeeper & Kafka**
   ```bash
   bin/zookeeper-server-start.sh config/zookeeper.properties &
   bin/kafka-server-start.sh config/server.properties &
   ```
2. **Verify Kafka is running**
   ```bash
   bin/kafka-topics.sh --list --bootstrap-server localhost:9092
   ```

### Running Spring Boot Application

```bash
mvn spring-boot:run
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License.

## References

0. [Apache Kafka for Event-Driven Spring Boot Microservices](https://www.udemy.com/course/apache-kafka-for-spring-boot-microservices)
1. [Apache Kafka Documentation](https://kafka.apache.org/documentation/)
2. [Spring Kafka Reference Guide](https://docs.spring.io/spring-kafka/docs/current/reference/html/)
3. [Confluent Kafka Documentation](https://docs.confluent.io/current/kafka.html)
