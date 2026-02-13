package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetCabinFilterByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.CabinFilterRow, error) {

	row, err := t.q.GetCabinFilterByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapCabinFilterRow(row), nil
}

func (t *techSpecQueries) CreateCabinFilter(
	ctx context.Context,
	p techspec.CreateCabinFilterParams,
) (*techspec.CabinFilterRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreateCabinFilter(ctx, sqlc.CreateCabinFilterParams{
		ID:         uID,
		PartNo:     p.PartNo,
		Type:       sqlnull.String(p.Type),
		Dimensions: sqlnull.String(p.Dimensions),
		Material:   sqlnull.String(p.Material),
		Notes:      sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapCabinFilterRow(row), nil
}

func (t *techSpecQueries) UpdateCabinFilterByPartNo(
	ctx context.Context,
	p techspec.UpdateCabinFilterParams,
) (*techspec.CabinFilterRow, error) {

	row, err := t.q.UpdateCabinFilterByPartNo(ctx, sqlc.UpdateCabinFilterByPartNoParams{
		PartNo:     p.PartNo,
		Type:       sqlnull.String(p.Type),
		Dimensions: sqlnull.String(p.Dimensions),
		Material:   sqlnull.String(p.Material),
		Notes:      sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapCabinFilterRow(row), nil
}

func (t *techSpecQueries) DeleteCabinFilterByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeleteCabinFilterByPartNo(ctx, partNo)
}

func mapCabinFilterRow(r sqlc.CabinFilter) *techspec.CabinFilterRow {

	return &techspec.CabinFilterRow{
		ID:         r.ID.String(),
		PartNo:     r.PartNo,
		Type:       sqlnull.StringPtr(r.Type),
		Dimensions: sqlnull.StringPtr(r.Dimensions),
		Material:   sqlnull.StringPtr(r.Material),
		Notes:      sqlnull.StringPtr(r.Notes),
	}
}
