package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type RotorRow struct {
	PartNo            string
	PulleyRibs        string
	PulleySize        string
	CompressorDetails string
	Notes             string
}

type RotorRepository interface {
	List(ctx context.Context) ([]RotorRow, error)
}

type rotorRepo struct{ q *db.DBContext }

func NewRotorRepository(q *db.DBContext) RotorRepository {
	return &rotorRepo{q}
}

func (r *rotorRepo) List(ctx context.Context) ([]RotorRow, error) {
	rows, err := r.q.Queries.GetRotorsForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]RotorRow, 0, len(rows))
	for _, v := range rows {
		out = append(out, RotorRow{
			PartNo:            v.PartNo,
			PulleyRibs:        v.PulleyRibs.String,
			PulleySize:        v.PulleySize.String,
			CompressorDetails: v.CompressorDetails.String,
			Notes:             v.Notes.String,
		})
	}
	return out, nil
}
