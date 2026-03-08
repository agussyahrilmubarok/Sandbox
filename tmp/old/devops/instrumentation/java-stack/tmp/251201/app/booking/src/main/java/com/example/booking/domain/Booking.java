package com.example.booking.domain;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;
import org.hibernate.annotations.UuidGenerator;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;
import org.springframework.data.jpa.domain.support.AuditingEntityListener;

import java.time.LocalDateTime;
import java.time.OffsetDateTime;
import java.util.UUID;

@Entity
@Table(name = "Bookings")
@EntityListeners(AuditingEntityListener.class)
@Getter
@Setter
public class Booking {

    @Id
    @Column(nullable = false, updatable = false)
    @GeneratedValue
    @UuidGenerator
    private UUID id;

    @Column(nullable = false)
    private String memberCode;

    @Column(nullable = false)
    private String courseCode;

    @Column(nullable = false)
    private LocalDateTime bookedAt;

    @Enumerated(EnumType.STRING)
    private Status status;

    private String notes;

    @CreatedDate
    @Column(nullable = false, updatable = false)
    private OffsetDateTime createdAt;

    @LastModifiedDate
    @Column(nullable = false)
    private OffsetDateTime updatedAt;

    public enum Status {
        PENDING,
        CONFIRMED,
        CANCELLED
    }
}
