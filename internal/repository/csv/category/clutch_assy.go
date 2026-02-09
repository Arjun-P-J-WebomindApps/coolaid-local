package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type ClutchAssyRow struct {
	PartNo            string
	PulleyRibs        string
	PulleySize        string
	CompressorDetails string
	ConnectorType     string
	Voltage           string
	ShaftType         string
	Notes             string
}

type ClutchAssyRepository interface {
	List(ctx context.Context) ([]ClutchAssyRow, error)
}

type clutchAssyRepo struct{ q *db.DBContext }

func NewClutchAssyRepository(q *db.DBContext) ClutchAssyRepository {
	return &clutchAssyRepo{q}
}

func (r *clutchAssyRepo) List(ctx context.Context) ([]ClutchAssyRow, error) {
	rows, err := r.q.Queries.GetClutchAssysForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]ClutchAssyRow, 0, len(rows))
	for _, v := range rows {
		out = append(out, ClutchAssyRow{
			PartNo:            v.PartNo,
			PulleyRibs:        v.PulleyRibs.String,
			PulleySize:        v.PulleySize.String,
			CompressorDetails: v.CompressorDetails.String,
			ConnectorType:     v.ConnectorType.String,
			Voltage:           v.Voltage.String,
			ShaftType:         v.ShaftType.String,
			Notes:             v.Notes.String,
		})
	}
	return out, nil
}
