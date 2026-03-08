package com.example.kafka.order.orders;

public enum OrderStatus {
    NEW,               // Order has been created
    PENDING,           // Waiting for confirmation or payment
    CONFIRMED,         // Order has been confirmed
    PROCESSING,        // Order is being processed
    SHIPPED,           // Order has been shipped
    DELIVERED,         // Order has been received by the customer
    COMPLETED,         // Order is completed successfully
    CANCELLED,         // Order has been cancelled
    REFUNDED,          // Payment has been refunded
    FAILED,            // Order processing failed
    ON_HOLD,           // Order is on hold due to some issue
    PARTIALLY_SHIPPED, // Only part of the order has been shipped
    AWAITING_PAYMENT,  // Waiting for payment from the customer
    AWAITING_FULFILLMENT // Waiting for fulfillment before shipping
}

