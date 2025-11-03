package com.example.kafka.order.histories;

import com.example.kafka.order.orders.OrderStatus;

import java.util.List;

public interface OrderHistoryService {

    List<OrderHistory> findAll();

    List<OrderHistory> findAllByOrderId(String orderId);

    void add(String orderId, OrderStatus status);

}
