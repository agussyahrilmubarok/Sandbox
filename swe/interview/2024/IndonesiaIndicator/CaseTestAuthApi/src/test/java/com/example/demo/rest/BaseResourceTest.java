package com.example.demo.rest;

import com.example.demo.TestcontainersConfig;
import com.example.demo.entity.User;
import com.example.demo.security.UserDetailsCustom;
import com.example.demo.security.UserDetailsServiceImpl;
import com.example.demo.security.jwt.JwtProvider;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.BeforeEach;
import org.mockito.Mockito;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.web.servlet.MockMvc;

@SpringBootTest
@AutoConfigureMockMvc
abstract class BaseResourceTest extends TestcontainersConfig {

    @Autowired
    protected MockMvc mockMvc;

    @Autowired
    protected ObjectMapper objectMapper;

    @Autowired
    protected JwtProvider jwtProvider;

    @MockBean
    protected UserDetailsServiceImpl userDetailsService;

    protected User mockUser;

    protected String mockJwtToken;

    @BeforeEach
    protected void initializedAuth() {
        this.mockUser = makeUser();
        final UserDetailsCustom mockUserDetails = new UserDetailsCustom(mockUser);
        this.mockJwtToken = makeJwtToken(mockUserDetails);

        Mockito.when(userDetailsService.loadUserByUsername(mockUser.getUsername()))
                .thenReturn(mockUserDetails);
    }

    private User makeUser() {
        final User user = new User();
        user.setId(1L);
        user.setUsername("johndoe");
        user.setEmail("johndoe@test.com");
        user.setPassword("Password123");
        return user;
    }

    private String makeJwtToken(UserDetailsCustom userDetails) {
        return "Bearer ".concat(jwtProvider.generateToken(userDetails));
    }
}
