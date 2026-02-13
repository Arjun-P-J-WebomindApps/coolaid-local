package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetChillerUnitByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.ChillerUnitRow, error) {

	row, err := t.q.GetChillerUnitByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapChillerUnitRow(row), nil
}

func (t *techSpecQueries) CreateChillerUnit(
	ctx context.Context,
	p techspec.CreateChillerUnitParams,
) (*techspec.ChillerUnitRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreateChillerUnit(ctx, sqlc.CreateChillerUnitParams{
		ID:      uID,
		PartNo:  p.PartNo,
		Type:    sqlnull.String(p.Type),
		Voltage: sqlnull.String(p.Voltage),
		Notes:   sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapChillerUnitRow(row), nil
}

func (t *techSpecQueries) UpdateChillerUnitByPartNo(
	ctx context.Context,
	p techspec.UpdateChillerUnitParams,
) (*techspec.ChillerUnitRow, error) {

	row, err := t.q.UpdateChillerUnitByPartNo(ctx, sqlc.UpdateChillerUnitByPartNoParams{
		PartNo:  p.PartNo,
		Type:    sqlnull.String(p.Type),
		Voltage: sqlnull.String(p.Voltage),
		Notes:   sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapChillerUnitRow(row), nil
}

func (t *techSpecQueries) DeleteChillerUnitByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeleteChillerUnitByPartNo(ctx, partNo)
}

func mapChillerUnitRow(r sqlc.ChillerUnit) *techspec.ChillerUnitRow {

	return &techspec.ChillerUnitRow{
		ID:      r.ID.String(),
		PartNo:  r.PartNo,
		Type:    sqlnull.StringPtr(r.Type),
		Voltage: sqlnull.StringPtr(r.Voltage),
		Notes:   sqlnull.StringPtr(r.Notes),
	}
}
