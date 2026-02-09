package auth

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/oplog"
)

type ID string

// ToUUID converts an auth.ID into a uuid.UUID with logging support
func (id ID) ToUUID(ctx context.Context) (uuid.UUID, error) {
	u, err := uuid.Parse(string(id))
	if err != nil {
		oplog.Error(ctx,
			"failed to convert auth.ID to uuid",
			"id=", id,
			"error=", err,
		)

		return uuid.Nil, fmt.Errorf("invalid id %q: %w", id, err)
	}

	return u, nil
}

type Queries interface {
	GetUserByEmail(ctx context.Context, email string) (*UserRow, error)
	GetUserByUsername(ctx context.Context, username string) (*UserRow, error)
	GetUserById(ctx context.Context, id ID) (*UserRow, error)

	CreateUser(ctx context.Context, p CreateUserParams) (*UserRow, error)

	GetUserTotpByUserID(ctx context.Context, userID ID) (*UserTotpRow, error)
	CreateUserTotp(ctx context.Context, p CreateUserTotpParams) error

	GetUserSessionByUserId(ctx context.Context, userID ID) (*SessionRow, error)
	GetUserSessionById(ctx context.Context, sessionID ID) (*SessionRow, error)
	CreateUserSession(ctx context.Context, p CreateSessionParams) (*SessionRow, error)
	DeleteUserSession(ctx context.Context, sessionID ID) error

	CreateRefreshToken(ctx context.Context, p CreateRefreshTokenParams) error
	DeleteRefreshTokensBySession(ctx context.Context, sessionID ID) error
	GetRefreshTokenByHash(ctx context.Context, hash string) (*RefreshTokenRow, error)

	CreateOTP(ctx context.Context, p CreateOTPParams) error
	GetLatestOTPFromUser(ctx context.Context, userID ID) (*OTPRow, error)
	DeleteUserOTPByUserId(ctx context.Context, userID ID) error
	MarkOTPAsUsed(ctx context.Context, userID ID) error

	UpdateUserPassword(ctx context.Context, userID ID, hashedPassword string) error
}

type DB interface {
	BeginTx(ctx context.Context) (*sql.Tx, Queries, error)
	Queries() Queries
}

type Crypto interface {
	HashPassword(string) (string, error)
	ComparePassword(plain, hash string) error

	HashToken(string) (string, error)
	CompareToken(hash, plain string) error

	GenerateSecureToken(n int) (string, error)
	GenerateNumericOTP(n int) (string, error)

	GenerateJWT(ctx context.Context, sessionID, userID ID) (string, error)
	ValidateJWT(token string) (*JWTClaims, error)
	ValidateJWTFromContext(ctx context.Context) (*JWTClaims, error)
}

type Mailer interface {
	Send(to, subject, html string) error
}

type CookieWriter interface {
	SetAuth(access, refresh string)
	ClearAuth()
}
