package com.example.kafka.core.events;

public class OrderApprovedEvent {
    private String orderId;

    public OrderApprovedEvent() {
    }

    public OrderApprovedEvent(String orderId) {
        this.orderId = orderId;
    }

    public String getOrderId() {
        return orderId;
    }

    public void setOrderId(String orderId) {
        this.orderId = orderId;
    }
}
