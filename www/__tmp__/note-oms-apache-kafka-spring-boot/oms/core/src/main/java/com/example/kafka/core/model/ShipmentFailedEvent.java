package com.example.kafka.core.model;

public class ShipmentFailedEvent {
    private String orderId;
    private String reason;

    public ShipmentFailedEvent() {
    }

    public ShipmentFailedEvent(String orderId, String reason) {
        this.orderId = orderId;
        this.reason = reason;
    }

    public String getOrderId() {
        return orderId;
    }

    public void setOrderId(String orderId) {
        this.orderId = orderId;
    }

    public String getReason() {
        return reason;
    }

    public void setReason(String reason) {
        this.reason = reason;
    }
}
