package crypto

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

// GenerateNumericOTPString generates a cryptographically secure numeric OTP
// of the specified length and returns it as a zero-padded string.
//
// - Uses crypto/rand for secure random number generation.
// - Ensures the OTP is exactly `length` digits (leading zeros preserved).
// - Returns an error if the length is invalid or random generation fails.
//
// This function is intended for OTP use cases such as:
// - password reset codes
// - email / SMS verification codes
//
// The returned OTP should never be stored in plain text.
// Always hash it before persisting.
func GenerateNumericOTPString(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("invalid otp length")
	}

	max := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length)), nil)

	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}

	// Zero-pad to required length
	return fmt.Sprintf("%0*d", length, n.Int64()), nil
}
