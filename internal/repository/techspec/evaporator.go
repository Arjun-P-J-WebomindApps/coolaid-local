package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetEvaporatorByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.EvaporatorRow, error) {

	row, err := t.q.GetEvaporatorByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapEvaporatorRow(row), nil
}

func (t *techSpecQueries) CreateEvaporator(
	ctx context.Context,
	p techspec.CreateEvaporatorParams,
) (*techspec.EvaporatorRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreateEvaporator(ctx, sqlc.CreateEvaporatorParams{
		ID:             uID,
		PartNo:         p.PartNo,
		Mounting:       sqlnull.String(p.Mounting),
		ExpValve:       sqlnull.String(p.ExpValve),
		AdditionalInfo: sqlnull.String(p.AdditionalInfo),
		Dimensions:     sqlnull.String(p.Dimensions),
		PipeConnector:  sqlnull.String(p.PipeConnector),
		Notes:          sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapEvaporatorRow(row), nil
}

func (t *techSpecQueries) UpdateEvaporatorByPartNo(
	ctx context.Context,
	p techspec.UpdateEvaporatorParams,
) (*techspec.EvaporatorRow, error) {

	row, err := t.q.UpdateEvaporatorByPartNo(ctx, sqlc.UpdateEvaporatorByPartNoParams{
		PartNo:         p.PartNo,
		Mounting:       sqlnull.String(p.Mounting),
		ExpValve:       sqlnull.String(p.ExpValve),
		AdditionalInfo: sqlnull.String(p.AdditionalInfo),
		Dimensions:     sqlnull.String(p.Dimensions),
		PipeConnector:  sqlnull.String(p.PipeConnector),
		Notes:          sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapEvaporatorRow(row), nil
}

func (t *techSpecQueries) DeleteEvaporatorByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeleteEvaporatorByPartNo(ctx, partNo)
}

func mapEvaporatorRow(r sqlc.Evaporator) *techspec.EvaporatorRow {

	return &techspec.EvaporatorRow{
		ID:             r.ID.String(),
		PartNo:         r.PartNo,
		Mounting:       sqlnull.StringPtr(r.Mounting),
		ExpValve:       sqlnull.StringPtr(r.ExpValve),
		AdditionalInfo: sqlnull.StringPtr(r.AdditionalInfo),
		Dimensions:     sqlnull.StringPtr(r.Dimensions),
		PipeConnector:  sqlnull.StringPtr(r.PipeConnector),
		Notes:          sqlnull.StringPtr(r.Notes),
	}
}
