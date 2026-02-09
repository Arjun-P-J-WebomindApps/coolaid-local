package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
)

func (s *Service) Create(
	ctx context.Context,
	input CreateCategoryInput,
) (*Category, error) {

	id := shared.NewID().String()

	row, err := s.DB.Queries().CreateCategory(ctx, CreateCategoryParams{
		ID:    id,
		Name:  input.Name,
		Image: input.Image,
	})
	if err != nil {
		return nil, ErrCategoryExists
	}

	return mapRowToModel(row), nil
}
