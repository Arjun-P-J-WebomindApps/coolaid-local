package search

import "context"

type Service struct {
	port        Port
	suggestions SuggestionRepository
}

func NewService(port Port, repo SuggestionRepository) *Service {
	return &Service{
		port:        port,
		suggestions: repo,
	}
}

func (s *Service) Search(ctx context.Context, req SearchRequest) (*SearchResponse, error) {

	if err := validateSearchRequest(req); err != nil {
		return nil, err
	}

	req = normalizeSearchRequest(req)

	resp, err := s.executeSearch(ctx, req)
	if err != nil {
		return nil, err
	}

	return s.rerankIfNeeded(resp, req), nil
}
