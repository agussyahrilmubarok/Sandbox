package com.example.kafka.payment.payments;

import org.springframework.data.jpa.repository.JpaRepository;


public interface PaymentRepository extends JpaRepository<Payment, String> {

    Payment findByOrderId(String orderId);

    Payment findTopByOrderIdOrderByDateCreatedDesc(String orderId);

}
