package search

import (
	"regexp"
	"strings"
)

var (
	tagRe = regexp.MustCompile("</?mark>")
)

// countMarkedTokens calculates how many tokens inside a snippet
// are wrapped in <mark> tags and returns (markedTokens, totalTokens).
//
// Why we use this:
// Typesense returns highlighted snippets using <mark> tags.
// To compute relevance closeness, we need to measure how much
// of the snippet actually matched the query.
// This function extracts that ratio in token form.
func countMarkedTokens(snippet string) (marked, total int) {

	// Remove highlight tags to count total tokens
	total = len(tokenizePlain(tagRe.ReplaceAllString(snippet, "")))
	if total == 0 {
		return 0, 0
	}

	// Count tokens inside <mark>...</mark>
	parts := strings.Split(snippet, "<mark>")
	for i := 1; i < len(parts); i++ {
		seg := parts[i]
		end := strings.Index(seg, "</mark>")
		if end < 0 {
			continue
		}
		marked += len(tokenizePlain(seg[:end]))
	}

	return marked, total
}

func computeClosenessFromHighlights(highlights []Highlight) float64 {

	sumMarked := 0
	sumTotal := 0

	for _, h := range highlights {
		m, t := countMarkedTokens(h.Snippet)
		sumMarked += m
		sumTotal += t
	}

	if sumTotal == 0 {
		return 0
	}

	return float64(sumMarked) / float64(sumTotal)
}
