package search

import (
	"regexp"
	"strings"
)

var (
	hasDigit           = regexp.MustCompile(`[0-9]`)
	validPart          = regexp.MustCompile(`^[a-zA-Z0-9\-\/]+$`)
	splitRe            = regexp.MustCompile("[ \t\n\r/_-]+")
	fieldsForCloseness = []string{"model", "category", "brand", "company", "part_no"}
)

// firstNumericDrivenToken detects a token that looks like a part number.
//
// Why we use this:
// In Coolaid search, part numbers are primarily numeric-driven.
// They always contain digits and may optionally include letters,
// hyphens, or slashes. This helper prioritizes part-number search
// over generic text search.
func firstNumericDrivenToken(q string) (string, bool) {
	for _, tok := range tokenizePlain(q) {

		// Must contain at least one digit
		if !hasDigit.MatchString(tok) {
			continue
		}

		// Must contain only valid part-number characters
		if !validPart.MatchString(tok) {
			continue
		}

		// Optional: enforce minimum length to avoid noise like "3"
		if len(tok) < 3 {
			continue
		}

		return tok, true
	}
	return "", false
}

// tokenizePlain normalizes and tokenizes input text into lowercase words.
//
// Why we use this:
// Search queries can contain inconsistent casing, extra spaces,
// or formatting noise. This function standardizes input so that
// all downstream search logic (part detection, scoring, matching)
// operates on clean, predictable tokens.
func tokenizePlain(s string) []string {
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "" {
		return nil
	}

	// Collapse multiple spaces into one
	s = strings.Join(strings.Fields(s), " ")

	// Split only by spaces (hyphen preserved intentionally)
	return strings.Fields(s)
}
