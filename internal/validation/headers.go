package validation

import (
	"fmt"
	"strings"
)

func ValidateCSVHeader(
	header []string,
	expected []string,
	requireOrder bool,
) error {

	normalize := func(s string) string {
		return strings.ToLower(strings.TrimSpace(s))
	}

	if requireOrder {
		if len(header) < len(expected) {
			return fmt.Errorf("header length mismatch")
		}
		for i := range expected {
			if normalize(header[i]) != normalize(expected[i]) {
				return fmt.Errorf(
					"header mismatch at position %d: got %q want %q",
					i+1, header[i], expected[i],
				)
			}
		}
		return nil
	}

	seen := map[string]struct{}{}
	for _, h := range header {
		seen[normalize(h)] = struct{}{}
	}

	for _, e := range expected {
		if _, ok := seen[normalize(e)]; !ok {
			return fmt.Errorf("missing required column %q", e)
		}
	}

	return nil
}
