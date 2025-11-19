package com.example.booking.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.validation.constraints.NotBlank;
import lombok.*;

import java.time.LocalDateTime;
import java.time.OffsetDateTime;
import java.util.UUID;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class BookingDTO {

    private UUID id;

    @JsonProperty("member_code")
    private String memberCode;

    @JsonProperty("course_code")
    private String courseCode;

    @JsonProperty("booked_at")
    private LocalDateTime bookedAt;

    private String status;

    private String notes;

    @JsonProperty("created_at")
    private OffsetDateTime createdAt;

    @JsonProperty("updated_at")
    private OffsetDateTime updatedAt;

    @Getter
    @Builder
    @NoArgsConstructor(access = AccessLevel.PRIVATE)
    @AllArgsConstructor(access = AccessLevel.PRIVATE)
    public static class Request {

        @NotBlank(message = "Member code is required.")
        @JsonProperty("member_code")
        private String memberCode;

        @NotBlank(message = "Course code is required.")
        @JsonProperty("course_code")
        private String courseCode;
    }
}
