package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetBlowerMotorByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.BlowerMotorRow, error) {

	row, err := t.q.GetBlowerMotorByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapBlowerMotorRow(row), nil
}

func (t *techSpecQueries) CreateBlowerMotor(
	ctx context.Context,
	p techspec.CreateBlowerMotorParams,
) (*techspec.BlowerMotorRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreateBlowerMotor(ctx, sqlc.CreateBlowerMotorParams{
		ID:            uID,
		PartNo:        p.PartNo,
		Mounting:      sqlnull.String(p.Mounting),
		ConnectorType: sqlnull.String(p.ConnectorType),
		Impeller:      sqlnull.String(p.Impeller),
		Resistance:    sqlnull.String(p.Resistance),
		MotorMounting: sqlnull.String(p.MotorMounting),
		MotorType:     sqlnull.String(p.MotorType),
		Voltage:       sqlnull.String(p.Voltage),
		Notes:         sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapBlowerMotorRow(row), nil
}

func (t *techSpecQueries) UpdateBlowerMotorByPartNo(
	ctx context.Context,
	p techspec.UpdateBlowerMotorParams,
) (*techspec.BlowerMotorRow, error) {

	row, err := t.q.UpdateBlowerMotorByPartNo(ctx, sqlc.UpdateBlowerMotorByPartNoParams{
		PartNo:        p.PartNo,
		Mounting:      sqlnull.String(p.Mounting),
		ConnectorType: sqlnull.String(p.ConnectorType),
		Impeller:      sqlnull.String(p.Impeller),
		Resistance:    sqlnull.String(p.Resistance),
		MotorMounting: sqlnull.String(p.MotorMounting),
		MotorType:     sqlnull.String(p.MotorType),
		Voltage:       sqlnull.String(p.Voltage),
		Notes:         sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapBlowerMotorRow(row), nil
}

func (t *techSpecQueries) DeleteBlowerMotorByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeleteBlowerMotorByPartNo(ctx, partNo)
}

func mapBlowerMotorRow(r sqlc.BlowerMotor) *techspec.BlowerMotorRow {

	return &techspec.BlowerMotorRow{
		ID:            r.ID.String(),
		PartNo:        r.PartNo,
		Mounting:      sqlnull.StringPtr(r.Mounting),
		ConnectorType: sqlnull.StringPtr(r.ConnectorType),
		Impeller:      sqlnull.StringPtr(r.Impeller),
		Resistance:    sqlnull.StringPtr(r.Resistance),
		MotorMounting: sqlnull.StringPtr(r.MotorMounting),
		MotorType:     sqlnull.StringPtr(r.MotorType),
		Voltage:       sqlnull.StringPtr(r.Voltage),
		Notes:         sqlnull.StringPtr(r.Notes),
	}
}
