package com.example.course.service;

import com.example.course.domain.Course;
import com.example.course.model.CourseDTO;
import com.example.course.repos.CourseRepository;
import com.example.course.util.BadRequestException;
import com.example.course.util.NotFoundException;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.data.domain.Sort;
import org.springframework.stereotype.Service;

import java.time.LocalDate;
import java.util.List;
import java.util.UUID;

@Service
@Slf4j
@RequiredArgsConstructor
public class CourseServiceImpl implements CourseService {

    private final CourseRepository courseRepository;

    @Override
    public List<CourseDTO> findAll() {
        final List<Course> courses = courseRepository.findAll(Sort.by("code"));

        log.info("Successfully fetched {} courses", courses.size());
        return courses.stream()
                .map(course -> mapToDTO(course, new CourseDTO()))
                .toList();
    }

    @Override
    public CourseDTO findById(UUID id) {
        return courseRepository.findById(id)
                .map(course -> {
                    log.info("Course found for id={}", id);
                    return mapToDTO(course, new CourseDTO());
                })
                .orElseThrow(() -> {
                    log.warn("Course not found for id={}", id);
                    return new NotFoundException("Course not found");
                });
    }

    @Override
    public CourseDTO findByCode(String code) {
        return courseRepository.findByCode(code)
                .map(course -> {
                    log.info("Course found for code={}", code);
                    return mapToDTO(course, new CourseDTO());
                })
                .orElseThrow(() -> {
                    log.warn("Course not found for code={}", code);
                    return new NotFoundException("Course not found");
                });
    }

    @Override
    public CourseDTO create(CourseDTO courseDTO) {
        if (codeExists(courseDTO.getCode())) {
            log.warn("Failed creating course — code={} already exists", courseDTO.getCode());
            throw new IllegalArgumentException("Code already exists");
        }

        Course course = new Course();
        mapToEntity(courseDTO, course);
        course = courseRepository.save(course);

        log.info("Successfully created course id={} code={}", course.getId(), course.getCode());
        return mapToDTO(course, new CourseDTO());
    }

    @Override
    public CourseDTO update(UUID id, CourseDTO courseDTO) {
        Course course = courseRepository.findById(id)
                .orElseThrow(() -> {
                    log.warn("Update failed — course not found for id={}", id);
                    return new NotFoundException("Course not found");
                });

        mapToEntity(courseDTO, course);
        course = courseRepository.save(course);

        log.info("Successfully updated course id={}", id);
        return mapToDTO(course, new CourseDTO());
    }

    @Override
    public void deleteById(UUID id) {
        Course course = courseRepository.findById(id)
                .orElseThrow(() -> {
                    log.warn("Delete failed — course not found for id={}", id);
                    return new NotFoundException("Course not found");
                });

        courseRepository.delete(course);
        log.info("Successfully deleted course id={}", id);
    }

    @Override
    public void reserveByCode(String code) {
        Course course = courseRepository.findByCode(code)
                .orElseThrow(() -> {
                    log.warn("Reserve failed — course not found for code={}", code);
                    return new NotFoundException("Course not found");
                });

        LocalDate today = LocalDate.now();
        if (today.isAfter(course.getEndDate())) {
            log.warn("Reserve failed — course has ended - code={}", code);
            throw new BadRequestException("Course has already ended");
        }

        if (course.getSeatAvailable() <= 0) {
            log.warn("Reserve failed — no seats available - code={}", code);
            throw new BadRequestException("No available seats to reserve");
        }

        course.setSeatAvailable(course.getSeatAvailable() - 1);
        courseRepository.save(course);

        log.info("Course reserved successfully - code={}, remainingSeats={}", code, course.getSeatAvailable());
    }

    @Override
    public void releaseByCode(String code) {
        Course course = courseRepository.findByCode(code)
                .orElseThrow(() -> {
                    log.warn("Release failed — course not found, code={}", code);
                    return new NotFoundException("Course not found");
                });

        LocalDate today = LocalDate.now();
        if (today.isAfter(course.getEndDate())) {
            log.warn("Release failed — course has ended, code={}", code);
            throw new BadRequestException("Course has already ended");
        }

        if (course.getSeatAvailable() >= course.getSeat()) {
            log.warn("Release failed — seatAvailable already full, code={}", code);
            throw new BadRequestException("All seats are already available");
        }

        course.setSeatAvailable(course.getSeatAvailable() + 1);
        courseRepository.save(course);

        log.info("Course released successfully — code={}, remainingSeats={}", code, course.getSeatAvailable());
    }

    private CourseDTO mapToDTO(final Course course, final CourseDTO courseDTO) {
        courseDTO.setId(course.getId());
        courseDTO.setCode(course.getCode());
        courseDTO.setName(course.getName());
        courseDTO.setPrice(course.getPrice());
        courseDTO.setStartDate(course.getStartDate());
        courseDTO.setEndDate(course.getEndDate());
        courseDTO.setSeat(course.getSeat());
        courseDTO.setSeatAvailable(course.getSeatAvailable());
        courseDTO.setCreatedAt(course.getCreatedAt());
        courseDTO.setUpdatedAt(course.getUpdatedAt());
        return courseDTO;
    }

    private Course mapToEntity(final CourseDTO courseDTO, final Course course) {
        course.setCode(courseDTO.getCode());
        course.setName(courseDTO.getName());
        course.setPrice(courseDTO.getPrice());
        course.setStartDate(courseDTO.getStartDate());
        course.setEndDate(courseDTO.getEndDate());
        course.setSeat(courseDTO.getSeat());
        course.setSeatAvailable(courseDTO.getSeatAvailable());
        return course;
    }

    private boolean codeExists(final String code) {
        return courseRepository.existsByCodeIgnoreCase(code);
    }
}
