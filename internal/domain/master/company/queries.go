package company

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/shared"
	"github.com/webomindapps-dev/coolaid-backend/oplog"
)

func (s *Service) GetByID(ctx context.Context, id string) (*Company, error) {
	row, err := s.DB.Queries().GetCompanyByID(ctx, shared.ID(id))
	if err != nil {
		return nil, ErrCompanyNotFound
	}
	return mapRowToModel(row), nil
}

func (s *Service) ListByName(ctx context.Context, name *string) ([]Company, error) {
	rows, err := s.DB.Queries().GetCompaniesByName(
		ctx,
		shared.ValueOrEmpty(name),
	)

	if err != nil {
		oplog.Error(ctx, err)
		return nil, ErrInternal
	}

	out := make([]Company, 0, len(rows))
	for _, r := range rows {
		out = append(out, *mapRowToModel(&r))
	}

	return out, nil
}
