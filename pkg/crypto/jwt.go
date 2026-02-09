package crypto

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/config"
)

// JWTClaims represents the custom JWT payload used by the application.
//
// It embeds jwt.RegisteredClaims to include standard fields such as:
// - exp (expiry)
// - iat (issued at)
// - nbf (not before)
// - iss (issuer)
// - sub (subject)
//
// In addition, it carries:
// - UserID: the authenticated user's ID
// - SessionID: the active session ID associated with this token
type JWTClaims struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
	jwt.RegisteredClaims
}

// GenerateJWT creates and signs a JWT access token using HS256.
//
// The generated token includes:
// - UserID and SessionID as custom claims
// - Standard registered claims (expiry, issuer, etc.)
//
// Token expiry duration is controlled via configuration
// (config.Auth.AccessTokenExpiryMinutes).
//
// This function ONLY creates and signs the token.
// It does NOT set cookies or persist anything.
func GenerateJWT(sessionID uuid.UUID, userID uuid.UUID) (string, error) {
	claims := JWTClaims{
		UserID:    userID.String(),
		SessionID: sessionID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(time.Duration(config.Auth.AccessTokenExpiryMinutes) * time.Minute),
			),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "coolaid",
			Subject:   userID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the configured secret key.
	return token.SignedString([]byte(config.Auth.JwtSecret))
}

// ValidateJWT parses and validates a JWT access token string.
//
// It performs the following checks:
// - Verifies the signing algorithm is HMAC (HS256)
// - Validates the token signature using the configured secret
// - Ensures the token is not expired and is otherwise valid
//
// Returns parsed JWTClaims if the token is valid.
// Returns an error if the token is invalid, expired, or malformed.
func ValidateJWT(tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {

		// Ensure the token was signed using HMAC (HS256).
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.Auth.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// ValidateJWTFromContext extracts and validates the JWT access token
// from the request context.
//
// Expected behavior:
// - Retrieves the underlying Gin context from the standard context
// - Reads the "access_token" cookie
// - Validates the token using ValidateJWT
//
// Returns parsed JWTClaims if the token is valid.
// Returns an error if the token is missing, invalid, or context is unavailable.
//
// This function is HTTP-framework aware and should be used
// only at transport / adapter boundaries.
func ValidateJWTFromContext(ctx context.Context) (*JWTClaims, error) {
	ginCtx, ok := config.GinContextFromCtx(ctx)
	if !ok {
		return nil, fmt.Errorf("no gin context")
	}

	tokenStr, err := ginCtx.Cookie("access_token")
	if err != nil || tokenStr == "" {
		return nil, fmt.Errorf("no access token")
	}

	return ValidateJWT(tokenStr)
}
