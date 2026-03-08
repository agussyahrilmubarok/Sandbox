package com.example.demo.rest.payload;

import lombok.Data;

import java.util.List;

@Data
public class ErrorPayload {
    private String message;

    private List<String> errors;
}
