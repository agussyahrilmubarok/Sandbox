package com.example.course.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.validation.constraints.*;
import lombok.*;

import java.math.BigDecimal;
import java.time.LocalDate;
import java.time.OffsetDateTime;
import java.util.UUID;

@Getter
@Setter
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class CourseDTO {

    private UUID id;

    @NotNull
    @Size(max = 20)
    private String code;

    @NotNull
    @Size(max = 255)
    private String name;

    @NotNull
    @DecimalMin(value = "0.0", inclusive = false)
    private BigDecimal price;

    @NotNull
    @JsonProperty("start_date")
    private LocalDate startDate;

    @NotNull
    @JsonProperty("end_date")
    private LocalDate endDate;

    @PositiveOrZero
    private int seat;

    @JsonProperty("seat_available")
    private int seatAvailable;

    @JsonProperty("created_at")
    private OffsetDateTime createdAt;

    @JsonProperty("updated_at")
    private OffsetDateTime updatedAt;

    @Getter
    @Builder
    @NoArgsConstructor(access = AccessLevel.PRIVATE)
    @AllArgsConstructor(access = AccessLevel.PRIVATE)
    public static class CodeRequest {

        @NotBlank(message = "Code is required.")
        private String code;
    }
}
