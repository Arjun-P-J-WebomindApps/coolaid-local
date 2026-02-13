package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetRadFanAssyByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.RadFanAssyRow, error) {

	row, err := t.q.GetRadFanAssyByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapRadFanAssyRow(row), nil
}

func (t *techSpecQueries) CreateRadFanAssy(
	ctx context.Context,
	p techspec.CreateRadFanAssyParams,
) (*techspec.RadFanAssyRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreateRadFanAssy(ctx, sqlc.CreateRadFanAssyParams{
		ID:               uID,
		PartNo:           p.PartNo,
		Voltage:          sqlnull.String(p.Voltage),
		MotorType:        sqlnull.String(p.MotorType),
		Resistance:       sqlnull.String(p.Resistance),
		NumberOfSockets:  sqlnull.Int32(p.NumberOfSockets),
		Shroud:           sqlnull.String(p.Shroud),
		ConnectorType:    sqlnull.String(p.ConnectorType),
		FanBladeDiameter: sqlnull.String(p.FanBladeDiameter),
		NumberOfBlades:   sqlnull.Int32(p.NumberOfBlades),
		Notes:            sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapRadFanAssyRow(row), nil
}

func (t *techSpecQueries) UpdateRadFanAssyByPartNo(
	ctx context.Context,
	p techspec.UpdateRadFanAssyParams,
) (*techspec.RadFanAssyRow, error) {

	row, err := t.q.UpdateRadFanAssyByPartNo(ctx, sqlc.UpdateRadFanAssyByPartNoParams{
		PartNo:           p.PartNo,
		Voltage:          sqlnull.String(p.Voltage),
		MotorType:        sqlnull.String(p.MotorType),
		Resistance:       sqlnull.String(p.Resistance),
		NumberOfSockets:  sqlnull.Int32(p.NumberOfSockets),
		Shroud:           sqlnull.String(p.Shroud),
		ConnectorType:    sqlnull.String(p.ConnectorType),
		FanBladeDiameter: sqlnull.String(p.FanBladeDiameter),
		NumberOfBlades:   sqlnull.Int32(p.NumberOfBlades),
		Notes:            sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapRadFanAssyRow(row), nil
}

func (t *techSpecQueries) DeleteRadFanAssyByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeleteRadFanAssyByPartNo(ctx, partNo)
}

func mapRadFanAssyRow(r sqlc.RadFanAssy) *techspec.RadFanAssyRow {

	return &techspec.RadFanAssyRow{
		ID:               r.ID.String(),
		PartNo:           r.PartNo,
		Voltage:          sqlnull.StringPtr(r.Voltage),
		MotorType:        sqlnull.StringPtr(r.MotorType),
		Resistance:       sqlnull.StringPtr(r.Resistance),
		NumberOfSockets:  sqlnull.Int32Ptr(r.NumberOfSockets),
		Shroud:           sqlnull.StringPtr(r.Shroud),
		ConnectorType:    sqlnull.StringPtr(r.ConnectorType),
		FanBladeDiameter: sqlnull.StringPtr(r.FanBladeDiameter),
		NumberOfBlades:   sqlnull.Int32Ptr(r.NumberOfBlades),
		Notes:            sqlnull.StringPtr(r.Notes),
	}
}
