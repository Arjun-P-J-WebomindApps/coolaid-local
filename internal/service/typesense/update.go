package typesense

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/search"
)

func (s *Service) Update(
	ctx context.Context,
	req search.IndexRequest,
) error {
	return s.Index(ctx, req)
}
