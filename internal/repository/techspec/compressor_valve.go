package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetCompressorValveByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.CompressorValveRow, error) {

	row, err := t.q.GetCompressorValveByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapCompressorValveRow(row), nil
}

func (t *techSpecQueries) CreateCompressorValve(
	ctx context.Context,
	p techspec.CreateCompressorValveParams,
) (*techspec.CompressorValveRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreateCompressorValve(ctx, sqlc.CreateCompressorValveParams{
		ID:                uID,
		PartNo:            p.PartNo,
		Type:              sqlnull.String(p.Type),
		Voltage:           sqlnull.String(p.Voltage),
		ConnectorType:     sqlnull.String(p.ConnectorType),
		CompressorDetails: sqlnull.String(p.CompressorDetails),
		Notes:             sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapCompressorValveRow(row), nil
}
func (t *techSpecQueries) UpdateCompressorValveByPartNo(
	ctx context.Context,
	p techspec.UpdateCompressorValveParams,
) (*techspec.CompressorValveRow, error) {

	row, err := t.q.UpdateCompressorValveByPartNo(ctx, sqlc.UpdateCompressorValveByPartNoParams{
		PartNo:            p.PartNo,
		Type:              sqlnull.String(p.Type),
		Voltage:           sqlnull.String(p.Voltage),
		ConnectorType:     sqlnull.String(p.ConnectorType),
		CompressorDetails: sqlnull.String(p.CompressorDetails),
		Notes:             sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapCompressorValveRow(row), nil
}

func (t *techSpecQueries) DeleteCompressorValveByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeleteCompressorValveByPartNo(ctx, partNo)
}

func mapCompressorValveRow(r sqlc.CompressorValf) *techspec.CompressorValveRow {

	return &techspec.CompressorValveRow{
		ID:                r.ID.String(),
		PartNo:            r.PartNo,
		Type:              sqlnull.StringPtr(r.Type),
		Voltage:           sqlnull.StringPtr(r.Voltage),
		ConnectorType:     sqlnull.StringPtr(r.ConnectorType),
		CompressorDetails: sqlnull.StringPtr(r.CompressorDetails),
		Notes:             sqlnull.StringPtr(r.Notes),
	}
}
