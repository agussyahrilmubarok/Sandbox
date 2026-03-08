package com.example.booking.service.v1;

import com.example.booking.client.CourseClient;
import com.example.booking.domain.Booking;
import com.example.booking.model.BookingDTO;
import com.example.booking.model.CourseDTO;
import com.example.booking.model.MemberDTO;
import com.example.booking.repos.BookingRepository;
import com.example.booking.service.BookingService;
import com.example.booking.util.BadRequestException;
import com.example.booking.util.NotFoundException;
import feign.FeignException;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestClientException;
import org.springframework.web.client.RestTemplate;

import java.time.LocalDateTime;

@Service("BookingServiceImplV1")
@Slf4j
@RequiredArgsConstructor
public class BookingServiceImpl implements BookingService {

    private final BookingRepository bookingRepository;
    private final RestTemplate restTemplate;
    private final CourseClient courseClient;

    @Override
    public BookingDTO bookCourse(BookingDTO.Request request) {
        MemberDTO memberDTO = null;
        CourseDTO courseDTO = null;

        try {
            memberDTO = getMemberByCode(request.getMemberCode());
            if (memberDTO == null) {
                log.warn("Member not found - code={}", request.getMemberCode());
                throw new NotFoundException("Member not found");
            }

            courseDTO = getCourseByCode(request.getCourseCode());
            if (courseDTO.getSeatAvailable() <= 0) {
                log.warn("No available seats - code={}", request.getCourseCode());
                throw new BadRequestException("No available seats");
            }

            courseClient.reserveCourse(CourseDTO.CodeRequest.builder().code(request.getCourseCode()).build());

            Booking booking = new Booking();
            booking.setMemberCode(memberDTO.getCode());
            booking.setCourseCode(courseDTO.getCode());
            booking.setBookedAt(LocalDateTime.now());
            booking.setStatus(Booking.Status.PENDING);
            booking.setNotes(String.format("Booking course %s by member %s", courseDTO.getCode(), memberDTO.getCode()));

            booking = bookingRepository.save(booking);

            log.info("Booking successfully created - member={}, course={}, bookingId={}", memberDTO.getCode(), courseDTO.getCode(), booking.getId());
            return mapToDTO(booking, new BookingDTO());
        } catch (Exception ex) {
            log.error("Booking failed for member={} course={}", request.getMemberCode(), request.getCourseCode(), ex);

            if (courseDTO != null) {
                try {
                    courseClient.releaseCourse(CourseDTO.CodeRequest.builder().code(request.getCourseCode()).build());
                    log.info("Course released after booking failure - code={}", request.getCourseCode());
                } catch (Exception releaseEx) {
                    log.error("Failed to release course after booking failure - code={}", request.getCourseCode(), releaseEx);
                }
            }

            if (ex instanceof NotFoundException || ex instanceof BadRequestException) {
                throw ex;
            } else {
                throw new BadRequestException("Failed to create booking: " + ex.getMessage(), ex);
            }
        }
    }

    private MemberDTO getMemberByCode(String code) {
        try {
            String url = "http://member-app/api/v1/members/find?code={code}";
            return restTemplate.getForObject(url, MemberDTO.class, code);
        } catch (RestClientException ex) {
            log.error("Failed to fetch member - code={}", code, ex);
            return null;
        }
    }

    private CourseDTO getCourseByCode(String code) {
        try {
            return courseClient.getCourseByCode(code);
        } catch (FeignException.NotFound ex) {
            log.warn("Course not found - code={}", code);
            throw new NotFoundException("Course not found");
        } catch (FeignException ex) {
            log.error("Failed to fetch course - code={}", code, ex);
            throw new BadRequestException("Failed to fetch course", ex);
        }
    }

    private BookingDTO mapToDTO(final Booking booking, final BookingDTO bookingDTO) {
        bookingDTO.setId(booking.getId());
        bookingDTO.setMemberCode(booking.getMemberCode());
        bookingDTO.setCourseCode(booking.getCourseCode());
        bookingDTO.setBookedAt(booking.getBookedAt());
        bookingDTO.setStatus(booking.getStatus().name());
        bookingDTO.setNotes(booking.getNotes());
        bookingDTO.setCreatedAt(booking.getCreatedAt());
        bookingDTO.setUpdatedAt(booking.getUpdatedAt());
        return bookingDTO;
    }
}
