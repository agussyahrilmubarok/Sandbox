package com.example.demo.rest;

import com.example.demo.rest.payload.JWTPayload;
import com.example.demo.rest.payload.LoginPayload;
import com.example.demo.rest.payload.MessagePayload;
import com.example.demo.rest.payload.RegisterPayload;
import com.example.demo.service.IAuthService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/v1/auth")
@RequiredArgsConstructor
@Tag(name = "Authentication", description = "API for authentication operations such as register, and login.")
public class AuthResource {

    private final IAuthService authService;

    @PostMapping("/register")
    @ResponseStatus(HttpStatus.CREATED)
    @Operation(summary = "Register a new user", description = "Registers a new user and returns a success message.",
            responses = {
                    @ApiResponse(responseCode = "201", description = "User registered successfully",
                            content = @Content(schema = @Schema(implementation = ResponseEntity.class)))
            })
    public ResponseEntity<MessagePayload> register(@Valid @RequestBody RegisterPayload payload) {
        return new ResponseEntity<>(authService.register(payload), HttpStatus.CREATED);
    }

    @PostMapping("/login")
    @Operation(summary = "Login user", description = "Logs in a user and returns JWT tokens.",
            responses = {
                    @ApiResponse(responseCode = "200", description = "User logged in successfully",
                            content = @Content(schema = @Schema(implementation = ResponseEntity.class)))
            })
    public ResponseEntity<JWTPayload> login(@Valid @RequestBody LoginPayload payload) {
        return new ResponseEntity<>(authService.login(payload), HttpStatus.OK);
    }
}
