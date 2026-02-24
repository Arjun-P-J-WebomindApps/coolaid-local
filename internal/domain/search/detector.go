package search

import "strings"

// extractPartToken checks if first token looks like a part number.
// Must contain at least one letter and one digit.
// Allows hyphen.
func extractPartToken(q string) (string, bool) {
	s := strings.TrimSpace(q)
	if s == "" {
		return "", false
	}

	s = strings.Join(strings.Fields(s), " ")
	first := strings.SplitN(s, " ", 2)[0]

	hasLetter := false
	hasDigit := false

	for _, r := range first {
		switch {
		case r >= 'A' && r <= 'Z':
			hasLetter = true
		case r >= 'a' && r <= 'z':
			hasLetter = true
		case r >= '0' && r <= '9':
			hasDigit = true
		case r == '-':
			// allowed
		default:
			return "", false
		}
	}

	if hasLetter && hasDigit {
		return first, true
	}

	return "", false
}
