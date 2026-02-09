package brand

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
)

func (s *Service) GetByID(ctx context.Context, id string) (*Brand, error) {
	row, err := s.DB.Queries().GetBrandByID(ctx, shared.ID(id))
	if err != nil {
		return nil, ErrBrandNotFound
	}
	return mapRowToModel(row), nil
}

func (s *Service) ListByName(ctx context.Context, name *string) ([]Brand, error) {
	rows, err := s.DB.Queries().GetBrandListByName(
		ctx,
		shared.ValueOrEmpty(name),
	)
	if err != nil {
		return nil, ErrInternal
	}

	out := make([]Brand, 0, len(rows))
	for _, r := range rows {
		out = append(out, *mapRowToModel(&r))
	}

	return out, nil
}
