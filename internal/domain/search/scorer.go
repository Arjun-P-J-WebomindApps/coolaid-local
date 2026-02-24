package search

import (
	"regexp"
	"strings"
)

var tagRe = regexp.MustCompile("</?mark>")

var fieldsForCloseness = []string{
	"model",
	"category",
	"brand",
	"company",
	"part_no",
}

func countMarkedTokens(snippet string) (marked int) {
	parts := strings.Split(snippet, "<mark>")
	for i := 1; i < len(parts); i++ {
		end := strings.Index(parts[i], "</mark>")
		if end < 0 {
			continue
		}
		marked += len(tokenizePlain(parts[i][:end]))
	}
	return
}

func markClosenessScore(highlights []Highlight) float64 {

	score := 0.0

	for _, h := range highlights {
		marked := countMarkedTokens(h.Snippet)

		switch h.Field {
		case "part_no":
			score += float64(marked) * 5
		case "model":
			score += float64(marked) * 3
		case "brand":
			score += float64(marked) * 2
		default:
			score += float64(marked)
		}
	}

	return score
}
