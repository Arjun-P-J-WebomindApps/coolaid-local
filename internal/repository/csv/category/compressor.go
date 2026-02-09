package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type CompressorRow struct {
	PartNo             string
	CompressorID       string
	Oil                string
	Refrigerant        string
	Voltage            string
	PulleyRibs         string
	PulleySize         string
	PipeConnector      string
	CompressorType     string
	CompressorMounting string
	ConnectorType      string
	Notes              string
}

type CompressorRepository interface {
	List(ctx context.Context) ([]CompressorRow, error)
}

type compressorRepo struct{ q *db.DBContext }

func NewCompressorRepository(q *db.DBContext) CompressorRepository {
	return &compressorRepo{q}
}

func (r *compressorRepo) List(ctx context.Context) ([]CompressorRow, error) {
	rows, err := r.q.Queries.GetCompressorsForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]CompressorRow, 0, len(rows))
	for _, v := range rows {
		out = append(out, CompressorRow{
			PartNo:             v.PartNo,
			CompressorID:       v.CompressorID.String,
			Oil:                v.Oil.String,
			Refrigerant:        v.Refrigerant.String,
			Voltage:            v.Voltage.String,
			PulleyRibs:         v.PulleyRibs.String,
			PulleySize:         v.PulleySize.String,
			PipeConnector:      v.PipeConnector.String,
			CompressorType:     v.CompType.String,
			CompressorMounting: v.CompMounting.String,
			ConnectorType:      v.ConnectorType.String,
			Notes:              v.Notes.String,
		})
	}
	return out, nil
}
