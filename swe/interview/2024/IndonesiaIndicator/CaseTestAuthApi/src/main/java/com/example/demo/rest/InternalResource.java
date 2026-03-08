package com.example.demo.rest;

import com.example.demo.common.Constants;
import com.example.demo.rest.payload.MessagePayload;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/v1/internal")
@Tag(name = "Internal", description = "API for authentication internal operations such as get resource.")
public class InternalResource {

    @GetMapping
    @Operation(summary = "Internal Resource", description = "User can access if logged in.",
            responses = {
                    @ApiResponse(responseCode = "200", description = "Internal resources accessed successfully",
                            content = @Content(schema = @Schema(implementation = MessagePayload.class)))
            })
    public ResponseEntity<MessagePayload> get() {
        return new ResponseEntity<>(new MessagePayload(Constants.GET_INTERNAL_MESSAGE), HttpStatus.OK);
    }
}
