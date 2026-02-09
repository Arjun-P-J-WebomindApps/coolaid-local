package crypto

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/auth"
	"github.com/webomindapps-dev/coolaid-backend/pkg/crypto"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) HashPassword(p string) (string, error) {
	return crypto.HashPassword(p) // your existing util
}

func (s *Service) ComparePassword(plain, hash string) error {
	return crypto.CompareHash(plain, hash)
}

func (s *Service) GenerateNumericOTP(n int) (string, error) {
	return crypto.GenerateNumericOTPString(n)
}

func (s *Service) HashToken(t string) (string, error) {
	return crypto.HashToken(t)
}

func (s *Service) ValidateJWT(token string) (*auth.JWTClaims, error) {
	jwtClaims, err := crypto.ValidateJWT(token)

	if err != nil {
		return nil, err
	}

	return &auth.JWTClaims{
		SessionID: jwtClaims.SessionID,
		UserID:    auth.ID(jwtClaims.UserID),
		ExpiresAt: jwtClaims.ExpiresAt.Time,
	}, nil
}

func (s *Service) CompareToken(hash, plain string) error {
	return crypto.CompareToken(hash, plain)
}

func (s *Service) GenerateSecureToken(n int) (string, error) {
	return crypto.GenerateSecureToken(n)
}

func (s *Service) GenerateJWT(ctx context.Context, sessionID, userID auth.ID) (string, error) {
	sID, errSID := sessionID.ToUUID(ctx)
	if errSID != nil {
		return "", errSID
	}

	uID, errUID := userID.ToUUID(ctx)
	if errUID != nil {
		return "", errUID
	}

	return crypto.GenerateJWT(sID, uID)
}

func (s *Service) ValidateJWTFromContext(ctx context.Context) (*auth.JWTClaims, error) {
	jwtClaims, err := crypto.ValidateJWTFromContext(ctx)

	if err != nil {
		return nil, err
	}

	return &auth.JWTClaims{
		SessionID: jwtClaims.SessionID,
		UserID:    auth.ID(jwtClaims.UserID),
		ExpiresAt: jwtClaims.ExpiresAt.Time,
	}, nil
}
