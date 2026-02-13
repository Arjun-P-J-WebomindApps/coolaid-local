package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetRadFanMotorByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.RadFanMotorRow, error) {

	row, err := t.q.GetRadFanMotorByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapRadFanMotorRow(row), nil
}

func (t *techSpecQueries) CreateRadFanMotor(
	ctx context.Context,
	p techspec.CreateRadFanMotorParams,
) (*techspec.RadFanMotorRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreateRadFanMotor(ctx, sqlc.CreateRadFanMotorParams{
		ID:               uID,
		PartNo:           p.PartNo,
		FanBladeDiameter: sqlnull.String(p.FanBladeDiameter),
		NumberOfBlades:   sqlnull.Int32(p.NumberOfBlades),
		Voltage:          sqlnull.String(p.Voltage),
		NumberOfSockets:  sqlnull.Int32(p.NumberOfSockets),
		ConnectorType:    sqlnull.String(p.ConnectorType),
		Notes:            sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapRadFanMotorRow(row), nil
}

func (t *techSpecQueries) UpdateRadFanMotorByPartNo(
	ctx context.Context,
	p techspec.UpdateRadFanMotorParams,
) (*techspec.RadFanMotorRow, error) {

	row, err := t.q.UpdateRadFanMotorByPartNo(ctx, sqlc.UpdateRadFanMotorByPartNoParams{
		PartNo:           p.PartNo,
		FanBladeDiameter: sqlnull.String(p.FanBladeDiameter),
		NumberOfBlades:   sqlnull.Int32(p.NumberOfBlades),
		Voltage:          sqlnull.String(p.Voltage),
		NumberOfSockets:  sqlnull.Int32(p.NumberOfSockets),
		ConnectorType:    sqlnull.String(p.ConnectorType),
		Notes:            sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapRadFanMotorRow(row), nil
}

func (t *techSpecQueries) DeleteRadFanMotorByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeleteRadFanMotorByPartNo(ctx, partNo)
}

func mapRadFanMotorRow(r sqlc.RadFanMotor) *techspec.RadFanMotorRow {

	return &techspec.RadFanMotorRow{
		ID:               r.ID.String(),
		PartNo:           r.PartNo,
		FanBladeDiameter: sqlnull.StringPtr(r.FanBladeDiameter),
		NumberOfBlades:   sqlnull.Int32Ptr(r.NumberOfBlades),
		Voltage:          sqlnull.StringPtr(r.Voltage),
		NumberOfSockets:  sqlnull.Int32Ptr(r.NumberOfSockets),
		ConnectorType:    sqlnull.StringPtr(r.ConnectorType),
		Notes:            sqlnull.StringPtr(r.Notes),
	}
}
