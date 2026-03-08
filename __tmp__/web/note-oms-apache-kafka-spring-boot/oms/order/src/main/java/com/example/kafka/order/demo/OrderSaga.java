package com.example.kafka.order.demo;

import com.example.kafka.core.commands.PaymentProcessCommand;
import com.example.kafka.core.commands.ProductReserveCommand;
import com.example.kafka.core.commands.ProductRollbackCommand;
import com.example.kafka.core.events.*;
import com.example.kafka.order.histories.OrderHistoryService;
import com.example.kafka.order.orders.Order;
import com.example.kafka.order.orders.OrderService;
import com.example.kafka.order.orders.OrderStatus;
import lombok.RequiredArgsConstructor;
import org.springframework.kafka.annotation.KafkaHandler;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.stereotype.Component;

@Component
@KafkaListener(topics = {"orders-events", "products-events", "payments-events"})
@RequiredArgsConstructor
public class OrderSaga {

    private final KafkaTemplate<String, Object> kafkaTemplate;
    private final OrderService orderService;
    private final OrderHistoryService orderHistoryService;

    @KafkaHandler
    public void orderCreatedEvent(@Payload OrderCreatedEvent event) {
        orderHistoryService.add(event.getOrderId(), OrderStatus.NEW);

        ProductReserveCommand productReserveCommand = new ProductReserveCommand();
        productReserveCommand.setOrderId(event.getOrderId());
        productReserveCommand.setProductId(event.getProductId());
        productReserveCommand.setQuantity(event.getQuantity());
        kafkaTemplate.send("products-commands", productReserveCommand);
    }

    @KafkaHandler
    public void stockReservedEvent(@Payload StockReservedEvent event) {
        Order order = new Order();
        order.setPrice(event.getPrice());
        order.setStatus(OrderStatus.PROCESSING);
        orderService.updateOrder(event.getOrderId(), order);
        orderHistoryService.add(event.getOrderId(), OrderStatus.PROCESSING);

        PaymentProcessCommand paymentProcessCommand = new PaymentProcessCommand();
        paymentProcessCommand.setOrderId(event.getOrderId());
        paymentProcessCommand.setProductId(event.getProductId());
        paymentProcessCommand.setQuantity(event.getQuantity());
        paymentProcessCommand.setPrice(event.getPrice());
        kafkaTemplate.send("payments-commands", paymentProcessCommand);
    }

    @KafkaHandler
    public void stockNotAvailable(@Payload StockNotAvailableEvent event) {
        Order order = new Order();
        order.setStatus(OrderStatus.FAILED);
        orderService.updateOrder(event.getOrderId(), order);
        orderHistoryService.add(event.getOrderId(), OrderStatus.FAILED);
    }

    @KafkaHandler
    public void paymentProcessedEvent(@Payload PaymentProcessedEvent event) {
        Order order = new Order();
        order.setStatus(OrderStatus.CONFIRMED);
        orderService.updateOrder(event.getOrderId(), order);
        orderHistoryService.add(event.getOrderId(), OrderStatus.CONFIRMED);
    }

    @KafkaHandler
    public void paymentFailedEvent(@Payload PaymentFailedEvent event) {
        Order order = new Order();
        order.setStatus(OrderStatus.CANCELLED);
        orderService.updateOrder(event.getOrderId(), order);
        orderHistoryService.add(event.getOrderId(), OrderStatus.CANCELLED);

        ProductRollbackCommand productRollbackCommand = new ProductRollbackCommand();
        productRollbackCommand.setOrderId(event.getOrderId());
        productRollbackCommand.setProductId(event.getProductId());
        productRollbackCommand.setQuantity(event.getQuantity());
        kafkaTemplate.send("products-commands", productRollbackCommand);
    }

}
