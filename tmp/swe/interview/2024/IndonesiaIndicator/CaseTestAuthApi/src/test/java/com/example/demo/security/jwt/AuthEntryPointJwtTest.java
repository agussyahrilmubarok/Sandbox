package com.example.demo.security.jwt;

import com.fasterxml.jackson.databind.ObjectMapper;
import jakarta.servlet.ServletException;
import jakarta.servlet.http.HttpServletResponse;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.Mockito;
import org.mockito.junit.jupiter.MockitoExtension;
import org.mockito.junit.jupiter.MockitoSettings;
import org.mockito.quality.Strictness;
import org.springframework.http.MediaType;
import org.springframework.mock.web.MockHttpServletRequest;
import org.springframework.mock.web.MockHttpServletResponse;
import org.springframework.security.core.AuthenticationException;

import java.io.IOException;
import java.util.Map;

@ExtendWith(MockitoExtension.class)
@MockitoSettings(strictness = Strictness.LENIENT)
class AuthEntryPointJwtTest {

    @Mock
    private AuthenticationException authenticationException;

    @InjectMocks
    private AuthEntryPointJwt authEntryPointJwt;

    @Test
    public void testCommence() throws IOException, ServletException {
        MockHttpServletRequest request = new MockHttpServletRequest();
        MockHttpServletResponse response = new MockHttpServletResponse();

        String errorMessage = "Unauthorized access";
        Mockito.when(authenticationException.getMessage()).thenReturn(errorMessage);

        authEntryPointJwt.commence(request, response, authenticationException);

        Assertions.assertEquals(response.getStatus(), HttpServletResponse.SC_UNAUTHORIZED);
        Assertions.assertEquals(response.getContentType(), MediaType.APPLICATION_JSON_VALUE);

        ObjectMapper mapper = new ObjectMapper();
        Map<String, Object> responseBody = mapper.readValue(response.getContentAsString(), Map.class);

        Assertions.assertEquals(responseBody.get("status"), HttpServletResponse.SC_UNAUTHORIZED);
        Assertions.assertEquals(responseBody.get("error"), "Unauthorized");
        Assertions.assertEquals(responseBody.get("message"), errorMessage);
        Assertions.assertEquals(responseBody.get("path"), request.getServletPath());
    }
}