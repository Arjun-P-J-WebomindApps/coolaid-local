package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type PressureSwitchRow struct {
	PartNo        string
	ConnectorType string
	ThreadType    string
	Notes         string
}

type PressureSwitchRepository interface {
	List(ctx context.Context) ([]PressureSwitchRow, error)
}

type pressureSwitchRepo struct{ q *db.DBContext }

func NewPressureSwitchRepository(q *db.DBContext) PressureSwitchRepository {
	return &pressureSwitchRepo{q}
}

func (r *pressureSwitchRepo) List(ctx context.Context) ([]PressureSwitchRow, error) {
	rows, err := r.q.Queries.GetPressureSwitchesForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]PressureSwitchRow, 0, len(rows))
	for _, v := range rows {
		out = append(out, PressureSwitchRow{
			PartNo:        v.PartNo,
			ConnectorType: v.ConnectorType.String,
			ThreadType:    v.ThreadType.String,
			Notes:         v.Notes.String,
		})
	}
	return out, nil
}
