package models

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/ptr"
	"github.com/webomindapps-dev/coolaid-backend/oplog"
)

func (s *Service) Update(
	ctx context.Context,
	input UpdateModelInput,
) (*Model, error) {

	// 1. Basic validation (domain rule)
	if input.ID == "" || ptr.String(input.Name) == "" || ptr.String(input.CompanyName) == "" {
		return nil, ErrInvalidInput
	}

	// 2. Resolve company (domain rule)
	company, err := s.CompanyService.DB.Queries().GetCompanyByName(ctx, ptr.String(input.CompanyName))
	if err != nil {
		return nil, ErrCompanyNotFound
	}

	// 3. Ensure model exists (domain rule)
	existing, err := s.DB.Queries().GetModelByID(ctx, shared.ID(input.ID))
	if err != nil {
		return nil, ErrModelNotFound
	}

	// 4. Enforce uniqueness (exclude self)
	_, err = s.DB.Queries().GetModelsByCompanyAndModelNames(
		ctx,
		ptr.String(input.Name),
		ptr.String(input.CompanyName),
	)
	if err == nil && existing.Name != ptr.String(input.Name) {
		return nil, ErrModelExists
	}

	// 5. Update model
	row, err := s.DB.Queries().UpdateModel(ctx, UpdateModelParams{
		ID:        input.ID,
		CompanyID: &company.ID,
		Name:      input.Name,
		ImageURL:  input.ImageURL,
	})
	if err != nil {
		oplog.Error(ctx,
			"UpdateModel persistence failure",
			"model_id=", input.ID,
			"company_id=", company.ID,
			"name=", input.Name,
			"err=", err,
		)
		return nil, ErrInternal
	}

	return mapRowToModel(row), nil
}
