package com.example.kafka.payment.demo;

import com.example.kafka.core.commands.PaymentProcessCommand;
import com.example.kafka.payment.payments.Payment;
import com.example.kafka.payment.payments.PaymentRepository;
import com.example.kafka.payment.payments.PaymentStatus;
import lombok.RequiredArgsConstructor;
import org.springframework.kafka.annotation.KafkaHandler;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.stereotype.Component;

import java.math.BigDecimal;
import java.util.UUID;

@Component
@KafkaListener(topics = "payments-commands")
@RequiredArgsConstructor
public class PaymentCommandHandler {

    private final PaymentRepository paymentRepository;

    @KafkaHandler
    private void productReverseCommand(@Payload PaymentProcessCommand command) {
        Payment payment = new Payment();
        payment.setId(UUID.randomUUID().toString());
        payment.setOrderId(command.getOrderId());
        payment.setProductId(command.getProductId());
        payment.setQuantity(command.getQuantity());
        payment.setPrice(command.getPrice());
        payment.setAmount(BigDecimal.ZERO);
        payment.setStatus(PaymentStatus.PENDING);
        paymentRepository.save(payment);
    }
}
