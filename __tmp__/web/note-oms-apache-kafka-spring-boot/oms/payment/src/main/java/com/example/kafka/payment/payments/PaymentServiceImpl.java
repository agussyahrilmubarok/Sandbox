package com.example.kafka.payment.payments;

import com.example.kafka.payment.demo.PaymentProducer;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
import java.util.List;
import java.util.stream.Collectors;

@Service
@Slf4j
@RequiredArgsConstructor
public class PaymentServiceImpl implements PaymentService {

    private final PaymentRepository paymentRepository;
    private final PaymentProducer paymentProducer;

    @Override
    public List<PaymentResponse> findAll() {
        List<Payment> payments = paymentRepository.findAll();

        return payments.stream().map(payment -> PaymentResponse.builder()
                .id(payment.getId())
                .orderId(payment.getOrderId())
                .productId(payment.getProductId())
                .quantity(payment.getQuantity())
                .price(payment.getPrice())
                .amount(payment.getAmount())
                .status(payment.getStatus().name())
                .dateCreated(payment.getDateCreated())
                .lastUpdated(payment.getLastUpdated())
                .build()).collect(Collectors.toList());
    }

    @Override
    public PaymentResponse setPayOrder(String orderId, BigDecimal amount, boolean schema) {
        Payment payment = paymentRepository.findTopByOrderIdOrderByDateCreatedDesc(orderId);
        if (schema) {
            payment.setAmount(amount);
            payment.setStatus(PaymentStatus.SUCCESS);
            Payment savedPayment = paymentRepository.save(payment);
            paymentProducer.paymentProcessedEvent(savedPayment);
            return PaymentResponse.builder()
                    .id(savedPayment.getId())
                    .orderId(savedPayment.getOrderId())
                    .productId(savedPayment.getProductId())
                    .quantity(savedPayment.getQuantity())
                    .price(savedPayment.getPrice())
                    .amount(savedPayment.getAmount())
                    .status(savedPayment.getStatus().name())
                    .dateCreated(savedPayment.getDateCreated())
                    .lastUpdated(savedPayment.getLastUpdated())
                    .build();
        } else {
            payment.setStatus(PaymentStatus.FAILED);
            Payment savedPayment = paymentRepository.save(payment);
            paymentProducer.paymentFailedEvent(savedPayment);
            return PaymentResponse.builder()
                    .id(savedPayment.getId())
                    .orderId(savedPayment.getOrderId())
                    .productId(savedPayment.getProductId())
                    .quantity(savedPayment.getQuantity())
                    .price(savedPayment.getPrice())
                    .amount(savedPayment.getAmount())
                    .status(savedPayment.getStatus().name())
                    .dateCreated(savedPayment.getDateCreated())
                    .lastUpdated(savedPayment.getLastUpdated())
                    .build();
        }
    }


}
