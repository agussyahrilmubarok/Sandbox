package com.example.flighthub.auth.service.impl;

import com.example.flighthub.auth.model.dto.request.TokenInvalidateRequest;
import com.example.flighthub.auth.model.entity.UserEntity;
import com.example.flighthub.auth.service.InvalidTokenService;
import com.example.flighthub.auth.service.TokenService;
import com.example.flighthub.base.AbstractBaseServiceTest;
import com.example.flighthub.builder.TokenBuilder;
import com.example.flighthub.builder.UserBuilder;
import io.jsonwebtoken.Claims;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;

import java.util.Set;

import static org.mockito.Mockito.*;

/**
 * Unit test class for {@link LogoutServiceImpl}.
 * This test class validates the functionality of the logout service, which is responsible
 * for handling the user logout process, including token invalidation.
 * It mocks the {@link TokenService} and {@link InvalidTokenService} to verify the logout logic
 * without actually interacting with the token management system.
 */
class LogoutServiceImplTest extends AbstractBaseServiceTest {

    @InjectMocks
    private LogoutServiceImpl logoutService;

    @Mock
    private TokenService tokenService;

    @Mock
    private InvalidTokenService invalidTokenService;

    @Test
    void givenAccessTokenAndRefreshToken_whenLogoutForAdmin_thenReturnLogout() {

        // Given
        final UserEntity mockUserEntity = new UserBuilder().withValidFields().build();

        final Claims mockAccessTokenClaims = TokenBuilder.getValidClaims(
                mockUserEntity.getId(),
                mockUserEntity.getFirstName()
        );

        final Claims mockRefreshTokenClaims = TokenBuilder.getValidClaims(
                mockUserEntity.getId(),
                mockUserEntity.getFirstName()
        );

        final String mockAccessTokenId = mockAccessTokenClaims.getId();
        final String mockRefreshTokenId = mockRefreshTokenClaims.getId();


        final String accessToken = "validAccessToken";
        final String refreshToken = "validRefreshToken";

        final TokenInvalidateRequest tokenInvalidateRequest = TokenInvalidateRequest.builder()
                .accessToken(accessToken)
                .refreshToken(refreshToken)
                .build();

        // When
        doNothing().when(tokenService).verifyAndValidate(Set.of(accessToken, refreshToken));
        when(tokenService.getPayload(accessToken)).thenReturn(mockAccessTokenClaims);
        doNothing().when(invalidTokenService).checkForInvalidityOfToken(mockAccessTokenId);
        when(tokenService.getPayload(refreshToken)).thenReturn(mockRefreshTokenClaims);
        doNothing().when(invalidTokenService).checkForInvalidityOfToken(mockRefreshTokenId);
        doNothing().when(invalidTokenService).invalidateTokens(Set.of(mockAccessTokenId, mockRefreshTokenId));

        // Then
        logoutService.logout(tokenInvalidateRequest);

        // Verify
        verify(tokenService).verifyAndValidate(Set.of(accessToken, refreshToken));
        verify(tokenService, times(2)).getPayload(anyString());
        verify(invalidTokenService, times(2)).checkForInvalidityOfToken(anyString());
        verify(invalidTokenService).invalidateTokens(anySet());

    }

}