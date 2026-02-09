package crypto

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a plain-text password using bcrypt.
//
// - Uses bcrypt with the default cost.
// - Returns a salted, irreversible hash suitable for database storage.
// - The original password cannot be recovered from the hash.
//
// This function should be used ONLY for user passwords.
// Do NOT use it for tokens or OTPs.
func HashPassword(input string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)

	return string(hashedBytes), err
}

// CompareHash compares a bcrypt-hashed password with a plain-text input.
//
// Expected usage:
//   CompareHash(plainPassword, storedHash)
//
// - Returns nil if the password matches.
// - Returns an error if the password is incorrect or hash is invalid.
//
// This function performs a constant-time comparison internally
// to protect against timing attacks.
func CompareHash(input string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(input))
}
