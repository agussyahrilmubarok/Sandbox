package com.example.demo.security;

import com.example.demo.entity.User;
import com.example.demo.exception.UserNotFoundException;
import com.example.demo.repository.UserRepository;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.Mockito;
import org.mockito.junit.jupiter.MockitoExtension;
import org.mockito.junit.jupiter.MockitoSettings;
import org.mockito.quality.Strictness;
import org.springframework.security.core.userdetails.UserDetails;

import java.util.Optional;

@ExtendWith(MockitoExtension.class)
@MockitoSettings(strictness = Strictness.LENIENT)
public class UserDetailsServiceImplTest {

    @Mock
    private UserRepository userRepository;

    @InjectMocks
    private UserDetailsServiceImpl userDetailsService;

    @Test
    void givenExistingEmail_whenLoadUserByUsername_shouldReturnUserDetails() {
        User mockUser = new User();
        mockUser.setId(1L);
        mockUser.setUsername("johndoe");
        mockUser.setPassword("password");

        Mockito.when(userRepository.findByUsername("johndoe"))
                .thenReturn(Optional.of(mockUser));

        UserDetails result = userDetailsService.loadUserByUsername("johndoe");

        Assertions.assertNotNull(result);
        Assertions.assertEquals("johndoe", result.getUsername());

        Mockito.verify(userRepository, Mockito.times(1)).findByUsername("johndoe");
    }

    @Test
    void givenNonExistingEmail_whenLoadUserByUsername_shouldThrowUsernameNotFoundException() {
        Mockito.when(userRepository.findByUsername("nonexistent"))
                .thenReturn(Optional.empty());

        Assertions.assertThrows(UserNotFoundException.class, () -> {
            userDetailsService.loadUserByUsername("nonexistent");
        });

        Mockito.verify(userRepository, Mockito.times(1)).findByUsername("nonexistent");
    }
}