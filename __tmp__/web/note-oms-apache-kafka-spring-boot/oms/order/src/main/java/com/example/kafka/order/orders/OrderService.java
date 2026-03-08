package com.example.kafka.order.orders;


public interface OrderService {

    OrderResponse placeOrder(OrderRequest dto);

    OrderResponse updateOrder(String orderId, Order order);
}
