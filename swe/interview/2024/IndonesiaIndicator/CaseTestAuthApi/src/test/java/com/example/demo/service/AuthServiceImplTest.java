package com.example.demo.service;

import com.example.demo.common.Constants;
import com.example.demo.entity.User;
import com.example.demo.exception.AlreadyExistsException;
import com.example.demo.repository.UserRepository;
import com.example.demo.rest.payload.JWTPayload;
import com.example.demo.rest.payload.LoginPayload;
import com.example.demo.rest.payload.MessagePayload;
import com.example.demo.rest.payload.RegisterPayload;
import com.example.demo.security.jwt.JwtProvider;
import com.example.demo.service.impl.AuthServiceImpl;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.Mockito;
import org.mockito.junit.jupiter.MockitoExtension;
import org.mockito.junit.jupiter.MockitoSettings;
import org.mockito.quality.Strictness;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.crypto.password.PasswordEncoder;

@ExtendWith(MockitoExtension.class)
@MockitoSettings(strictness = Strictness.LENIENT)
public class AuthServiceImplTest {

    @Mock
    private UserRepository userRepository;

    @Mock
    private PasswordEncoder passwordEncoder;

    @Mock
    private AuthenticationManager authenticationManager;

    @Mock
    private JwtProvider jwtProvider;

    @InjectMocks
    private AuthServiceImpl authService;

    @Test
    void givenRegisterPayload_whenRegister_shouldReturnSuccess() {
        User mockUser = new User();
        mockUser.setUsername("johndoe");
        mockUser.setEmail("johndoe@test.com");
        mockUser.setPassword("PasswordEncoded");
        Mockito.when(userRepository.existsByUsernameIgnoreCase(Mockito.anyString()))
                .thenReturn(false);
        Mockito.when(userRepository.existsByEmailIgnoreCase(Mockito.anyString()))
                .thenReturn(false);
        Mockito.when(userRepository.save(Mockito.any(User.class)))
                .thenReturn(mockUser);

        RegisterPayload payload = new RegisterPayload();
        payload.setUsername("johndoe");
        payload.setEmail("johndoe@test.com");
        payload.setPassword("password");
        MessagePayload result = authService.register(payload);

        Assertions.assertNotNull(result);
        Assertions.assertEquals(Constants.REGISTER_SUCCESS, result.getMessage());

        Mockito.verify(userRepository).existsByUsernameIgnoreCase(Mockito.anyString());
        Mockito.verify(userRepository).existsByEmailIgnoreCase(Mockito.anyString());
        Mockito.verify(userRepository).save(Mockito.any(User.class));
    }

    @Test
    void givenExistingUsername_whenRegister_shouldThrowUsernameAlreadyExistsException() {
        RegisterPayload payload = new RegisterPayload();
        payload.setUsername("johndoe");
        payload.setEmail("johndoe@test.com");
        payload.setPassword("password");

        Mockito.when(userRepository.existsByUsernameIgnoreCase(Mockito.anyString()))
                .thenReturn(true);

        Assertions.assertThrows(AlreadyExistsException.class, () -> {
            authService.register(payload);
        });

        Mockito.verify(userRepository).existsByUsernameIgnoreCase(Mockito.anyString());
        Mockito.verify(userRepository, Mockito.never()).existsByEmailIgnoreCase(Mockito.anyString());
        Mockito.verify(userRepository, Mockito.never()).save(Mockito.any(User.class));
    }

    @Test
    void givenExistingEmail_whenRegister_shouldThrowEmailAlreadyExistsException() {
        RegisterPayload payload = new RegisterPayload();
        payload.setUsername("johndoe");
        payload.setEmail("johndoe@test.com");
        payload.setPassword("password");

        Mockito.when(userRepository.existsByUsernameIgnoreCase(Mockito.anyString()))
                .thenReturn(false);
        Mockito.when(userRepository.existsByEmailIgnoreCase(Mockito.anyString()))
                .thenReturn(true);

        Assertions.assertThrows(AlreadyExistsException.class, () -> {
            authService.register(payload);
        });

        Mockito.verify(userRepository).existsByUsernameIgnoreCase(Mockito.anyString());
        Mockito.verify(userRepository).existsByEmailIgnoreCase(Mockito.anyString());
        Mockito.verify(userRepository, Mockito.never()).save(Mockito.any(User.class));
    }

    @Test
    void givenLoginPayload_whenLogin_shouldReturnJwtToken() {
        LoginPayload payload = new LoginPayload();
        payload.setUsername("johndoe");
        payload.setPassword("password");
        Authentication mockAuthentication = new
                UsernamePasswordAuthenticationToken(payload.getUsername(), payload.getPassword());

        Mockito.when(authenticationManager.authenticate(Mockito.any(UsernamePasswordAuthenticationToken.class)))
                .thenReturn(mockAuthentication);
        Mockito.when(jwtProvider.generateToken(Mockito.any(Authentication.class)))
                .thenReturn("mockedToken");

        JWTPayload result = authService.login(payload);

        Assertions.assertNotNull(result);
        Assertions.assertEquals("mockedToken", result.getToken());

        Mockito.verify(authenticationManager).authenticate(Mockito.any(UsernamePasswordAuthenticationToken.class));
        Mockito.verify(jwtProvider).generateToken(mockAuthentication);
    }

}
