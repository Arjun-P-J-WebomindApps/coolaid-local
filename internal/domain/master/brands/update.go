package brand

import "context"

func (s *Service) Update(
	ctx context.Context,
	input UpdateBrandInput,
) (*Brand, error) {

	row, err := s.DB.Queries().UpdateBrand(ctx, UpdateBrandParams{
		ID:    input.ID,
		Name:  input.Name,
		Image: input.Image,
	})
	if err != nil {
		return nil, ErrBrandNotFound
	}

	return mapRowToModel(row), nil
}
