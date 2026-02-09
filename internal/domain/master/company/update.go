package company

import (
	"context"
	"time"
)

func (s *Service) Update(
	ctx context.Context,
	input UpdateCompanyInput,
) (*Company, error) {

	row, err := s.DB.Queries().UpdateCompany(ctx, UpdateCompanyParams{
		ID:        input.ID,
		Name:      input.Name,
		Status:    input.Status,
		ImageURL:  input.ImageURL,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, ErrCompanyNotFound
	}

	return mapRowToModel(row), nil
}
