package search

import (
	"context"
	"sort"
)

func validateSearchRequest(req SearchRequest) error {
	if req.Collection == "" {
		return ErrMissingCollection
	}
	if req.Query == "" {
		return ErrInvalidQuery
	}
	return nil
}

func normalizeSearchRequest(req SearchRequest) SearchRequest {

	if req.Page <= 0 {
		req.Page = 1
	}

	if req.PerPage <= 0 {
		req.PerPage = 20
	}

	if req.PerPage > 100 {
		req.PerPage = 100
	}

	return req
}

func (s *Service) executeSearch(ctx context.Context, req SearchRequest) (*SearchResponse, error) {
	return s.port.Search(ctx, req)
}

func (s *Service) rerankIfNeeded(resp *SearchResponse, req SearchRequest) *SearchResponse {

	// Only rerank if multi-field search
	if req.QueryBy == "" || len(resp.Hits) == 0 {
		return resp
	}

	type scored struct {
		hit   SearchHit
		score float64
	}

	scoredHits := make([]scored, 0, len(resp.Hits))

	for _, h := range resp.Hits {
		score := computeClosenessFromHighlights(h.Highlights)
		scoredHits = append(scoredHits, scored{
			hit:   h,
			score: score,
		})
	}

	sort.SliceStable(scoredHits, func(i, j int) bool {
		return scoredHits[i].score > scoredHits[j].score
	})

	newHits := make([]SearchHit, 0, len(scoredHits))
	for _, s := range scoredHits {
		newHits = append(newHits, s.hit)
	}

	resp.Hits = newHits
	return resp
}
