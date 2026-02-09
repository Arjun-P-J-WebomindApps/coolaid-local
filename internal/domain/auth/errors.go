package auth

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidOTP         = errors.New("invalid otp")
	ErrExpiredOTP         = errors.New("otp expired")

	ErrUserExists   = errors.New("user already exists")
	ErrUnauthorized = errors.New("unauthorized")
	ErrMFARequired  = errors.New("mfa required")

	ErrSessionExpired = errors.New("session expired")
	ErrSessionRevoked = errors.New("session revoked")

	ErrInternal = errors.New("internal error")

	ErrRefreshExpired = errors.New("refresh expired")
)
