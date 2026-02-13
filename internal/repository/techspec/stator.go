package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetStatorByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.StatorRow, error) {

	row, err := t.q.GetStatorByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapStatorRow(row), nil
}

func (t *techSpecQueries) CreateStator(
	ctx context.Context,
	p techspec.CreateStatorParams,
) (*techspec.StatorRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreateStator(ctx, sqlc.CreateStatorParams{
		ID:                uID,
		PartNo:            p.PartNo,
		Voltage:           sqlnull.String(p.Voltage),
		CompressorDetails: sqlnull.String(p.CompressorDetails),
		Notes:             sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapStatorRow(row), nil
}

func (t *techSpecQueries) UpdateStatorByPartNo(
	ctx context.Context,
	p techspec.UpdateStatorParams,
) (*techspec.StatorRow, error) {

	row, err := t.q.UpdateStatorByPartNo(ctx, sqlc.UpdateStatorByPartNoParams{
		PartNo:            p.PartNo,
		Voltage:           sqlnull.String(p.Voltage),
		CompressorDetails: sqlnull.String(p.CompressorDetails),
		Notes:             sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapStatorRow(row), nil
}

func (t *techSpecQueries) DeleteStatorByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeleteStatorByPartNo(ctx, partNo)
}

func mapStatorRow(r sqlc.Stator) *techspec.StatorRow {

	return &techspec.StatorRow{
		ID:                r.ID.String(),
		PartNo:            r.PartNo,
		Voltage:           sqlnull.StringPtr(r.Voltage),
		CompressorDetails: sqlnull.StringPtr(r.CompressorDetails),
		Notes:             sqlnull.StringPtr(r.Notes),
	}
}
