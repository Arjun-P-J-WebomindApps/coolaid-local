package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type CondenserFanAssyRow struct {
	PartNo           string
	Voltage          string
	MotorType        string
	Resistance       string
	FanBladeDiameter string
	NumberOfBlades   int32
	Shroud           string
	ConnectorType    string
	Notes            string
}

type CondenserFanAssyRepository interface {
	List(ctx context.Context) ([]CondenserFanAssyRow, error)
}

type condenserFanAssyRepo struct{ q *db.DBContext }

func NewCondenserFanAssyRepository(q *db.DBContext) CondenserFanAssyRepository {
	return &condenserFanAssyRepo{q}
}

func (r *condenserFanAssyRepo) List(ctx context.Context) ([]CondenserFanAssyRow, error) {
	rows, err := r.q.Queries.GetCondFanAssysForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]CondenserFanAssyRow, 0, len(rows))
	for _, v := range rows {
		out = append(out, CondenserFanAssyRow{
			PartNo:           v.PartNo,
			Voltage:          v.Voltage.String,
			MotorType:        v.MotorType.String,
			Resistance:       v.Resistance.String,
			FanBladeDiameter: v.FanBladeDiameter.String,
			NumberOfBlades:   v.NumberOfBlades.Int32,
			Shroud:           v.Shroud.String,
			ConnectorType:    v.ConnectorType.String,
			Notes:            v.Notes.String,
		})
	}
	return out, nil
}
