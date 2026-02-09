package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
)

func (s *Service) Delete(ctx context.Context, id string) error {
	if err := s.DB.Queries().DeleteCategory(ctx, shared.ID(id)); err != nil {
		return ErrCategoryNotFound
	}
	return nil
}
