package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetCondenserByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.CondenserRow, error) {

	row, err := t.q.GetCondenserByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapCondenserRow(row), nil
}

func (t *techSpecQueries) CreateCondenser(
	ctx context.Context,
	p techspec.CreateCondenserParams,
) (*techspec.CondenserRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreateCondenser(ctx, sqlc.CreateCondenserParams{
		ID:             uID,
		PartNo:         p.PartNo,
		Size:           sqlnull.String(p.Size),
		PipeConnector:  sqlnull.String(p.PipeConnector),
		Drier:          sqlnull.String(p.Drier),
		PressureSwitch: sqlnull.String(p.PressureSwitch),
		Notes:          sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapCondenserRow(row), nil
}

func (t *techSpecQueries) UpdateCondenserByPartNo(
	ctx context.Context,
	p techspec.UpdateCondenserParams,
) (*techspec.CondenserRow, error) {

	row, err := t.q.UpdateCondenserByPartNo(ctx, sqlc.UpdateCondenserByPartNoParams{
		PartNo:         p.PartNo,
		Size:           sqlnull.String(p.Size),
		PipeConnector:  sqlnull.String(p.PipeConnector),
		Drier:          sqlnull.String(p.Drier),
		PressureSwitch: sqlnull.String(p.PressureSwitch),
		Notes:          sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapCondenserRow(row), nil
}

func (t *techSpecQueries) DeleteCondenserByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeleteCondenserByPartNo(ctx, partNo)
}

func mapCondenserRow(r sqlc.Condenser) *techspec.CondenserRow {

	return &techspec.CondenserRow{
		ID:             r.ID.String(),
		PartNo:         r.PartNo,
		Size:           sqlnull.StringPtr(r.Size),
		PipeConnector:  sqlnull.StringPtr(r.PipeConnector),
		Drier:          sqlnull.StringPtr(r.Drier),
		PressureSwitch: sqlnull.StringPtr(r.PressureSwitch),
		Notes:          sqlnull.StringPtr(r.Notes),
	}
}
