package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetFilterDrierByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.FilterDrierRow, error) {

	row, err := t.q.GetFilterDrierByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapFilterDrierRow(row), nil
}

func (t *techSpecQueries) CreateFilterDrier(
	ctx context.Context,
	p techspec.CreateFilterDrierParams,
) (*techspec.FilterDrierRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreateFilterDrier(ctx, sqlc.CreateFilterDrierParams{
		ID:             uID,
		PartNo:         p.PartNo,
		PipeConnector:  sqlnull.String(p.PipeConnector),
		Size:           sqlnull.String(p.Size),
		PressureSwitch: sqlnull.String(p.PressureSwitch),
		Notes:          sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapFilterDrierRow(row), nil
}

func (t *techSpecQueries) UpdateFilterDrierByPartNo(
	ctx context.Context,
	p techspec.UpdateFilterDrierParams,
) (*techspec.FilterDrierRow, error) {

	row, err := t.q.UpdateFilterDrierByPartNo(ctx, sqlc.UpdateFilterDrierByPartNoParams{
		PartNo:         p.PartNo,
		PipeConnector:  sqlnull.String(p.PipeConnector),
		Size:           sqlnull.String(p.Size),
		PressureSwitch: sqlnull.String(p.PressureSwitch),
		Notes:          sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapFilterDrierRow(row), nil
}

func (t *techSpecQueries) DeleteFilterDrierByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeleteFilterDrierByPartNo(ctx, partNo)
}

func mapFilterDrierRow(r sqlc.FilterDrier) *techspec.FilterDrierRow {

	return &techspec.FilterDrierRow{
		ID:             r.ID.String(),
		PartNo:         r.PartNo,
		PipeConnector:  sqlnull.StringPtr(r.PipeConnector),
		Size:           sqlnull.StringPtr(r.Size),
		PressureSwitch: sqlnull.StringPtr(r.PressureSwitch),
		Notes:          sqlnull.StringPtr(r.Notes),
	}
}
