package com.example.kafka.payment.payments;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/v1/payments")
@RequiredArgsConstructor
public class PaymentResource {

    private final PaymentService paymentService;

    @GetMapping
    public List<PaymentResponse> getAllPayments() {
        return paymentService.findAll();
    }

    @PostMapping("/schema/failed")
    public PaymentResponse schemaFailed(@Valid @RequestBody PaymentRequest request) {
        return paymentService.setPayOrder(request.getOrderId(), request.getAmount(), false);
    }

    @PostMapping("/schema/success")
    public PaymentResponse schemaSuccess(@Valid @RequestBody PaymentRequest request) {
        return paymentService.setPayOrder(request.getOrderId(), request.getAmount(), true);

    }
}
