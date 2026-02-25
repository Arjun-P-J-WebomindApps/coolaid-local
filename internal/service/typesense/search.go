package typesense

import (
	"context"
	"fmt"
	"time"

	"github.com/typesense/typesense-go/typesense/api"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/search"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
)

func (s *Service) Search(
	ctx context.Context,
	req search.SearchRequest,
) (*search.SearchResponse, error) {

	params := &api.SearchCollectionParams{
		Q:       req.Query,
		QueryBy: req.QueryBy,
		Page:    &req.Page,
		PerPage: &req.PerPage,
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	resp, err := s.client.
		Collection(req.Collection).
		Documents().
		Search(ctx, params)

	if err != nil {
		return nil, fmt.Errorf("typesense search failed: %w", err)
	}

	var hits []api.SearchResultHit
	if resp.Hits != nil {
		hits = *resp.Hits
	}

	result := &search.SearchResponse{
		Found: int(ptr.IntValue(resp.Found)),
		Page:  int(ptr.IntValue(resp.Page)),
		Hits:  make([]search.SearchHit, 0, len(hits)),
	}

	for _, h := range hits {
		result.Hits = append(result.Hits, mapHit(h))
	}

	return result, nil
}
