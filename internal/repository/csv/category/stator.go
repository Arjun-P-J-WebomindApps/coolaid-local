package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type StatorRow struct {
	PartNo            string
	Voltage           string
	CompressorDetails string
	Notes             string
}

type StatorRepository interface {
	List(ctx context.Context) ([]StatorRow, error)
}

type statorRepo struct{ q *db.DBContext }

func NewStatorRepository(q *db.DBContext) StatorRepository {
	return &statorRepo{q}
}

func (r *statorRepo) List(ctx context.Context) ([]StatorRow, error) {
	rows, err := r.q.Queries.GetStatorsForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]StatorRow, 0, len(rows))
	for _, v := range rows {
		out = append(out, StatorRow{
			PartNo:            v.PartNo,
			Voltage:           v.Voltage.String,
			CompressorDetails: v.CompressorDetails.String,
			Notes:             v.Notes.String,
		})
	}
	return out, nil
}
