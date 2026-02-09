package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type HeaterRow struct {
	PartNo string
	Size   string
	Pipe   string
	Type   string
	Notes  string
}

type HeaterRepository interface {
	List(ctx context.Context) ([]HeaterRow, error)
}

type heaterRepo struct{ q *db.DBContext }

func NewHeaterRepository(q *db.DBContext) HeaterRepository {
	return &heaterRepo{q}
}

func (r *heaterRepo) List(ctx context.Context) ([]HeaterRow, error) {
	rows, err := r.q.Queries.GetHeaterCoresForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]HeaterRow, 0, len(rows))
	for _, v := range rows {
		out = append(out, HeaterRow{
			PartNo: v.PartNo,
			Size:   v.Size.String,
			Pipe:   v.Pipe.String,
			Type:   v.Type.String,
			Notes:  v.Notes.String,
		})
	}
	return out, nil
}
