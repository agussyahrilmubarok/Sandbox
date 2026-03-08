package com.example.booking.repos;

import com.example.booking.domain.Booking;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;
import java.util.UUID;

public interface BookingRepository extends JpaRepository<Booking, UUID> {

    List<Booking> findByMemberCode(String memberCode);

    List<Booking> findByCourseCode(String courseCode);

    List<Booking> findByStatus(Booking.Status status);
}
