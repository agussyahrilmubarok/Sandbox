package com.example.kafka.order.orders;

import org.springframework.data.jpa.repository.JpaRepository;


public interface OrderRepository extends JpaRepository<Order, String> {

    boolean existsByIdIgnoreCase(String id);

}
