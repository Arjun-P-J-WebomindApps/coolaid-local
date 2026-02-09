package models

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
	"github.com/webomindapps-dev/coolaid-backend/oplog"
)

func (s *Service) Create(
	ctx context.Context,
	input CreateModelInput,
) (*Model, error) {

	// 1. Basic validation (domain rule)
	if input.Name == "" || input.CompanyName == "" {
		return nil, ErrInvalidInput
	}

	// 2. Resolve company (domain rule)
	company, err := s.CompanyService.DB.Queries().GetCompanyByName(ctx, input.CompanyName)
	if err != nil {
		return nil, ErrCompanyNotFound
	}

	// 3. Enforce uniqueness (domain rule)
	_, err = s.DB.Queries().GetModelsByCompanyAndModelNames(
		ctx,
		company.ID,
		input.Name,
	)
	if err == nil {
		return nil, ErrModelExists
	}

	// 4. Create model
	id := shared.NewID().String()

	row, err := s.DB.Queries().CreateModel(ctx, CreateModelParams{
		ID:        id,
		CompanyID: company.ID,
		Name:      input.Name,
		ImageURL:  input.ImageURL,
	})
	if err != nil {
		// invariant break → log here if you want
		oplog.Error(ctx,
			"CreateModel persistence failure",
			"company_id=", company.ID,
			"name=", input.Name,
			"err=", err,
		)
		return nil, ErrInternal
	}

	return mapRowToModel(row), nil
}
