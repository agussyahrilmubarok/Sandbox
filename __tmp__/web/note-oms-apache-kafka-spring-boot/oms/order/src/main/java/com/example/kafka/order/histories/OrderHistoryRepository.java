package com.example.kafka.order.histories;

import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface OrderHistoryRepository extends JpaRepository<OrderHistory, String> {

    List<OrderHistory> findAllByOrderId(String orderId);

}
