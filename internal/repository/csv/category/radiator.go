package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type RadiatorRow struct {
	PartNo       string
	Size         string
	Transmission string
	TempSensor   string
	Tank         string
	Notes        string
}

type RadiatorRepository interface {
	List(ctx context.Context) ([]RadiatorRow, error)
}

type radiatorRepo struct{ q *db.DBContext }

func NewRadiatorRepository(q *db.DBContext) RadiatorRepository {
	return &radiatorRepo{q}
}

func (r *radiatorRepo) List(ctx context.Context) ([]RadiatorRow, error) {
	rows, err := r.q.Queries.GetRadiatorsForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]RadiatorRow, 0, len(rows))
	for _, v := range rows {
		out = append(out, RadiatorRow{
			PartNo:       v.PartNo,
			Size:         v.Size.String,
			Transmission: v.Transmission.String,
			TempSensor:   v.TempSensor.String,
			Tank:         v.Tank.String,
			Notes:        v.Notes.String,
		})
	}
	return out, nil
}
