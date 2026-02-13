package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
)

func (s *Service) ListByName(
	ctx context.Context,
	name string,
) ([]Category, error) {

	rows, err := s.DB.Queries().GetCategoriesByName(ctx, name)
	if err != nil {
		return nil, ErrInternal
	}

	out := make([]Category, 0, len(rows))
	for _, r := range rows {
		out = append(out, *mapRowToModel(&r))
	}

	return out, nil
}

func (s *Service) GetByID(
	ctx context.Context,
	id string,
) (*Category, error) {

	if id == "" {
		return nil, ErrInvalidInput
	}

	row, err := s.DB.Queries().GetCategoryByID(ctx, shared.ID(id))
	if err != nil {
		return nil, ErrCategoryNotFound
	}

	return mapRowToModel(row), nil
}

func (s *Service) GetByName(
	ctx context.Context,
	name string,
) (*Category, error) {

	if name == "" {
		return nil, ErrInvalidInput
	}

	row, err := s.DB.Queries().GetCategoryByName(ctx, name)
	if err != nil {
		return nil, ErrCategoryNotFound
	}

	return mapRowToModel(row), nil
}
