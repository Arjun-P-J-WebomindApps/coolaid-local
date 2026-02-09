package auth

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
	"github.com/webomindapps-dev/coolaid-backend/config"
)

func (s *Service) VerifyTOTP(
	ctx context.Context,
	username string,
	code string,
	meta SessionMeta,
	writer CookieWriter,
) (*AuthResult, error) {

	tx, qtx, err := s.DB.BeginTx(ctx)
	if err != nil {
		return nil, ErrInternal
	}
	defer tx.Rollback()

	user, err := qtx.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	t, err := qtx.GetUserTotpByUserID(ctx, user.ID)
	if err != nil || !totp.Validate(code, t.Secret) {
		return nil, ErrInvalidOTP
	}

	session, err := qtx.GetUserSessionByUserId(ctx, user.ID)
	if err != nil || session.ExpiresAt.Before(time.Now()) {
		if err == nil {
			_ = qtx.DeleteUserSession(ctx, session.SessionID)
		}

		session, err = qtx.CreateUserSession(ctx, CreateSessionParams{
			UserID:    user.ID,
			SessionID: ID(uuid.New().String()),
			UserAgent: meta.UserAgent,
			ExpiresAt: time.Now().Add(24 * time.Hour),
			IP:        "",
		})
		if err != nil {
			return nil, ErrInternal
		}
	}

	_ = qtx.DeleteRefreshTokensBySession(ctx, session.SessionID)

	raw, _ := s.Crypto.GenerateSecureToken(64)
	hash, _ := s.Crypto.HashToken(raw)

	errRefresh := qtx.CreateRefreshToken(ctx, CreateRefreshTokenParams{
		UserID:    user.ID,
		SessionID: session.SessionID,
		TokenHash: hash,
		ExpiresAt: time.Now().Add(time.Duration(config.Auth.RefreshTokenExpiryHours) * time.Hour),
	})

	if errRefresh != nil {
		return nil, ErrRefreshExpired
	}

	jwt, _ := s.Crypto.GenerateJWT(ctx, ID(session.SessionID), ID(user.ID))

	writer.SetAuth(jwt, raw)

	tx.Commit()

	return &AuthResult{
		UserID:   ID(user.ID),
		Username: user.Username,
	}, nil
}
