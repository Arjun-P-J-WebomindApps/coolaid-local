package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type ResistorRow struct {
	PartNo        string
	Type          string
	ConnectorType string
	Voltage       string
	Notes         string
}

type ResistorRepository interface {
	List(ctx context.Context) ([]ResistorRow, error)
}

type resistorRepo struct{ q *db.DBContext }

func NewResistorRepository(q *db.DBContext) ResistorRepository {
	return &resistorRepo{q}
}

func (r *resistorRepo) List(ctx context.Context) ([]ResistorRow, error) {
	rows, err := r.q.Queries.GetResistorsForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]ResistorRow, 0, len(rows))
	for _, v := range rows {
		out = append(out, ResistorRow{
			PartNo:        v.PartNo,
			Type:          v.Type.String,
			ConnectorType: v.ConnectorType.String,
			Voltage:       v.Voltage.String,
			Notes:         v.Notes.String,
		})
	}
	return out, nil
}
