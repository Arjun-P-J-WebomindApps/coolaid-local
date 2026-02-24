package search

import "strings"

// tokenizePlain splits string into lowercase tokens.
// Keeps hyphen intact.
func tokenizePlain(s string) []string {
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "" {
		return nil
	}

	s = strings.Join(strings.Fields(s), " ")
	return strings.Fields(s)
}
