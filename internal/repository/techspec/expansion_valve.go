package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetExpansionValveByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.ExpansionValveRow, error) {

	row, err := t.q.GetExpansionValveByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapExpansionValveRow(row), nil
}

func (t *techSpecQueries) CreateExpansionValve(
	ctx context.Context,
	p techspec.CreateExpansionValveParams,
) (*techspec.ExpansionValveRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreateExpansionValve(ctx, sqlc.CreateExpansionValveParams{
		ID:          uID,
		PartNo:      p.PartNo,
		Type:        sqlnull.String(p.Type),
		Material:    sqlnull.String(p.Material),
		Refrigerant: sqlnull.String(p.Refrigerant),
		Notes:       sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapExpansionValveRow(row), nil
}

func (t *techSpecQueries) UpdateExpansionValveByPartNo(
	ctx context.Context,
	p techspec.UpdateExpansionValveParams,
) (*techspec.ExpansionValveRow, error) {

	row, err := t.q.UpdateExpansionValveByPartNo(ctx, sqlc.UpdateExpansionValveByPartNoParams{
		PartNo:      p.PartNo,
		Type:        sqlnull.String(p.Type),
		Material:    sqlnull.String(p.Material),
		Refrigerant: sqlnull.String(p.Refrigerant),
		Notes:       sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapExpansionValveRow(row), nil
}

func (t *techSpecQueries) DeleteExpansionValveByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeleteExpansionValveByPartNo(ctx, partNo)
}

func mapExpansionValveRow(r sqlc.ExpansionValf) *techspec.ExpansionValveRow {

	return &techspec.ExpansionValveRow{
		ID:          r.ID.String(),
		PartNo:      r.PartNo,
		Type:        sqlnull.StringPtr(r.Type),
		Material:    sqlnull.StringPtr(r.Material),
		Refrigerant: sqlnull.StringPtr(r.Refrigerant),
		Notes:       sqlnull.StringPtr(r.Notes),
	}
}
