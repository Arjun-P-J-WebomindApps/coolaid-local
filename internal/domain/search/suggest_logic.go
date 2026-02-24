package search

import "context"

// --------------------------------------------------
// Validation
// --------------------------------------------------

func validateSuggestRequest(req SuggestRequest) error {
	if req.Collection == "" {
		return ErrMissingCollection
	}
	if req.Query == "" {
		return ErrInvalidQuery
	}
	if req.QueryBy == "" {
		return ErrMissingQueryBy
	}
	return nil
}

// --------------------------------------------------
// Decision: Should we do part search?
// --------------------------------------------------

func (s *Service) shouldDoPartSearch(ctx context.Context, token string) bool {

	models, err := s.suggestions.GetSimilarModels(ctx, token)
	if err != nil {
		// Safe fallback: assume part search if DB fails
		return true
	}

	return len(models) == 0
}

// --------------------------------------------------
// Build Part Suggestions
// --------------------------------------------------

func (s *Service) buildPartSuggestions(ctx context.Context, token string) (*SuggestResponse, error) {

	parts, _ := s.suggestions.GetPartSuggestions(ctx, token)
	oems, _ := s.suggestions.GetOemSuggestions(ctx, token)
	vendors, _ := s.suggestions.GetVendorSuggestions(ctx, token)

	suggestions := interleaveSuggestions(parts, oems, vendors, 15)

	hits := buildHitsFromSuggestions(suggestions)

	return &SuggestResponse{
		Hits: hits,
	}, nil
}

// --------------------------------------------------
// Fallback to Typesense Search
// --------------------------------------------------

func (s *Service) fallbackSuggestSearch(ctx context.Context, req SuggestRequest) (*SuggestResponse, error) {

	searchResp, err := s.port.Search(ctx, SearchRequest{
		Collection: req.Collection,
		Query:      req.Query,
		QueryBy:    req.QueryBy,
		Page:       1,
		PerPage:    15,
	})

	if err != nil {
		return nil, err
	}

	return &SuggestResponse{
		Hits: searchResp.Hits,
	}, nil
}

// --------------------------------------------------
// Build SearchHit from string suggestions
// --------------------------------------------------

func buildHitsFromSuggestions(values []string) []SearchHit {

	hits := make([]SearchHit, 0, len(values))

	for _, v := range values {
		hits = append(hits, SearchHit{
			Document: map[string]any{
				"id":      "SearchByPartNo",
				"part_no": v,
			},
		})
	}

	return hits
}

// --------------------------------------------------
// Interleave Logic
// --------------------------------------------------

func interleaveSuggestions(parts, oems, vendors []string, max int) []string {

	result := make([]string, 0, max)

	i, j, k := 0, 0, 0

	for len(result) < max {
		added := false

		if i < len(parts) {
			result = append(result, parts[i])
			i++
			added = true
		}
		if len(result) >= max {
			break
		}

		if j < len(oems) {
			result = append(result, oems[j])
			j++
			added = true
		}
		if len(result) >= max {
			break
		}

		if k < len(vendors) {
			result = append(result, vendors[k])
			k++
			added = true
		}

		if !added {
			break
		}
	}

	return result
}
