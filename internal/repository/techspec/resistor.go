package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetResistorByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.ResistorRow, error) {

	row, err := t.q.GetResistorByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapResistorRow(row), nil
}

func (t *techSpecQueries) CreateResistor(
	ctx context.Context,
	p techspec.CreateResistorParams,
) (*techspec.ResistorRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreateResistor(ctx, sqlc.CreateResistorParams{
		ID:            uID,
		PartNo:        p.PartNo,
		Type:          sqlnull.String(p.Type),
		ConnectorType: sqlnull.String(p.ConnectorType),
		Voltage:       sqlnull.String(p.Voltage),
		Notes:         sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapResistorRow(row), nil
}

func (t *techSpecQueries) UpdateResistorByPartNo(
	ctx context.Context,
	p techspec.UpdateResistorParams,
) (*techspec.ResistorRow, error) {

	row, err := t.q.UpdateResistorByPartNo(ctx, sqlc.UpdateResistorByPartNoParams{
		PartNo:        p.PartNo,
		Type:          sqlnull.String(p.Type),
		ConnectorType: sqlnull.String(p.ConnectorType),
		Voltage:       sqlnull.String(p.Voltage),
		Notes:         sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapResistorRow(row), nil
}

func (t *techSpecQueries) DeleteResistorByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeleteResistorByPartNo(ctx, partNo)
}

func mapResistorRow(r sqlc.Resistor) *techspec.ResistorRow {

	return &techspec.ResistorRow{
		ID:            r.ID.String(),
		PartNo:        r.PartNo,
		Type:          sqlnull.StringPtr(r.Type),
		ConnectorType: sqlnull.StringPtr(r.ConnectorType),
		Voltage:       sqlnull.StringPtr(r.Voltage),
		Notes:         sqlnull.StringPtr(r.Notes),
	}
}
