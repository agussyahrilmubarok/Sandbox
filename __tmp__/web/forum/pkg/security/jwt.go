package security

import (
	"errors"
	"time"

	"forum.id/pkg/config"
	"github.com/golang-jwt/jwt"
)

func getJWTConfig() (string, time.Duration) {
	var cfg = config.LoadViper()
	secretKey := cfg.GetString("JWT.SECRET")
	if secretKey == "" {
		secretKey = "default_secret_key"
	}
	expiration := cfg.GetDuration("JWT.EXPIRATION")
	if expiration == 0 {
		expiration = time.Hour * 24 // Default expiration is 24 hours
	}
	return secretKey, expiration
}

// GenerateJWT creates a JWT token for the user
func GenerateJWT(ID int64) string {
	secretKey, expiration := getJWTConfig()

	claims := jwt.MapClaims{
		"id":  ID,
		"exp": time.Now().Add(expiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return ""
	}

	return tokenString
}

// ValidateToken checks if the provided JWT token is valid
func ValidateToken(tokenString string) error {
	secretKey, _ := getJWTConfig()

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return err
	}

	// Check if the token is valid
	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

// ValidateAndParseToken validates and parses the provided JWT token
func ValidateAndParseToken(tokenString string) (*jwt.Token, error) {
	secretKey, _ := getJWTConfig()

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
