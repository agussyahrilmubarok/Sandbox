package com.example.flighthub.common.config;

import io.swagger.v3.oas.annotations.OpenAPIDefinition;
import io.swagger.v3.oas.annotations.enums.SecuritySchemeIn;
import io.swagger.v3.oas.annotations.enums.SecuritySchemeType;
import io.swagger.v3.oas.annotations.info.Contact;
import io.swagger.v3.oas.annotations.info.Info;
import io.swagger.v3.oas.annotations.security.SecurityScheme;

/**
 * Configuration class for setting up OpenAPI documentation for the application.
 * This class configures the metadata and security settings for the API documentation using OpenAPI 3.0 annotations.
 */
@OpenAPIDefinition(
        info = @Info(
                contact = @Contact(
                        name = "MZS Software House",
                        url = "https://github.com/mzs-house/flighthub"
                ),
                description = "Case Study - Flight Hub API" +
                        " (Java 21, Spring Boot, MongoDB, JUnit, Spring Security, JWT, Docker, Kubernetes, Prometheus, Grafana, Github Actions (CI/CD) ) ",
                title = "flighthub",
                version = "1.0.0"
        )
)
@SecurityScheme(
        name = "bearerAuth",
        description = "JWT Token",
        scheme = "bearer",
        type = SecuritySchemeType.HTTP,
        bearerFormat = "JWT",
        in = SecuritySchemeIn.HEADER
)
public class OpenApiConfig {

}
