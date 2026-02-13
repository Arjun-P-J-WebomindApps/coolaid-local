package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetClutchAssyByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.ClutchAssyRow, error) {

	row, err := t.q.GetClutchAssyByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapClutchAssyRow(row), nil
}

func (t *techSpecQueries) CreateClutchAssy(
	ctx context.Context,
	p techspec.CreateClutchAssyParams,
) (*techspec.ClutchAssyRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreateClutchAssy(ctx, sqlc.CreateClutchAssyParams{
		ID:                uID,
		PartNo:            p.PartNo,
		PulleyRibs:        sqlnull.String(p.PulleyRibs),
		PulleySize:        sqlnull.String(p.PulleySize),
		CompressorDetails: sqlnull.String(p.CompressorDetails),
		ConnectorType:     sqlnull.String(p.ConnectorType),
		Voltage:           sqlnull.String(p.Voltage),
		ShaftType:         sqlnull.String(p.ShaftType),
		Notes:             sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapClutchAssyRow(row), nil
}

func (t *techSpecQueries) UpdateClutchAssyByPartNo(
	ctx context.Context,
	p techspec.UpdateClutchAssyParams,
) (*techspec.ClutchAssyRow, error) {

	row, err := t.q.UpdateClutchAssyByPartNo(ctx, sqlc.UpdateClutchAssyByPartNoParams{
		PartNo:            p.PartNo,
		PulleyRibs:        sqlnull.String(p.PulleyRibs),
		PulleySize:        sqlnull.String(p.PulleySize),
		CompressorDetails: sqlnull.String(p.CompressorDetails),
		ConnectorType:     sqlnull.String(p.ConnectorType),
		Voltage:           sqlnull.String(p.Voltage),
		ShaftType:         sqlnull.String(p.ShaftType),
		Notes:             sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapClutchAssyRow(row), nil
}

func (t *techSpecQueries) DeleteClutchAssyByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeleteClutchAssyByPartNo(ctx, partNo)
}

func mapClutchAssyRow(r sqlc.ClutchAssy) *techspec.ClutchAssyRow {

	return &techspec.ClutchAssyRow{
		ID:                r.ID.String(),
		PartNo:            r.PartNo,
		PulleyRibs:        sqlnull.StringPtr(r.PulleyRibs),
		PulleySize:        sqlnull.StringPtr(r.PulleySize),
		CompressorDetails: sqlnull.StringPtr(r.CompressorDetails),
		ConnectorType:     sqlnull.StringPtr(r.ConnectorType),
		Voltage:           sqlnull.StringPtr(r.Voltage),
		ShaftType:         sqlnull.StringPtr(r.ShaftType),
		Notes:             sqlnull.StringPtr(r.Notes),
	}
}
