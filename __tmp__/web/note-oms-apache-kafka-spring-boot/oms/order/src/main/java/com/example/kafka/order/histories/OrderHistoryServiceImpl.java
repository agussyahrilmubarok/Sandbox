package com.example.kafka.order.histories;

import com.example.kafka.order.orders.OrderStatus;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.UUID;

@Service
@Slf4j
@RequiredArgsConstructor
public class OrderHistoryServiceImpl implements OrderHistoryService {

    private final OrderHistoryRepository orderHistoryRepository;

    @Override
    public List<OrderHistory> findAll() {
        return orderHistoryRepository.findAll();
    }

    @Override
    public List<OrderHistory> findAllByOrderId(String orderId) {
        return orderHistoryRepository.findAllByOrderId(orderId);
    }

    @Override
    public void add(String orderId, OrderStatus status) {
        OrderHistory history = new OrderHistory();
        history.setId(UUID.randomUUID().toString());
        history.setOrderId(orderId);
        history.setStatus(status);
        orderHistoryRepository.save(history);
    }
}
