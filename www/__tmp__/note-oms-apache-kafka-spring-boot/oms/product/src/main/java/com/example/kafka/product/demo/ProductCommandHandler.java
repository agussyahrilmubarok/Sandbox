package com.example.kafka.product.demo;

import com.example.kafka.core.commands.ProductReserveCommand;
import com.example.kafka.core.commands.ProductRollbackCommand;
import com.example.kafka.core.events.StockNotAvailableEvent;
import com.example.kafka.core.events.StockReservedEvent;
import com.example.kafka.core.events.StockRollbackEvent;
import com.example.kafka.product.products.Product;
import com.example.kafka.product.products.ProductRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.kafka.annotation.KafkaHandler;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.stereotype.Component;

import java.math.BigDecimal;

@Component
@KafkaListener(topics = "products-commands")
@RequiredArgsConstructor
public class ProductCommandHandler {

    private final KafkaTemplate<String, Object> kafkaTemplate;
    private final ProductRepository productRepository;

    @KafkaHandler
    private void productReverseCommand(@Payload ProductReserveCommand command) {
        Product productReserved = reserveStock(command.getProductId(), command.getQuantity());
        if (productReserved == null) {
            StockNotAvailableEvent event = new StockNotAvailableEvent();
            event.setOrderId(command.getOrderId());
            event.setProductId(command.getProductId());
            kafkaTemplate.send("products-events", event);
        } else {
            StockReservedEvent event = new StockReservedEvent();
            event.setOrderId(command.getOrderId());
            event.setProductId(command.getProductId());
            event.setQuantity(command.getQuantity());
            event.setPrice(productReserved.getPrice().multiply(BigDecimal.valueOf(command.getQuantity())));
            kafkaTemplate.send("products-events", event);
        }
    }

    @KafkaHandler
    private void productRollbackCommand(@Payload ProductRollbackCommand command) {
        Product productAdjustment = adjustmentStock(command.getProductId(), command.getQuantity());
        StockRollbackEvent event = new StockRollbackEvent();
        event.setOrderId(command.getOrderId());
        event.setProductId(productAdjustment.getId());
        event.setQuantity(command.getQuantity());
        kafkaTemplate.send("products-events", event);
    }

    private Product reserveStock(String id, Integer quantity) {
        Product product = productRepository.findById(id).orElse(null);
        if (product == null || product.getStock() < quantity) {
            return null;
        } else {
            product.setStock(product.getStock() - quantity);
            return productRepository.save(product);
        }
    }

    private Product adjustmentStock(String id, Integer quantity) {
        Product product = productRepository.findById(id).orElseThrow();
        Integer newStock = product.getStock() + quantity;
        product.setStock(newStock);
        return productRepository.save(product);
    }
}
