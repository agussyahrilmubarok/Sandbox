package com.example.demo.security.jwt;

import com.example.demo.exception.JwtTokenException;
import io.jsonwebtoken.io.Decoders;
import io.jsonwebtoken.security.Keys;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.junit.jupiter.MockitoExtension;
import org.mockito.junit.jupiter.MockitoSettings;
import org.mockito.quality.Strictness;

import java.security.Key;
import java.security.SecureRandom;

@ExtendWith(MockitoExtension.class)
@MockitoSettings(strictness = Strictness.LENIENT)
class JwtProviderTest {

    @InjectMocks
    private JwtProvider jwtProvider;

    @Test
    public void testValidateToken_WrongToken() {
        JwtTokenException exception = Assertions.assertThrows(JwtTokenException.class, () -> jwtProvider.validateToken("token"));
        Assertions.assertEquals("Invalid Token", exception.getMessage());
    }

    private Key getSignInKey() {
        byte[] keyBytes = Decoders.BASE64.decode(generateRandomHex(64));
        return Keys.hmacShaKeyFor(keyBytes);
    }

    private String generateRandomHex(int length) {
        if (length <= 0 || length % 2 != 0) {
            throw new IllegalArgumentException("Length must be a positive even number");
        }

        byte[] randomBytes = new byte[length / 2];
        SecureRandom secureRandom = new SecureRandom();
        secureRandom.nextBytes(randomBytes);

        StringBuilder hexString = new StringBuilder();
        for (byte b : randomBytes) {
            hexString.append(String.format("%02X", b));
        }

        return hexString.toString();
    }
}