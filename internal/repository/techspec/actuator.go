package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetActuatorByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.ActuatorRow, error) {

	row, err := t.q.GetActuatorByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapActuatorRow(row), nil
}

func (t *techSpecQueries) CreateActuator(
	ctx context.Context,
	p techspec.CreateActuatorParams,
) (*techspec.ActuatorRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreateActuator(ctx, sqlc.CreateActuatorParams{
		ID:            uID,
		PartNo:        p.PartNo,
		ConnectorType: sqlnull.String(p.ConnectorType),
		Mounting:      sqlnull.String(p.Mounting),
		Voltage:       sqlnull.String(p.Voltage),
		RotationAngle: sqlnull.String(p.RotationAngle),
		Notes:         sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapActuatorRow(row), nil
}

func (t *techSpecQueries) UpdateActuatorByPartNo(
	ctx context.Context,
	p techspec.UpdateActuatorParams,
) (*techspec.ActuatorRow, error) {

	row, err := t.q.UpdateActuatorByPartNo(ctx, sqlc.UpdateActuatorByPartNoParams{
		PartNo:        p.PartNo,
		ConnectorType: sqlnull.String(p.ConnectorType),
		Mounting:      sqlnull.String(p.Mounting),
		Voltage:       sqlnull.String(p.Voltage),
		RotationAngle: sqlnull.String(p.RotationAngle),
		Notes:         sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapActuatorRow(row), nil
}

func (t *techSpecQueries) DeleteActuatorByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeleteActuatorByPartNo(ctx, partNo)
}

func mapActuatorRow(r sqlc.Actuator) *techspec.ActuatorRow {

	return &techspec.ActuatorRow{
		ID:            r.ID.String(),
		PartNo:        r.PartNo,
		ConnectorType: sqlnull.StringPtr(r.ConnectorType),
		Mounting:      sqlnull.StringPtr(r.Mounting),
		Voltage:       sqlnull.StringPtr(r.Voltage),
		RotationAngle: sqlnull.StringPtr(r.RotationAngle),
		Notes:         sqlnull.StringPtr(r.Notes),
	}
}
