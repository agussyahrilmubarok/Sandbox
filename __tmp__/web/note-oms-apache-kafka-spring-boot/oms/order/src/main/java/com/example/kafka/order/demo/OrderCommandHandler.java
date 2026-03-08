package com.example.kafka.order.demo;

import com.example.kafka.core.commands.OrderApproveCommand;
import com.example.kafka.core.commands.OrderRejectCommand;
import com.example.kafka.core.events.OrderApprovedEvent;
import com.example.kafka.order.histories.OrderHistoryRepository;
import com.example.kafka.order.orders.Order;
import com.example.kafka.order.orders.OrderRepository;
import com.example.kafka.order.orders.OrderStatus;
import lombok.RequiredArgsConstructor;
import org.springframework.kafka.annotation.KafkaHandler;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.stereotype.Component;

@Component
@KafkaListener(topics = {"orders-commands"})
@RequiredArgsConstructor
public class OrderCommandHandler {

    private final KafkaTemplate<String, Object> kafkaTemplate;
    private final OrderRepository orderRepository;
    private final OrderHistoryRepository orderHistoryRepository;

    @KafkaHandler
    public void orderApproveCommand(@Payload OrderApproveCommand command) {
        Order order = approveOrder(command.getOrderId());
        if (order != null) {
            OrderApprovedEvent event = new OrderApprovedEvent();
            event.setOrderId(order.getId());
            kafkaTemplate.send("orders-events", event);
        }
    }

    @KafkaHandler
    public void orderRejectCommand(@Payload OrderRejectCommand command) {
        Order order = rejectOrder(command.getOrderId());
        // TODO
    }

    private Order approveOrder(String orderId) {
        Order order = orderRepository.findById(orderId).orElse(null);
        if (order == null) {
            return null;
        } else {
            order.setStatus(OrderStatus.CONFIRMED);
            orderRepository.save(order);
            OrderApprovedEvent event = new OrderApprovedEvent();
            event.setOrderId(orderId);
            return order;
        }
    }

    private Order rejectOrder(String orderId) {
        Order order = orderRepository.findById(orderId).orElse(null);
        if (order == null) {
            return null;
        } else {
            order.setStatus(OrderStatus.CANCELLED);
            orderRepository.save(order);
            OrderApprovedEvent event = new OrderApprovedEvent();
            event.setOrderId(orderId);
            return order;
        }
    }

}
