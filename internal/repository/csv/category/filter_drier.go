package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type FilterDrierRow struct {
	PartNo         string
	PipeConnector  string
	Size           string
	PressureSwitch string
	Notes          string
}

type FilterDrierRepository interface {
	List(ctx context.Context) ([]FilterDrierRow, error)
}

type filterDrierRepo struct{ q *db.DBContext }

func NewFilterDrierRepository(q *db.DBContext) FilterDrierRepository {
	return &filterDrierRepo{q}
}

func (r *filterDrierRepo) List(ctx context.Context) ([]FilterDrierRow, error) {
	rows, err := r.q.Queries.GetFilterDriersForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]FilterDrierRow, 0, len(rows))
	for _, v := range rows {
		out = append(out, FilterDrierRow{
			PartNo:         v.PartNo,
			PipeConnector:  v.PipeConnector.String,
			Size:           v.Size.String,
			PressureSwitch: v.PressureSwitch.String,
			Notes:          v.Notes.String,
		})
	}
	return out, nil
}
