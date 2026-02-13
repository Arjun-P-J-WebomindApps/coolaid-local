package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetPressureSwitchByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.PressureSwitchRow, error) {

	row, err := t.q.GetPressureSwitchByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapPressureSwitchRow(row), nil
}

func (t *techSpecQueries) CreatePressureSwitch(
	ctx context.Context,
	p techspec.CreatePressureSwitchParams,
) (*techspec.PressureSwitchRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreatePressureSwitch(ctx, sqlc.CreatePressureSwitchParams{
		ID:            uID,
		PartNo:        p.PartNo,
		ConnectorType: sqlnull.String(p.ConnectorType),
		ThreadType:    sqlnull.String(p.ThreadType),
		Notes:         sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapPressureSwitchRow(row), nil
}

func (t *techSpecQueries) UpdatePressureSwitchByPartNo(
	ctx context.Context,
	p techspec.UpdatePressureSwitchParams,
) (*techspec.PressureSwitchRow, error) {

	row, err := t.q.UpdatePressureSwitchByPartNo(ctx, sqlc.UpdatePressureSwitchByPartNoParams{
		PartNo:        p.PartNo,
		ConnectorType: sqlnull.String(p.ConnectorType),
		ThreadType:    sqlnull.String(p.ThreadType),
		Notes:         sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapPressureSwitchRow(row), nil
}

func (t *techSpecQueries) DeletePressureSwitchByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeletePressureSwitchByPartNo(ctx, partNo)
}

func mapPressureSwitchRow(r sqlc.PressureSwitch) *techspec.PressureSwitchRow {

	return &techspec.PressureSwitchRow{
		ID:            r.ID.String(),
		PartNo:        r.PartNo,
		ConnectorType: sqlnull.StringPtr(r.ConnectorType),
		ThreadType:    sqlnull.StringPtr(r.ThreadType),
		Notes:         sqlnull.StringPtr(r.Notes),
	}
}
