package auth

import "time"

/*
Domain-level representations of DB rows.
These are NOT sqlc / ORM models.
The DB implementation adapts to these.
*/

type UserRow struct {
	ID       ID
	Name     string
	Username string
	Email    string
	Password string
	Mobile   string
	Role     string
	IsActive bool
}

type UserTotpRow struct {
	UserID ID
	Secret string
}

type SessionRow struct {
	SessionID ID
	UserID    ID
	ExpiresAt time.Time
	RevokedAt *time.Time
}

type RefreshTokenRow struct {
	ID        ID
	UserID    ID
	SessionID ID
	TokenHash string
	ExpiresAt time.Time
	RevokedAt *time.Time
	CreatedAt time.Time
}

type OTPRow struct {
	UserID    ID
	OtpCode   string
	ExpiresAt time.Time
	IsUsed    bool
}
