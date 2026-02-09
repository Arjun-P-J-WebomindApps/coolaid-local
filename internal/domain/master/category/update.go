package category

import "context"

func (s *Service) Update(
	ctx context.Context,
	input UpdateCategoryInput,
) (*Category, error) {

	row, err := s.DB.Queries().UpdateCategory(ctx, UpdateCategoryParams{
		ID:    input.ID,
		Name:  input.Name,
		Image: input.Image,
	})
	if err != nil {
		return nil, ErrCategoryNotFound
	}

	return mapRowToModel(row), nil
}
