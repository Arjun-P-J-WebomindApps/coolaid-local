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

	//Typesense
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

	hits := make([]api.SearchResultHit, 0, 0)

	if resp.Hits != nil {
		hits = *resp.Hits
	}

	result := &search.SearchResponse{
		Found: int(ptr.IntValue(resp.Found)),
		Page:  int(ptr.IntValue(resp.Page)),
		Hits:  make([]search.SearchHit, 0, len(hits)),
	}

	for _, h := range hits {

		hlts := make([]api.SearchHighlight, 0, 0)

		if h.Highlights != nil {
			hlts = *h.Highlights
		}

		highlights := make([]search.Highlight, 0, len(hlts))

		for _, hl := range hlts {

			matchedTokens := make([]string, 0)
			if hl.MatchedTokens != nil {
				matchedTokens = make([]string, 0, len(*hl.MatchedTokens))
				for _, t := range *hl.MatchedTokens {
					matchedTokens = append(matchedTokens, fmt.Sprint(t))
				}
			}

			highlights = append(highlights, search.Highlight{
				Field:         ptr.String(hl.Field),
				Snippet:       ptr.String(hl.Snippet),
				MatchedTokens: matchedTokens,
			})
		}

		result.Hits = append(result.Hits, search.SearchHit{
			Document:   h.Document,
			Highlights: highlights,
		})
	}

	return result, nil
}
