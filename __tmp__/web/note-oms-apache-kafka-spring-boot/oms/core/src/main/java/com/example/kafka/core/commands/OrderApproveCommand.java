package com.example.kafka.core.commands;

public class OrderApproveCommand {
    private String orderId;

    public OrderApproveCommand() {
    }

    public OrderApproveCommand(String orderId) {
        this.orderId = orderId;
    }

    public String getOrderId() {
        return orderId;
    }

    public void setOrderId(String orderId) {
        this.orderId = orderId;
    }
}
