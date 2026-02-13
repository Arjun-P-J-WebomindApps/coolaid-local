package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetHeaterCoreByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.HeaterCoreRow, error) {

	row, err := t.q.GetHeaterCoreByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapHeaterCoreRow(row), nil
}

func (t *techSpecQueries) CreateHeaterCore(
	ctx context.Context,
	p techspec.CreateHeaterCoreParams,
) (*techspec.HeaterCoreRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreateHeaterCore(ctx, sqlc.CreateHeaterCoreParams{
		ID:     uID,
		PartNo: p.PartNo,
		Size:   sqlnull.String(p.Size),
		Pipe:   sqlnull.String(p.Pipe),
		Type:   sqlnull.String(p.Type),
		Notes:  sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapHeaterCoreRow(row), nil
}

func (t *techSpecQueries) UpdateHeaterCoreByPartNo(
	ctx context.Context,
	p techspec.UpdateHeaterCoreParams,
) (*techspec.HeaterCoreRow, error) {

	row, err := t.q.UpdateHeaterCoreByPartNo(ctx, sqlc.UpdateHeaterCoreByPartNoParams{
		PartNo: p.PartNo,
		Size:   sqlnull.String(p.Size),
		Pipe:   sqlnull.String(p.Pipe),
		Type:   sqlnull.String(p.Type),
		Notes:  sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapHeaterCoreRow(row), nil
}

func (t *techSpecQueries) DeleteHeaterCoreByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeleteHeaterCoreByPartNo(ctx, partNo)
}

func mapHeaterCoreRow(r sqlc.HeaterCore) *techspec.HeaterCoreRow {

	return &techspec.HeaterCoreRow{
		ID:     r.ID.String(),
		PartNo: r.PartNo,
		Size:   sqlnull.StringPtr(r.Size),
		Pipe:   sqlnull.StringPtr(r.Pipe),
		Type:   sqlnull.StringPtr(r.Type),
		Notes:  sqlnull.StringPtr(r.Notes),
	}
}
