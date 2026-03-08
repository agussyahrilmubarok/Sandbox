package com.example.kafka.core.events;

import java.math.BigDecimal;

public class StockReservedEvent {
    private String orderId;
    private String productId;
    private Integer quantity;
    private BigDecimal price;

    public StockReservedEvent() {
    }

    public StockReservedEvent(String orderId, String productId, Integer quantity, BigDecimal price) {
        this.orderId = orderId;
        this.productId = productId;
        this.quantity = quantity;
        this.price = price;
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

    public Integer getQuantity() {
        return quantity;
    }

    public void setQuantity(Integer quantity) {
        this.quantity = quantity;
    }

    public BigDecimal getPrice() {
        return price;
    }

    public void setPrice(BigDecimal price) {
        this.price = price;
    }
}
