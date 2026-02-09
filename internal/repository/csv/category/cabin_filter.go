package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type CabinFilterRow struct {
	PartNo     string
	Type       string
	Dimensions string
	Material   string
	Notes      string
}

type CabinFilterRepository interface {
	List(ctx context.Context) ([]CabinFilterRow, error)
}

type cabinFilterRepo struct{ q *db.DBContext }

func NewCabinFilterRepository(q *db.DBContext) CabinFilterRepository {
	return &cabinFilterRepo{q}
}

func (r *cabinFilterRepo) List(ctx context.Context) ([]CabinFilterRow, error) {
	rows, err := r.q.Queries.GetCabinFiltersForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]CabinFilterRow, 0, len(rows))
	for _, v := range rows {
		out = append(out, CabinFilterRow{
			PartNo:     v.PartNo,
			Type:       v.Type.String,
			Dimensions: v.Dimensions.String,
			Material:   v.Material.String,
			Notes:      v.Notes.String,
		})
	}
	return out, nil
}
