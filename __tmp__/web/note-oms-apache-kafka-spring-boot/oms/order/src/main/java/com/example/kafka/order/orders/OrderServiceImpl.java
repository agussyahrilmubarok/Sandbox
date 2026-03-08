package com.example.kafka.order.orders;

import com.example.kafka.order.demo.OrderProducer;
import com.example.kafka.order.histories.OrderHistoryRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import java.util.UUID;

@Service
@Slf4j
@RequiredArgsConstructor
public class OrderServiceImpl implements OrderService {

    private final OrderRepository orderRepository;
    private final OrderHistoryRepository orderHistoryRepository;
    private final OrderProducer orderProducer;

    @Override
    public OrderResponse placeOrder(OrderRequest dto) {
        Order order = new Order();
        order.setId(UUID.randomUUID().toString());
        order.setProductId(dto.getProductId());
        order.setQuantity(dto.getTotal());
        order.setStatus(OrderStatus.NEW);
        Order savedOrder = orderRepository.save(order);

        orderProducer.sendOrderCreatedEvent(order);

        return OrderResponse.builder()
                .id(savedOrder.getId())
                .productId(savedOrder.getProductId())
                .total(savedOrder.getQuantity())
                .price(order.getPrice())
                .status(savedOrder.getStatus().toString())
                .dateCreated(savedOrder.getDateCreated())
                .lastUpdated(savedOrder.getLastUpdated())
                .build();
    }

    @Override
    public OrderResponse updateOrder(String orderId, Order orderUpdate) {
        Order order = orderRepository.findById(orderId).orElseThrow();
        if (orderUpdate.getPrice() != null) {
            order.setPrice(orderUpdate.getPrice());
        }

        if (orderUpdate.getStatus() != null) {
            order.setStatus(orderUpdate.getStatus());
        }

        Order savedOrder = orderRepository.save(order);
        return OrderResponse.builder()
                .id(savedOrder.getId())
                .productId(savedOrder.getProductId())
                .total(savedOrder.getQuantity())
                .price(order.getPrice())
                .status(savedOrder.getStatus().toString())
                .dateCreated(savedOrder.getDateCreated())
                .lastUpdated(savedOrder.getLastUpdated())
                .build();
    }

}
