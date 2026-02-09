package brand

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
)

func (s *Service) Create(
	ctx context.Context,
	input CreateBrandInput,
) (*Brand, error) {

	_, err := s.DB.Queries().GetBrandByName(ctx, input.Name)
	if err == nil {
		return nil, ErrBrandExists
	}

	row, err := s.DB.Queries().CreateBrand(ctx, CreateBrandParams{
		ID:    shared.NewID().String(),
		Name:  input.Name,
		Image: input.Image,
	})
	if err != nil {
		return nil, ErrInternal
	}

	return mapRowToModel(row), nil
}
