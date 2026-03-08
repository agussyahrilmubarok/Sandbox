package com.example.kafka.core.commands;

public class OrderRejectCommand {
    private String orderId;

    public OrderRejectCommand() {
    }

    public OrderRejectCommand(String orderId) {
        this.orderId = orderId;
    }

    public String getOrderId() {
        return orderId;
    }

    public void setOrderId(String orderId) {
        this.orderId = orderId;
    }
}
