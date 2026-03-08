package middleware

import (
	"errors"
	"strings"

	"forum.id/pkg/security"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// JWTAuthorization is a middleware that checks if the request has a valid JWT token
func JWTAuthorization(c *fiber.Ctx) error {
	// Get the Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Authorization header is required")
	}

	// Extract the token from the Authorization header
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid authorization header format")
	}

	// Validate and parse the token
	token, err := security.ValidateAndParseToken(tokenString)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
	}

	// Extract the user ID from the token claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := extractUserIDFromClaims(claims)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token: user ID not found")
		}

		// Set userID into the local context for use in handlers
		c.Locals("userID", userID)
	} else {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	// Token is valid, proceed to the next middleware/handler
	return c.Next()
}

// Helper function to extract user ID from JWT claims
func extractUserIDFromClaims(claims jwt.MapClaims) (int64, error) {
	userIDFloat, ok := claims["id"].(float64) // JWT often stores numbers as float64
	if !ok {
		return 0, errors.New("user ID not found in token")
	}

	// Convert float64 to int64
	return int64(userIDFloat), nil
}
