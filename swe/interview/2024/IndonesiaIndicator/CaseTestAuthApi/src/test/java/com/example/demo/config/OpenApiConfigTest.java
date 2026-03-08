package com.example.demo.config;

import io.swagger.v3.oas.annotations.OpenAPIDefinition;
import io.swagger.v3.oas.annotations.enums.SecuritySchemeIn;
import io.swagger.v3.oas.annotations.enums.SecuritySchemeType;
import io.swagger.v3.oas.annotations.info.Contact;
import io.swagger.v3.oas.annotations.info.Info;
import io.swagger.v3.oas.annotations.security.SecurityScheme;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;


class OpenApiConfigTest {
    @Test
    void openApiInfo() {
        OpenAPIDefinition openAPIDefinition = OpenApiConfig.class.getAnnotation(OpenAPIDefinition.class);

        Assertions.assertEquals("1.0.0", openAPIDefinition.info().version());
        Assertions.assertEquals("API for user authentication", openAPIDefinition.info().title());
        Assertions.assertEquals("Web service is Restful API for user authentication that allows users to register, log in and \n" +
                "   access protected resources.", openAPIDefinition.info().description());

    }

    @Test
    void securityScheme() {
        SecurityScheme securityScheme = OpenApiConfig.class.getAnnotation(SecurityScheme.class);

        Assertions.assertEquals("bearerAuth", securityScheme.name());
        Assertions.assertEquals("JWT Token", securityScheme.description());
        Assertions.assertEquals("bearer", securityScheme.scheme());
        Assertions.assertEquals(SecuritySchemeType.HTTP, securityScheme.type());
        Assertions.assertEquals("JWT", securityScheme.bearerFormat());
        Assertions.assertEquals(SecuritySchemeIn.HEADER, securityScheme.in());
    }

    @Test
    void contactInfo() {
        Info info = OpenApiConfig.class.getAnnotation(OpenAPIDefinition.class).info();
        Contact contact = info.contact();

        Assertions.assertEquals("Agus Syahril Mubarok", contact.name());
        Assertions.assertEquals("agussmkertjhaan@gmail.com", contact.email());
    }
}