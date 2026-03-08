package com.example.demo.rest;

import com.example.demo.common.Constants;
import org.junit.jupiter.api.Test;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.result.MockMvcResultMatchers;

class InternalResourceTest extends BaseResourceTest {

    @Test
    void givenAccessInternal_whenHaveToken_returnSuccess() throws Exception {
        mockMvc.perform(MockMvcRequestBuilders.get("/api/v1/internal")
                        .contentType(MediaType.APPLICATION_JSON)
                        .header(HttpHeaders.AUTHORIZATION, mockJwtToken))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.content().contentType(MediaType.APPLICATION_JSON))
                .andExpect(MockMvcResultMatchers.jsonPath("$.message")
                        .value(Constants.GET_INTERNAL_MESSAGE));

    }
}