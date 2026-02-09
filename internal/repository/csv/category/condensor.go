package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type CondenserRow struct {
	PartNo         string
	Size           string
	PipeConnector  string
	Drier          string
	PressureSwitch string
	Notes          string
}

type CondenserRepository interface {
	List(ctx context.Context) ([]CondenserRow, error)
}

type condenserRepo struct{ q *db.DBContext }

func NewCondenserRepository(q *db.DBContext) CondenserRepository {
	return &condenserRepo{q}
}

func (r *condenserRepo) List(ctx context.Context) ([]CondenserRow, error) {
	rows, err := r.q.Queries.GetCondensersForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]CondenserRow, 0, len(rows))
	for _, v := range rows {
		out = append(out, CondenserRow{
			PartNo:         v.PartNo,
			Size:           v.Size.String,
			PipeConnector:  v.PipeConnector.String,
			Drier:          v.Drier.String,
			PressureSwitch: v.PressureSwitch.String,
			Notes:          v.Notes.String,
		})
	}
	return out, nil
}
