package utils

import (
	"fmt"
	"strings"
	"time"

	"errors"
	"go-api-starter/core/config"
	"go-api-starter/core/constants"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TokenClaims struct {
	UserID   uuid.UUID `json:"user_id"`
	Email    string    `json:"email"`
	UserName string    `json:"username"`
	Role     string    `json:"roles"`
	Scope    string    `json:"scope"` // access, refresh, reset_password, email_verification

	jwt.RegisteredClaims
}

func ValidateAndParseToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.Get().JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		if claims.ExpiresAt.Unix() < time.Now().Unix() {
			return nil, errors.New("token expired")
		}
		return claims, nil
	}

	return nil, errors.New("invalid token claims")
}

// ValidateJWTToken validates JWT token including expiration check
func ValidateJWTToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.Get().JWT.Secret), nil
	})

	if err != nil {
		return err
	}

	// Check if token is valid and not expired
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		if claims.ExpiresAt.Unix() < time.Now().Unix() {
			return errors.New("token expired")
		}
		return nil
	}

	return errors.New("invalid token claims")
}

func GetTokenFromHeader(c echo.Context) (string, error) {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return "", errors.New("missing Authorization header")
	}

	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		return "", errors.New("invalid token format")
	}

	if err := ValidateJWTToken(token); err != nil {
		return "", errors.New("invalid token: " + err.Error())
	}

	return token, nil
}

func GenerateToken(userID uuid.UUID, email, userName *string, scope string, expireTime ...time.Duration) (string, error) {
	cfg := config.Get()

	// Handle nil email, userID, and userName

	emailValue := ""
	if email != nil {
		emailValue = *email
	}

	userNameValue := ""
	if userName != nil {
		userNameValue = *userName
	}

	// Default scope if not provided
	if scope == "" {
		scope = constants.ScopeTokenAccess
	}

	// Set expiration based on scope
	var expiration time.Duration
	if len(expireTime) > 0 {
		expiration = expireTime[0]
	} else {
		switch scope {
		case constants.ScopeTokenRefresh:
			expiration = constants.DefaultRefreshTokenExpiry // 7 days
		case constants.ScopeTokenResetPassword, constants.ScopeTokenEmailVerification:
			expiration = constants.DefaultResetPasswordTokenExpiry // 5 minutes
		default: // access
			expiration = constants.DefaultAccessTokenExpiry // 24 hours
		}
	}

	claims := TokenClaims{
		UserID:   userID,
		Email:    emailValue,
		UserName: userNameValue,
		Scope:    scope,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.JWT.Secret))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}

func ParseDataFromToken(c echo.Context) *TokenClaims {
	tokenString, err := GetTokenFromHeader(c)
	if err != nil {
		return nil
	}

	claims, err := ValidateAndParseToken(tokenString)
	if err != nil {
		return nil
	}

	return claims
}

// ValidateTokenScope validates if the token has the required scope
func ValidateTokenScope(claims *TokenClaims, requiredScope string) bool {
	return claims != nil && claims.Scope == requiredScope
}

// GenerateAccessToken generates an access token
func GenerateAccessToken(userID uuid.UUID, email, userName *string) (string, error) {
	return GenerateToken(userID, email, userName, constants.ScopeTokenAccess)
}

// GenerateRefreshToken generates a refresh token
func GenerateRefreshToken(userID uuid.UUID, email, userName *string) (string, error) {
	return GenerateToken(userID, email, userName, constants.ScopeTokenRefresh)
}

// GenerateResetPasswordToken generates a reset password token
func GenerateResetPasswordToken(userID uuid.UUID, email *string) (string, error) {
	return GenerateToken(userID, email, nil, constants.ScopeTokenResetPassword)
}

// GenerateEmailVerificationToken generates an email verification token
func GenerateEmailVerificationToken(userID uuid.UUID, email *string) (string, error) {
	return GenerateToken(userID, email, nil, constants.ScopeTokenEmailVerification)
}
