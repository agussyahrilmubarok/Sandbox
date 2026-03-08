package com.example.kafka.core.events;

public class StockNotAvailableEvent {
    private String orderId;
    private String productId;

    public StockNotAvailableEvent() {
    }

    public StockNotAvailableEvent(String orderId, String productId) {
        this.orderId = orderId;
        this.productId = productId;
    }

    public String getOrderId() {
        return orderId;
    }

    public void setOrderId(String orderId) {
        this.orderId = orderId;
    }

    public String getProductId() {
        return productId;
    }

    public void setProductId(String productId) {
        this.productId = productId;
    }
}
