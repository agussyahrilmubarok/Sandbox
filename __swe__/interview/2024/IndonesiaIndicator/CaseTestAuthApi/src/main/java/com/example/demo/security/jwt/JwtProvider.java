package com.example.demo.security.jwt;

import com.example.demo.exception.JwtTokenException;
import com.example.demo.security.UserDetailsCustom;
import io.jsonwebtoken.*;
import io.jsonwebtoken.io.Decoders;
import io.jsonwebtoken.security.Keys;
import lombok.extern.log4j.Log4j2;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.security.core.Authentication;
import org.springframework.stereotype.Component;

import java.security.Key;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;

@Component
@Log4j2
public class JwtProvider {

    @Value("${jwt.secretKey}")
    private String jwtSecret;

    @Value("${jwt.expireInMs}")
    private int jwtExpirationInMs;

    public String generateToken(Authentication auth) {
        UserDetailsCustom userDetails = (UserDetailsCustom) auth.getPrincipal();
        Map<String, Object> claims = new HashMap<>();
        claims.put("id", userDetails.getId());

        return createToken(claims, userDetails.getUsername());
    }

    public String generateToken(UserDetailsCustom userDetails) {
        Map<String, Object> claims = new HashMap<>();
        claims.put("id", userDetails.getId());

        return createToken(claims, userDetails.getUsername());
    }

    public boolean validateToken(final String token) {
        try {
            Jwts.parserBuilder().setSigningKey(getSignInKey()).build().parseClaimsJws(token);
            return true;
        } catch (MalformedJwtException e) {
            log.error("JwtUtils | validateJwtToken | Invalid JWT token: {}", e.getMessage());
            throw new JwtTokenException("Invalid JWT token");
        } catch (ExpiredJwtException e) {
            log.error("JwtUtils | validateJwtToken | JWT token is expired: {}", e.getMessage());
            throw new JwtTokenException("JWT token is expired");
        } catch (UnsupportedJwtException e) {
            log.error("JwtUtils | validateJwtToken | JWT token is unsupported: {}", e.getMessage());
            throw new JwtTokenException("JWT token is unsupported");
        } catch (IllegalArgumentException e) {
            log.error("JwtUtils | validateJwtToken | JWT claims string is empty: {}", e.getMessage());
            throw new JwtTokenException("Invalid Token");
        }
    }

    public String getUsernameFromToken(final String token) {
        return extractClaims(token).getSubject();

    }

    private String createToken(Map<String, Object> claims, String subject) {
        Date now = new Date();
        Date expirationDate = new Date(now.getTime() + jwtExpirationInMs);

        return Jwts.builder()
                .setClaims(claims)
                .setSubject(subject)
                .setIssuedAt(now)
                .setExpiration(expirationDate)
                .signWith(getSignInKey(), SignatureAlgorithm.HS256)
                .compact();
    }

    private Key getSignInKey() {
        byte[] keyBytes = Decoders.BASE64.decode(jwtSecret);
        return Keys.hmacShaKeyFor(keyBytes);
    }

    private Claims extractClaims(String token) {
        return Jwts.parserBuilder()
                .setSigningKey(getSignInKey())
                .build()
                .parseClaimsJws(token)
                .getBody();
    }
}
