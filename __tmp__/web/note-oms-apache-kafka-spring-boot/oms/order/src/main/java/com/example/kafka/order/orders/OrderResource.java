package com.example.kafka.order.orders;

import com.example.kafka.order.histories.OrderHistory;
import com.example.kafka.order.histories.OrderHistoryService;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/v1/orders")
@RequiredArgsConstructor
public class OrderResource {

    private final OrderService orderService;
    private final OrderHistoryService orderHistoryService;

    @PostMapping
    @ResponseStatus(HttpStatus.ACCEPTED)
    public OrderResponse placeOrder(@Valid @RequestBody OrderRequest dto) {
        return orderService.placeOrder(dto);
    }

    @GetMapping("/{orderId}/history")
    @ResponseStatus(HttpStatus.OK)
    public List<OrderHistory> findAllByOrderId(@PathVariable String orderId) {
        return orderHistoryService.findAllByOrderId(orderId);
    }

    @GetMapping("/histories")
    @ResponseStatus(HttpStatus.OK)
    public List<OrderHistory> findAllOrderHistories() {
        return orderHistoryService.findAll();
    }
}
