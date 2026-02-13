package techspecrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (t *techSpecQueries) GetCompressorByPartNo(
	ctx context.Context,
	partNo string,
) (*techspec.CompressorRow, error) {

	row, err := t.q.GetCompressorByPartNo(ctx, partNo)
	if err != nil {
		return nil, err
	}

	return mapCompressorRow(row), nil
}

func (t *techSpecQueries) CreateCompressor(
	ctx context.Context,
	p techspec.CreateCompressorParams,
) (*techspec.CompressorRow, error) {

	uID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, techspec.ErrInternal
	}

	row, err := t.q.CreateCompressor(ctx, sqlc.CreateCompressorParams{
		ID:            uID,
		PartNo:        p.PartNo,
		CompressorID:  sqlnull.String(p.CompressorID),
		Oil:           sqlnull.String(p.Oil),
		Refrigerant:   sqlnull.String(p.Refrigerant),
		Voltage:       sqlnull.String(p.Voltage),
		PulleyRibs:    sqlnull.String(p.PulleyRibs),
		PulleySize:    sqlnull.String(p.PulleySize),
		PipeConnector: sqlnull.String(p.PipeConnector),
		CompType:      sqlnull.String(p.CompType),
		CompMounting:  sqlnull.String(p.CompMounting),
		ConnectorType: sqlnull.String(p.ConnectorType),
		Notes:         sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapCompressorRow(row), nil
}

func (t *techSpecQueries) UpdateCompressorByPartNo(
	ctx context.Context,
	p techspec.UpdateCompressorParams,
) (*techspec.CompressorRow, error) {

	row, err := t.q.UpdateCompressorByPartNo(ctx, sqlc.UpdateCompressorByPartNoParams{
		PartNo:        p.PartNo,
		CompressorID:  sqlnull.String(p.CompressorID),
		Oil:           sqlnull.String(p.Oil),
		Refrigerant:   sqlnull.String(p.Refrigerant),
		Voltage:       sqlnull.String(p.Voltage),
		PulleyRibs:    sqlnull.String(p.PulleyRibs),
		PulleySize:    sqlnull.String(p.PulleySize),
		PipeConnector: sqlnull.String(p.PipeConnector),
		CompType:      sqlnull.String(p.CompType),
		CompMounting:  sqlnull.String(p.CompMounting),
		ConnectorType: sqlnull.String(p.ConnectorType),
		Notes:         sqlnull.String(p.Notes),
	})
	if err != nil {
		return nil, err
	}

	return mapCompressorRow(row), nil
}

func (t *techSpecQueries) DeleteCompressorByPartNo(
	ctx context.Context,
	partNo string,
) error {

	return t.q.DeleteCompressorByPartNo(ctx, partNo)
}

func mapCompressorRow(r sqlc.Compressor) *techspec.CompressorRow {

	return &techspec.CompressorRow{
		ID:            r.ID.String(),
		PartNo:        r.PartNo,
		CompressorID:  sqlnull.StringPtr(r.CompressorID),
		Oil:           sqlnull.StringPtr(r.Oil),
		Refrigerant:   sqlnull.StringPtr(r.Refrigerant),
		Voltage:       sqlnull.StringPtr(r.Voltage),
		PulleyRibs:    sqlnull.StringPtr(r.PulleyRibs),
		PulleySize:    sqlnull.StringPtr(r.PulleySize),
		PipeConnector: sqlnull.StringPtr(r.PipeConnector),
		CompType:      sqlnull.StringPtr(r.CompType),
		CompMounting:  sqlnull.StringPtr(r.CompMounting),
		ConnectorType: sqlnull.StringPtr(r.ConnectorType),
		Notes:         sqlnull.StringPtr(r.Notes),
	}
}
