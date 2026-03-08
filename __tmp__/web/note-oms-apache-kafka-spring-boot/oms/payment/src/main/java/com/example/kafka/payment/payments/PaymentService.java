package com.example.kafka.payment.payments;

import java.math.BigDecimal;
import java.util.List;

public interface PaymentService {

    List<PaymentResponse> findAll();

    PaymentResponse setPayOrder(String orderId, BigDecimal amount, boolean schema);

}
