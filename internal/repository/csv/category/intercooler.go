package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type IntercoolerRow struct {
	PartNo     string
	Size       string
	TempSensor string
	Notes      string
}

type IntercoolerRepository interface {
	List(ctx context.Context) ([]IntercoolerRow, error)
}

type intercoolerRepo struct{ q *db.DBContext }

func NewIntercoolerRepository(q *db.DBContext) IntercoolerRepository {
	return &intercoolerRepo{q}
}

func (r *intercoolerRepo) List(ctx context.Context) ([]IntercoolerRow, error) {
	rows, err := r.q.Queries.GetIntercoolersForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]IntercoolerRow, 0, len(rows))
	for _, v := range rows {
		out = append(out, IntercoolerRow{
			PartNo:     v.PartNo,
			Size:       v.Size.String,
			TempSensor: v.TempSensor.String,
			Notes:      v.Notes.String,
		})
	}
	return out, nil
}
