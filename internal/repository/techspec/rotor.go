package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetRotorByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.RotorRow, error) {

	row, err := t.q.GetRotorByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapRotorRow(row), nil
}

func (t *techSpecQueries) CreateRotor(
	ctx context.Context,
	p techspec.CreateRotorParams,
) (*techspec.RotorRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreateRotor(ctx, sqlc.CreateRotorParams{
		ID:                uID,
		PartNo:            p.PartNo,
		PulleyRibs:        sqlnull.String(p.PulleyRibs),
		PulleySize:        sqlnull.String(p.PulleySize),
		CompressorDetails: sqlnull.String(p.CompressorDetails),
		Notes:             sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapRotorRow(row), nil
}

func (t *techSpecQueries) UpdateRotorByPartNo(
	ctx context.Context,
	p techspec.UpdateRotorParams,
) (*techspec.RotorRow, error) {

	row, err := t.q.UpdateRotorByPartNo(ctx, sqlc.UpdateRotorByPartNoParams{
		PartNo:            p.PartNo,
		PulleyRibs:        sqlnull.String(p.PulleyRibs),
		PulleySize:        sqlnull.String(p.PulleySize),
		CompressorDetails: sqlnull.String(p.CompressorDetails),
		Notes:             sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapRotorRow(row), nil
}

func (t *techSpecQueries) DeleteRotorByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeleteRotorByPartNo(ctx, partNo)
}

func mapRotorRow(r sqlc.Rotor) *techspec.RotorRow {

	return &techspec.RotorRow{
		ID:                r.ID.String(),
		PartNo:            r.PartNo,
		PulleyRibs:        sqlnull.StringPtr(r.PulleyRibs),
		PulleySize:        sqlnull.StringPtr(r.PulleySize),
		CompressorDetails: sqlnull.StringPtr(r.CompressorDetails),
		Notes:             sqlnull.StringPtr(r.Notes),
	}
}
