package com.example.kafka.order.orders;

import jakarta.validation.constraints.NotNull;
import jakarta.validation.constraints.Size;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Builder
@Data
@NoArgsConstructor
@AllArgsConstructor
public class OrderRequest {

    @NotNull
    @Size(max = 255)
    private String productId;

    @NotNull
    private Integer total;

}
