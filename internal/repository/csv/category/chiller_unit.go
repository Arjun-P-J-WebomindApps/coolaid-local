package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type ChillerUnitRow struct {
	PartNo  string
	Type    string
	Voltage string
	Notes   string
}

type ChillerUnitRepository interface {
	List(ctx context.Context) ([]ChillerUnitRow, error)
}

type chillerUnitRepo struct{ q *db.DBContext }

func NewChillerUnitRepository(q *db.DBContext) ChillerUnitRepository {
	return &chillerUnitRepo{q}
}

func (r *chillerUnitRepo) List(ctx context.Context) ([]ChillerUnitRow, error) {
	rows, err := r.q.Queries.GetChillerUnitsForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]ChillerUnitRow, 0, len(rows))
	for _, v := range rows {
		out = append(out, ChillerUnitRow{
			PartNo:  v.PartNo,
			Type:    v.Type.String,
			Voltage: v.Voltage.String,
			Notes:   v.Notes.String,
		})
	}
	return out, nil
}
