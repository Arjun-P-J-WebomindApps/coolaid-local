package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type EvaporatorRow struct {
	PartNo         string
	Mounting       string
	ExpansionValve string
	AdditionalInfo string
	Dimensions     string
	PipeConnector  string
	Notes          string
}

type EvaporatorRepository interface {
	List(ctx context.Context) ([]EvaporatorRow, error)
}

type evaporatorRepo struct{ q *db.DBContext }

func NewEvaporatorRepository(q *db.DBContext) EvaporatorRepository {
	return &evaporatorRepo{q}
}

func (r *evaporatorRepo) List(ctx context.Context) ([]EvaporatorRow, error) {
	rows, err := r.q.Queries.GetEvaporatorsForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]EvaporatorRow, 0, len(rows))
	for _, v := range rows {
		out = append(out, EvaporatorRow{
			PartNo:         v.PartNo,
			Mounting:       v.Mounting.String,
			ExpansionValve: v.ExpValve.String,
			AdditionalInfo: v.AdditionalInfo.String,
			Dimensions:     v.Dimensions.String,
			PipeConnector:  v.PipeConnector.String,
			Notes:          v.Notes.String,
		})
	}
	return out, nil
}
