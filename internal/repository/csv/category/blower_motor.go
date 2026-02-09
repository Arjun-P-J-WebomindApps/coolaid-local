package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type BlowerMotorRow struct {
	PartNo        string
	Mounting      string
	ConnectorType string
	Impeller      string
	Resistance    string
	MotorMounting string
	MotorType     string
	Voltage       string
	Notes         string
}

type BlowerMotorRepository interface {
	List(ctx context.Context) ([]BlowerMotorRow, error)
}

type blowerMotorRepo struct{ q *db.DBContext }

func NewBlowerMotorRepository(q *db.DBContext) BlowerMotorRepository {
	return &blowerMotorRepo{q}
}

func (r *blowerMotorRepo) List(ctx context.Context) ([]BlowerMotorRow, error) {
	rows, err := r.q.Queries.GetBlowerMotorsForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]BlowerMotorRow, 0, len(rows))
	for _, v := range rows {
		out = append(out, BlowerMotorRow{
			PartNo:        v.PartNo,
			Mounting:      v.Mounting.String,
			ConnectorType: v.ConnectorType.String,
			Impeller:      v.Impeller.String,
			Resistance:    v.Resistance.String,
			MotorMounting: v.MotorMounting.String,
			MotorType:     v.MotorType.String,
			Voltage:       v.Voltage.String,
			Notes:         v.Notes.String,
		})
	}
	return out, nil
}
