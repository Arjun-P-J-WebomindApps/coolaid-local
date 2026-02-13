package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetRadiatorByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.RadiatorRow, error) {

	row, err := t.q.GetRadiatorByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapRadiatorRow(row), nil
}

func (t *techSpecQueries) CreateRadiator(
	ctx context.Context,
	p techspec.CreateRadiatorParams,
) (*techspec.RadiatorRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreateRadiator(ctx, sqlc.CreateRadiatorParams{
		ID:           uID,
		PartNo:       p.PartNo,
		Size:         sqlnull.String(p.Size),
		Transmission: sqlnull.String(p.Transmission),
		TempSensor:   sqlnull.String(p.TempSensor),
		Tank:         sqlnull.String(p.Tank),
		Notes:        sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapRadiatorRow(row), nil
}

func (t *techSpecQueries) UpdateRadiatorByPartNo(
	ctx context.Context,
	p techspec.UpdateRadiatorParams,
) (*techspec.RadiatorRow, error) {

	row, err := t.q.UpdateRadiatorByPartNo(ctx, sqlc.UpdateRadiatorByPartNoParams{
		PartNo:       p.PartNo,
		Size:         sqlnull.String(p.Size),
		Transmission: sqlnull.String(p.Transmission),
		TempSensor:   sqlnull.String(p.TempSensor),
		Tank:         sqlnull.String(p.Tank),
		Notes:        sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapRadiatorRow(row), nil
}

func (t *techSpecQueries) DeleteRadiatorByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeleteRadiatorByPartNo(ctx, partNo)
}

func mapRadiatorRow(r sqlc.Radiator) *techspec.RadiatorRow {

	return &techspec.RadiatorRow{
		ID:           r.ID.String(),
		PartNo:       r.PartNo,
		Size:         sqlnull.StringPtr(r.Size),
		Transmission: sqlnull.StringPtr(r.Transmission),
		TempSensor:   sqlnull.StringPtr(r.TempSensor),
		Tank:         sqlnull.StringPtr(r.Tank),
		Notes:        sqlnull.StringPtr(r.Notes),
	}
}
