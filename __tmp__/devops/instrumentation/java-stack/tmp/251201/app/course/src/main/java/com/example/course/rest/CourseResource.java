package com.example.course.rest;

import com.example.course.model.CourseDTO;
import com.example.course.service.CourseService;
import io.swagger.v3.oas.annotations.Operation;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.math.BigDecimal;
import java.time.LocalDate;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.UUID;


@RestController
@RequestMapping(value = "/api/v1/courses", produces = MediaType.APPLICATION_JSON_VALUE)
@Slf4j
@RequiredArgsConstructor
public class CourseResource {

    private final CourseService courseService;

    @GetMapping
    @Operation(
            summary = "Get all courses",
            description = "Retrieves all course data",
            tags = {"Courses"}
    )
    public ResponseEntity<List<CourseDTO>> findAll() {
        List<CourseDTO> courses = courseService.findAll();

        log.info("Successfully retrieved all courses, count={}", courses.size());
        return ResponseEntity.ok(courses);
    }

    @GetMapping("/find")
    @Operation(
            summary = "Find course by code",
            description = "Retrieve a single course by their code",
            tags = {"Courses"}
    )
    public ResponseEntity<CourseDTO> find(@RequestParam(value = "code", required = true) String courseCode) {
        CourseDTO course = courseService.findByCode(courseCode);

        log.info("Successfully retrieved course with code={}", courseCode);
        return ResponseEntity.ok(course);
    }

    @PostMapping("/reserve")
    @Operation(
            summary = "Reserve a seat for a course by course code",
            description = "Reserve a seat for a course by specifying its course code",
            tags = {"Courses"}
    )
    public ResponseEntity<Map<String, String>> reserveByCode(@RequestBody @Valid CourseDTO.CodeRequest payload) {
        courseService.reserveByCode(payload.getCode());

        Map<String, String> response = Map.of(
                "message", "Seat reserved successfully"
        );
        log.info("Seat reserved for courseCode={}", payload.getCode());
        return ResponseEntity.ok(response);
    }

    @PostMapping("/release")
    @Operation(
            summary = "Release a seat for a course by course code",
            description = "Release a seat for a course by specifying its course code",
            tags = {"Courses"}
    )
    public ResponseEntity<Map<String, String>> releaseByCode(@RequestBody @Valid CourseDTO.CodeRequest payload) {
        courseService.releaseByCode(payload.getCode());

        Map<String, String> response = Map.of(
                "message", "Seat released successfully"
        );
        log.info("Seat released for courseCode={}", payload.getCode());
        return ResponseEntity.ok(response);
    }

    @PostMapping("/init-dummy")
    @Operation(
            summary = "Initialize dummy course data",
            description = "Create dummy data for testing",
            tags = {"Dummy"}
    )
    public ResponseEntity<List<CourseDTO>> initDummy() {
        List<CourseDTO> dummies = List.of(
                // COURSE-001 — available
                CourseDTO.builder()
                        .id(UUID.randomUUID())
                        .code("COURSE-001")
                        .name("Go Programming Basics")
                        .price(new BigDecimal("199.99"))
                        .startDate(LocalDate.now())
                        .endDate(LocalDate.now().plusDays(7))
                        .seat(10)
                        .seatAvailable(10)
                        .build(),
                // COURSE-002 — seat = 0, not available
                CourseDTO.builder()
                        .id(UUID.randomUUID())
                        .code("COURSE-002")
                        .name("Advanced Golang")
                        .price(new BigDecimal("249.99"))
                        .startDate(LocalDate.now().plusDays(1))
                        .endDate(LocalDate.now().plusDays(8))
                        .seat(0)
                        .seatAvailable(0)
                        .build(),
                // COURSE-003 — started yesterday
                CourseDTO.builder()
                        .id(UUID.randomUUID())
                        .code("COURSE-003")
                        .name("Docker for Beginners")
                        .price(new BigDecimal("149.99"))
                        .startDate(LocalDate.now().minusDays(1))
                        .endDate(LocalDate.now().plusDays(3))
                        .seat(20)
                        .seatAvailable(20)
                        .build(),
                // COURSE-004 — already ended
                CourseDTO.builder()
                        .id(UUID.randomUUID())
                        .code("COURSE-004")
                        .name("Kubernetes Mastery")
                        .price(new BigDecimal("299.99"))
                        .startDate(LocalDate.now().minusDays(7))
                        .endDate(LocalDate.now().minusDays(1))
                        .seat(15)
                        .seatAvailable(15)
                        .build(),
                // COURSE-005 — valid
                CourseDTO.builder()
                        .id(UUID.randomUUID())
                        .code("COURSE-005")
                        .name("Machine Learning Fundamentals")
                        .price(new BigDecimal("399.99"))
                        .startDate(LocalDate.now().minusDays(2))
                        .endDate(LocalDate.now().plusDays(5))
                        .seat(50)
                        .seatAvailable(50)
                        .build()

        );
        for (CourseDTO dummy : dummies) {
            try {
                courseService.create(dummy);
                log.info("Created dummy course: code={}", dummy.getCode());
            } catch (Exception e) {
                log.error("Failed to create dummy course: code={}", dummy.getCode(), e);
            }
        }

        log.info("Dummy courses initialization completed. Total: {}", dummies.size());
        return ResponseEntity.ok(dummies);
    }

    @DeleteMapping("/clean-dummy")
    @Operation(
            summary = "Remove all dummy courses",
            description = "Delete all dummy course data",
            tags = {"Dummy"}
    )
    public ResponseEntity<Map<String, String>> cleanDummy() {
        List<CourseDTO> courses = courseService.findAll();
        for (CourseDTO course : courses) {
            try {
                courseService.deleteById(course.getId());
                log.info("Deleted dummy course: code={}", course.getCode());
            } catch (Exception e) {
                log.error("Failed to delete dummy course: code={}", course.getCode(), e);
            }
        }

        Map<String, String> response = new HashMap<>();
        response.put("message", "Dummy data cleaned successfully");
        log.info("Dummy courses cleanup completed. Total deleted: {}", courses.size());
        return ResponseEntity.ok(response);
    }
}
