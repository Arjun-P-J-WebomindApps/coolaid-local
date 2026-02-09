package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

/*
ExpansionValveRow represents a single expansion valve record.
*/
type ExpansionValveRow struct {
	PartNo      string
	Type        string
	Material    string
	Refrigerant string
	Notes       string
}

/*
ExpansionValveRepository exposes typed access to expansion valve data.
*/
type ExpansionValveRepository interface {
	List(ctx context.Context) ([]ExpansionValveRow, error)
}

type expansionValveRepo struct {
	q *db.DBContext
}

func NewExpansionValveRepository(q *db.DBContext) ExpansionValveRepository {
	return &expansionValveRepo{q: q}
}

func (r *expansionValveRepo) List(ctx context.Context) ([]ExpansionValveRow, error) {
	items, err := r.q.Queries.GetExpansionValvesForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]ExpansionValveRow, 0, len(items))
	for _, it := range items {
		out = append(out, ExpansionValveRow{
			PartNo:      it.PartNo,
			Type:        it.Type.String,
			Material:    it.Material.String,
			Refrigerant: it.Refrigerant.String,
			Notes:       it.Notes.String,
		})
	}

	return out, nil
}
