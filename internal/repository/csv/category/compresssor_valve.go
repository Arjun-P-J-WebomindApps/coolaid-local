package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type CompressorValveRow struct {
	PartNo            string
	Type              string
	Voltage           string
	ConnectorType     string
	CompressorDetails string
	Notes             string
}

type CompressorValveRepository interface {
	List(ctx context.Context) ([]CompressorValveRow, error)
}

type compressorValveRepo struct{ q *db.DBContext }

func NewCompressorValveRepository(q *db.DBContext) CompressorValveRepository {
	return &compressorValveRepo{q}
}

func (r *compressorValveRepo) List(ctx context.Context) ([]CompressorValveRow, error) {
	rows, err := r.q.Queries.GetCompressorValvesForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]CompressorValveRow, 0, len(rows))
	for _, v := range rows {
		out = append(out, CompressorValveRow{
			PartNo:            v.PartNo,
			Type:              v.Type.String,
			Voltage:           v.Voltage.String,
			ConnectorType:     v.ConnectorType.String,
			CompressorDetails: v.CompressorDetails.String,
			Notes:             v.Notes.String,
		})
	}
	return out, nil
}
