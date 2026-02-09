package auth

import "time"

// ---------- USER ----------

type CreateUserParams struct {
	ID        ID
	Name      string
	Username  string
	Email     string
	Password  string
	Mobile    string
	Role      string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ---------- TOTP ----------

type CreateUserTotpParams struct {
	UserID ID
	Secret string
}

// ---------- SESSION ----------

type CreateSessionParams struct {
	UserID    ID
	SessionID ID
	IP        string
	UserAgent string
	ExpiresAt time.Time
}

// ---------- REFRESH TOKEN ----------

type CreateRefreshTokenParams struct {
	UserID    ID
	SessionID ID
	TokenHash string
	ExpiresAt time.Time
}

// ---------- OTP ----------

type CreateOTPParams struct {
	UserID    ID
	OtpCode   string
	ExpiresAt time.Time
}
