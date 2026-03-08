package com.example.kafka.core.events;

public class PaymentProcessedEvent {
    private String orderId;
    private String paymentId;

    public PaymentProcessedEvent() {
    }

    public PaymentProcessedEvent(String orderId, String paymentId) {
        this.orderId = orderId;
        this.paymentId = paymentId;
    }

    public String getOrderId() {
        return orderId;
    }

    public void setOrderId(String orderId) {
        this.orderId = orderId;
    }

    public String getPaymentId() {
        return paymentId;
    }

    public void setPaymentId(String paymentId) {
        this.paymentId = paymentId;
    }
}
