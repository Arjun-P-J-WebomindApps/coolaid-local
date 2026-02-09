package auth

import "time"

type AuthResult struct {
	UserID   ID
	Username string
}

type RegisterResult struct {
	UserID     ID
	Username   string
	Email      string
	OtpSecret  string
	OtpAuthURL string
}

type SessionMeta struct {
	IP        string
	UserAgent string
}

type JWTClaims struct {
	SessionID string
	UserID    ID
	ExpiresAt time.Time
}
