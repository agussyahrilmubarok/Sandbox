package com.example.demo.service.impl;

import com.example.demo.common.Constants;
import com.example.demo.entity.User;
import com.example.demo.exception.AlreadyExistsException;
import com.example.demo.repository.UserRepository;
import com.example.demo.rest.payload.JWTPayload;
import com.example.demo.rest.payload.LoginPayload;
import com.example.demo.rest.payload.MessagePayload;
import com.example.demo.rest.payload.RegisterPayload;
import com.example.demo.security.jwt.JwtProvider;
import com.example.demo.service.IAuthService;
import lombok.RequiredArgsConstructor;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class AuthServiceImpl implements IAuthService {

    private final UserRepository userRepository;
    private final PasswordEncoder passwordEncoder;
    private final AuthenticationManager authenticationManager;
    private final JwtProvider jwtProvider;

    @Override
    public MessagePayload register(RegisterPayload param) {
        if (userRepository.existsByUsernameIgnoreCase(param.getUsername())) {
            throw new AlreadyExistsException(Constants.USERNAME_ALREADY_EXISTS);
        }

        if (userRepository.existsByEmailIgnoreCase(param.getEmail())) {
            throw new AlreadyExistsException(Constants.EMAIL_ALREADY_EXISTS);
        }

        final User user = new User();
        user.setUsername(param.getUsername());
        user.setEmail(param.getEmail());
        user.setPassword(passwordEncoder.encode(param.getPassword()));
        userRepository.save(user);

        return new MessagePayload(Constants.REGISTER_SUCCESS);
    }

    @Override
    public JWTPayload login(LoginPayload param) {
        UsernamePasswordAuthenticationToken authToken =
                new UsernamePasswordAuthenticationToken(param.getUsername(), param.getPassword());

        Authentication auth = authenticationManager.authenticate(authToken);
        SecurityContextHolder.getContext().setAuthentication(auth);
        String jwtToken = jwtProvider.generateToken(auth);

        return new JWTPayload(jwtToken);
    }
}
