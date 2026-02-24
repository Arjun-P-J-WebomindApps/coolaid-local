package typesense

import (
	"context"
	"fmt"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/search"
)

func (s *Service) Index(
	ctx context.Context,
	req search.IndexRequest,
) error {

	_, err := s.client.
		Collection(req.Collection).
		Documents().
		Upsert(ctx, req.Payload)

	if err != nil {
		return fmt.Errorf("typesense index failed: %w", err)
	}

	return nil
}
