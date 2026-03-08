package com.finly.security.service;

import com.finly.security.UserDetailsImpl;
import com.finly.user.User;
import com.finly.user.repository.UserRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;

@Service
@Slf4j
@RequiredArgsConstructor
public class UserDetailsServiceImpl implements UserDetailsService {

    private final UserRepository userRepository;

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        log.info("Attempting to load user by email: {}", username);
        User user = userRepository.findByEmail(username)
                .orElseThrow(() -> {
                    log.warn("User not found for email: {}", username);
                    return new UsernameNotFoundException("Email or password is wrong");
                });
        log.info("User successfully found for email: {}", username);
        return new UserDetailsImpl(user);
    }
}
