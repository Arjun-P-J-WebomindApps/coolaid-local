package search

import (
	"context"
)

type Service struct {
	port Port
	DB   DB
}

func NewService(port Port, db DB) *Service {
	return &Service{
		port: port,
		DB:   db,
	}
}

func (s *Service) SearchProducts(ctx context.Context, req SearchRequest) (*SearchResponse, error) {
	// Validate
	if err := validateSearchRequest(req); err != nil {
		return nil, err
	}

	req = normalizeSearchRequest(req)

	// Detect numeric-driven token
	token, isPriority := firstNumericDrivenToken(req.Query)

	if isPriority && s.shouldDoPartSearch(ctx, token) {

		// Build DB-driven part suggestions
		parts, _ := s.DB.Queries().GetPartSuggestions(ctx, token)
		oems, _ := s.DB.Queries().GetOemSuggestions(ctx, token)
		vendors, _ := s.DB.Queries().GetVendorSuggestions(ctx, token)

		values := interleaveSuggestions(parts, oems, vendors, req.PerPage)

		hits := buildHitsFromSuggestions(values)

		return &SearchResponse{
			Found: len(hits),
			Page:  1,
			Hits:  hits,
		}, nil
	}

	// Normal Typesense Search
	resp, err := s.port.Search(ctx, req)
	if err != nil {
		return nil, err
	}

	//  Rerank
	resp = s.rerankIfNeeded(resp, req)

	return resp, nil

}
