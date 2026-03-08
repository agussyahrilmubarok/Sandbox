package com.example.kafka.payment.payments;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.math.BigDecimal;
import java.time.OffsetDateTime;

@Builder
@Data
@NoArgsConstructor
@AllArgsConstructor
public class PaymentResponse {

    private String id;

    private String orderId;

    private String productId;

    private Integer quantity;

    private BigDecimal price;

    private BigDecimal amount;

    private String status;

    private OffsetDateTime dateCreated;

    private OffsetDateTime lastUpdated;
}
