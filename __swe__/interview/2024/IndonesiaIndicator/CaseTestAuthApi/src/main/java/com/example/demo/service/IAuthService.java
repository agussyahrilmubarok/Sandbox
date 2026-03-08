package com.example.demo.service;

import com.example.demo.rest.payload.JWTPayload;
import com.example.demo.rest.payload.LoginPayload;
import com.example.demo.rest.payload.MessagePayload;
import com.example.demo.rest.payload.RegisterPayload;

public interface IAuthService {
    MessagePayload register(RegisterPayload param);

    JWTPayload login(LoginPayload param);
}
