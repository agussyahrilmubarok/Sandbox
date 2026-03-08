package com.example.demo.rest.payload;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class JWTPayload {
    private String token;
}
