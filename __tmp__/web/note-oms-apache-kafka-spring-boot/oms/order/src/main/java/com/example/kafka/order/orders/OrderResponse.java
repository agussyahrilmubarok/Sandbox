package com.example.kafka.order.orders;

import jakarta.validation.constraints.NotNull;
import jakarta.validation.constraints.Size;
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
public class OrderResponse {

    private String id;

    @NotNull
    @Size(max = 255)
    private String productId;

    @NotNull
    private Integer total;

    private BigDecimal price;

    private String status;

    private OffsetDateTime dateCreated;

    private OffsetDateTime lastUpdated;

}
