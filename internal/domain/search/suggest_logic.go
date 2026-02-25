package search

import (
	"context"
)

// --------------------------------------------------
// Decision: Should we do part search?
// --------------------------------------------------

func (s *Service) shouldDoPartSearch(ctx context.Context, token string) bool {

	models, err := s.DB.Queries().GetSimilarModels(ctx, token)
	if err != nil {
		// Safe fallback: assume part search if DB fails
		return true
	}

	return len(models) == 0
}

// --------------------------------------------------
// Build SearchHit from string suggestions
// --------------------------------------------------

func buildHitsFromSuggestions(values []string) []SearchHit {

	hits := make([]SearchHit, 0, len(values))

	for _, v := range values {

		hits = append(hits, SearchHit{
			Document: ProductSearchDocument{
				ID:       "SearchByPartNo",
				Company:  "",
				Model:    "",
				Category: "",
				Brand:    "",
				PartNo:   v,
			},
			Highlights: []Highlight{
				{
					Field:         "part_no",
					Snippet:       "",
					MatchedTokens: []string{""},
				},
			},
		})
	}

	return hits
}

func buildSuggestionToken(s PartSuggestionResponse) string {
	if s.MatchedValue == "" || s.MatchedValue == s.PartNo {
		return s.PartNo + ":" + s.PartNo
	}

	return s.MatchedValue + ":" + s.PartNo
}

// --------------------------------------------------
// Interleave Logic
// --------------------------------------------------

func interleaveSuggestions(parts, oems, vendors []PartSuggestionResponse, max int) []string {

	result := make([]string, 0, max)

	i, j, k := 0, 0, 0

	for len(result) < max {
		added := false

		if i < len(parts) {
			result = append(result, buildSuggestionToken(parts[i]))
			i++
			added = true
		}
		if len(result) >= max {
			break
		}

		if j < len(oems) {
			result = append(result, buildSuggestionToken(oems[j]))
			j++
			added = true
		}
		if len(result) >= max {
			break
		}

		if k < len(vendors) {
			result = append(result, buildSuggestionToken(vendors[k]))
			k++
			added = true
		}

		if !added {
			break
		}
	}

	return result
}
