package com.example.booking.client;

import com.example.booking.model.CourseDTO;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestParam;

@FeignClient(name = "course-app")
public interface CourseClient {

    @GetMapping("/api/v1/courses/find")
    CourseDTO getCourseByCode(@RequestParam("code") String code);

    @PostMapping("/api/v1/courses/reserve")
    void reserveCourse(@RequestBody CourseDTO.CodeRequest request);

    @PostMapping("/api/v1/courses/release")
    void releaseCourse(@RequestBody CourseDTO.CodeRequest request);
}
