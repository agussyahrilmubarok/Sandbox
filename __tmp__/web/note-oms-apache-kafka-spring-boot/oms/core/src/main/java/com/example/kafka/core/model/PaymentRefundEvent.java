package com.example.kafka.core.model;

public class PaymentRefundEvent {
    private String orderId;
    private Double amount;

    public PaymentRefundEvent() {
    }

    public PaymentRefundEvent(String orderId, Double amount) {
        this.orderId = orderId;
        this.amount = amount;
    }

    public String getOrderId() {
        return orderId;
    }

    public void setOrderId(String orderId) {
        this.orderId = orderId;
    }

    public Double getAmount() {
        return amount;
    }

    public void setAmount(Double amount) {
        this.amount = amount;
    }
}
