package com.example.kafka.core.model;

public class ShipmentInitiatedEvent {
    private String orderId;

    public ShipmentInitiatedEvent() {
    }

    public ShipmentInitiatedEvent(String orderId) {
        this.orderId = orderId;
    }

    public String getOrderId() {
        return orderId;
    }

    public void setOrderId(String orderId) {
        this.orderId = orderId;
    }
}
