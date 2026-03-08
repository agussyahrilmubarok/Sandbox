package com.example.kafka.payment.demo;

import com.example.kafka.core.events.PaymentFailedEvent;
import com.example.kafka.core.events.PaymentProcessedEvent;
import com.example.kafka.payment.payments.Payment;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Service;

@Service
@Slf4j
@RequiredArgsConstructor
public class PaymentProducer {

    private final KafkaTemplate<String, Object> kafkaTemplate;

    public void paymentProcessedEvent(Payment payment) {
        PaymentProcessedEvent event = new PaymentProcessedEvent();
        event.setOrderId(payment.getOrderId());
        event.setPaymentId(payment.getId());
        kafkaTemplate.send("payments-events", event);
    }

    public void paymentFailedEvent(Payment payment) {
        PaymentFailedEvent event = new PaymentFailedEvent();
        event.setOrderId(payment.getOrderId());
        event.setProductId(payment.getProductId());
        event.setQuantity(payment.getQuantity());
        event.setReason("TIMEOUT PAYMENT");
        kafkaTemplate.send("payments-events", event);
    }
}
