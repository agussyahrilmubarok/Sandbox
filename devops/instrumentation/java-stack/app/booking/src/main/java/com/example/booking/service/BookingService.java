package com.example.booking.service;

import com.example.booking.model.BookingDTO;

public interface BookingService {

    BookingDTO bookCourse(BookingDTO.Request request);
}
