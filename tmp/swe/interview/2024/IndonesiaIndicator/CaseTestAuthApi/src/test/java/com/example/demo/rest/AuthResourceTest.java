package com.example.demo.rest;

import com.example.demo.common.Constants;
import com.example.demo.rest.payload.JWTPayload;
import com.example.demo.rest.payload.LoginPayload;
import com.example.demo.rest.payload.MessagePayload;
import com.example.demo.rest.payload.RegisterPayload;
import com.example.demo.service.impl.AuthServiceImpl;
import org.junit.jupiter.api.Test;
import org.mockito.Mockito;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.result.MockMvcResultMatchers;

class AuthResourceTest extends BaseResourceTest {

    @MockBean
    private AuthServiceImpl authService;

    @Test
    void givenRegisterPayload_WhenRegister_ReturnSuccess() throws Exception {
        RegisterPayload payload = new RegisterPayload();
        payload.setUsername("janesmith");
        payload.setEmail("janesmith@test.com");
        payload.setPassword("Password123");

        Mockito.when(authService.register(payload))
                .thenReturn(new MessagePayload(Constants.REGISTER_SUCCESS));

        mockMvc.perform(MockMvcRequestBuilders.post("/api/v1/auth/register")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(payload)))
                .andExpect(MockMvcResultMatchers.status().isCreated())
                .andExpect(MockMvcResultMatchers.content().contentType(MediaType.APPLICATION_JSON))
                .andExpect(MockMvcResultMatchers.jsonPath("$.message")
                        .value(Constants.REGISTER_SUCCESS));

        Mockito.verify(authService).register(payload);
    }

    @Test
    void givenLoginPayload_WhenLogin_ReturnSuccess() throws Exception {
        LoginPayload payload = new LoginPayload();
        payload.setUsername("janesmith");
        payload.setPassword("Password123");

        JWTPayload mockResponse = new JWTPayload("mockedToken");

        Mockito.when(authService.login(payload)).thenReturn(mockResponse);

        mockMvc.perform(MockMvcRequestBuilders.post("/api/v1/auth/login")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(payload)))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.content().contentType(MediaType.APPLICATION_JSON));

        Mockito.verify(authService).login(payload);
    }
}