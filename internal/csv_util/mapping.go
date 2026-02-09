package csv_util

import "strings"

func BuildHeaderIndex(headers []string) map[string]int {
	idx := make(map[string]int, len(headers))

	for i, h := range headers {
		key := strings.ToLower(strings.TrimSpace(h))
		idx[key] = i
	}

	return idx
}
