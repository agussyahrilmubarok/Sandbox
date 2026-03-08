package com.example.course.service;

import com.example.course.model.CourseDTO;

import java.util.List;
import java.util.UUID;

public interface CourseService {

    List<CourseDTO> findAll();

    CourseDTO findById(UUID id);

    CourseDTO findByCode(String code);

    CourseDTO create(CourseDTO courseDTO);

    CourseDTO update(UUID id, CourseDTO courseDTO);

    void deleteById(UUID id);

    void reserveByCode(String code);

    void releaseByCode(String code);
}
