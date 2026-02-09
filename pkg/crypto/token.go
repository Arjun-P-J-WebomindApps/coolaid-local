package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
)

// GenerateSecureToken generates a cryptographically secure random token.
//
// - nBytes defines the number of random bytes to generate.
// - The returned value is hex-encoded for safe storage and transport.
// - Used for refresh tokens, password reset tokens, etc.
//
// This function does NOT store the token anywhere.
// Callers are responsible for hashing it before persistence.
func GenerateSecureToken(nBytes int) (string, error) {
	b := make([]byte, nBytes)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

// HashToken creates a deterministic SHA-256 hash of the input string.
//
// - Used for hashing sensitive tokens (refresh tokens, OTPs) before storing.
// - Deterministic hashing allows lookup by hash without storing raw values.
// - This is NOT a password hash (bcrypt is used for passwords).
//
// Same input will always produce the same hash.
func HashToken(input string) (string, error) {
	sum := sha256.Sum256([]byte(input))
	return hex.EncodeToString(sum[:]), nil
}

// CompareToken compares a stored token hash with a plain-text token.
//
// Expected usage:
//
//	CompareToken(hashedFromDB, rawTokenFromUser)
//
// - Hashes the plain token using SHA-256.
// - Compares hashes using constant-time comparison to prevent timing attacks.
// - Returns nil if tokens match, error otherwise.
//
// This function should be used for refresh tokens and OTP verification.
func CompareToken(hashed, plain string) error {
	sum := sha256.Sum256([]byte(plain))
	candidate := hex.EncodeToString(sum[:])

	if subtle.ConstantTimeCompare([]byte(hashed), []byte(candidate)) == 1 {
		return nil
	}
	return errors.New("token mismatch")
}
