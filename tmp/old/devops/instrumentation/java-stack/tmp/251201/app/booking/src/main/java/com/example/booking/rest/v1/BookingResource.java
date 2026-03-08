package com.example.booking.rest.v1;

import com.example.booking.model.BookingDTO;
import com.example.booking.service.BookingService;
import jakarta.validation.Valid;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController("BookingResourceV1")
@RequestMapping(value = "/api/v1/booking", produces = MediaType.APPLICATION_JSON_VALUE)
@Slf4j
public class BookingResource {

    private final BookingService bookingService;

    public BookingResource(@Qualifier("BookingServiceImplV1") BookingService bookingService) {
        this.bookingService = bookingService;
    }

    @PostMapping("/course")
    public ResponseEntity<BookingDTO> bookCourse(@RequestBody @Valid BookingDTO.Request payload) {
        return ResponseEntity.ok(bookingService.bookCourse(payload));
    }
}
