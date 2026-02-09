package company

import (
	"context"
	"time"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
)

func (s *Service) Create(
	ctx context.Context,
	input CreateCompanyInput,
) (*Company, error) {

	id := shared.NewID().String()

	row, err := s.DB.Queries().CreateCompany(ctx, CreateCompanyParams{
		ID:        id,
		Name:      input.Name,
		Status:    input.Status,
		ImageURL:  input.ImageURL,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, ErrInternal
	}

	return mapRowToModel(row), nil
}
