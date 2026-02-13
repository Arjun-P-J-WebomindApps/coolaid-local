package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetIntercoolerByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.IntercoolerRow, error) {

	row, err := t.q.GetIntercoolerByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapIntercoolerRow(row), nil
}

func (t *techSpecQueries) CreateIntercooler(
	ctx context.Context,
	p techspec.CreateIntercoolerParams,
) (*techspec.IntercoolerRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreateIntercooler(ctx, sqlc.CreateIntercoolerParams{
		ID:         uID,
		PartNo:     p.PartNo,
		Size:       sqlnull.String(p.Size),
		TempSensor: sqlnull.String(p.TempSensor),
		Notes:      sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapIntercoolerRow(row), nil
}

func (t *techSpecQueries) UpdateIntercoolerByPartNo(
	ctx context.Context,
	p techspec.UpdateIntercoolerParams,
) (*techspec.IntercoolerRow, error) {

	row, err := t.q.UpdateIntercoolerByPartNo(ctx, sqlc.UpdateIntercoolerByPartNoParams{
		PartNo:     p.PartNo,
		Size:       sqlnull.String(p.Size),
		TempSensor: sqlnull.String(p.TempSensor),
		Notes:      sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapIntercoolerRow(row), nil
}

func (t *techSpecQueries) DeleteIntercoolerByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeleteIntercoolerByPartNo(ctx, partNo)
}

func mapIntercoolerRow(r sqlc.Intercooler) *techspec.IntercoolerRow {

	return &techspec.IntercoolerRow{
		ID:         r.ID.String(),
		PartNo:     r.PartNo,
		Size:       sqlnull.StringPtr(r.Size),
		TempSensor: sqlnull.StringPtr(r.TempSensor),
		Notes:      sqlnull.StringPtr(r.Notes),
	}
}
