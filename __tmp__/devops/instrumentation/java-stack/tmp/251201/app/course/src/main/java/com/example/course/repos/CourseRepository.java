package com.example.course.repos;

import com.example.course.domain.Course;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;
import java.util.UUID;

public interface CourseRepository extends JpaRepository<Course, UUID> {

    Optional<Course> findByCode(String code);

    boolean existsByCodeIgnoreCase(String code);
}
