package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetCondFanAssyByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.CondFanAssyRow, error) {

	row, err := t.q.GetCondFanAssyByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapCondFanAssyRow(row), nil
}

func (t *techSpecQueries) CreateCondFanAssy(
	ctx context.Context,
	p techspec.CreateCondFanAssyParams,
) (*techspec.CondFanAssyRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreateCondFanAssy(ctx, sqlc.CreateCondFanAssyParams{
		ID:               uID,
		PartNo:           p.PartNo,
		Voltage:          sqlnull.String(p.Voltage),
		MotorType:        sqlnull.String(p.MotorType),
		Resistance:       sqlnull.String(p.Resistance),
		FanBladeDiameter: sqlnull.String(p.FanBladeDiameter),
		NumberOfBlades:   sqlnull.Int32(p.NumberOfBlades),
		Shroud:           sqlnull.String(p.Shroud),
		ConnectorType:    sqlnull.String(p.ConnectorType),
		Notes:            sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapCondFanAssyRow(row), nil
}

func (t *techSpecQueries) UpdateCondFanAssyByPartNo(
	ctx context.Context,
	p techspec.UpdateCondFanAssyParams,
) (*techspec.CondFanAssyRow, error) {

	row, err := t.q.UpdateCondFanAssyByPartNo(ctx, sqlc.UpdateCondFanAssyByPartNoParams{
		PartNo:           p.PartNo,
		Voltage:          sqlnull.String(p.Voltage),
		MotorType:        sqlnull.String(p.MotorType),
		Resistance:       sqlnull.String(p.Resistance),
		FanBladeDiameter: sqlnull.String(p.FanBladeDiameter),
		NumberOfBlades:   sqlnull.Int32(p.NumberOfBlades),
		Shroud:           sqlnull.String(p.Shroud),
		ConnectorType:    sqlnull.String(p.ConnectorType),
		Notes:            sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapCondFanAssyRow(row), nil
}

func (t *techSpecQueries) DeleteCondFanAssyByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeleteCondFanAssyByPartNo(ctx, partNo)
}

func mapCondFanAssyRow(r sqlc.CondFanAssy) *techspec.CondFanAssyRow {

	return &techspec.CondFanAssyRow{
		ID:               r.ID.String(),
		PartNo:           r.PartNo,
		Voltage:          sqlnull.StringPtr(r.Voltage),
		MotorType:        sqlnull.StringPtr(r.MotorType),
		Resistance:       sqlnull.StringPtr(r.Resistance),
		FanBladeDiameter: sqlnull.StringPtr(r.FanBladeDiameter),
		NumberOfBlades:   sqlnull.Int32Ptr(r.NumberOfBlades),
		Shroud:           sqlnull.StringPtr(r.Shroud),
		ConnectorType:    sqlnull.StringPtr(r.ConnectorType),
		Notes:            sqlnull.StringPtr(r.Notes),
	}
}
