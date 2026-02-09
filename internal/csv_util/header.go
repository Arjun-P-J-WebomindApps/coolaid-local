package csv_util

import (
	"fmt"
	"strings"
)

func ParseHeader(records [][]string) ([]string, error) {
	header := records[0]

	seen := map[string]struct{}{}

	for i := range header {
		header[i] = strings.TrimSpace(header[i])
		key := strings.ToLower(header[i])

		if header[i] == "" {
			return nil, fmt.Errorf("empty column name")
		}

		if _, ok := seen[key]; ok {
			return nil, fmt.Errorf("duplicate header: %q", header[i])
		}

		seen[key] = struct{}{}
	}

	return header, nil
}
