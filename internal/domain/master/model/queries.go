package models

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
)

func (s *Service) GetByID(ctx context.Context, id string) (*Model, error) {
	row, err := s.DB.Queries().GetModelByID(ctx, shared.ID(id))
	if err != nil {
		return nil, ErrModelNotFound
	}
	return mapRowToModel(row), nil
}

func (s *Service) ListByName(ctx context.Context, modelName *string, companyName *string) ([]ModelWithCompanyRow, error) {
	rows, err := s.DB.Queries().GetModelsByCompanyAndModelNames(
		ctx,
		shared.ValueOrEmpty(modelName),
		shared.ValueOrEmpty(companyName),
	)
	if err != nil {
		return nil, ErrInternal
	}

	out := make([]ModelWithCompanyRow, 0, len(rows))
	for _, r := range rows {
		out = append(out, *mapRowToCompanyModel(&r))
	}

	return out, nil
}
