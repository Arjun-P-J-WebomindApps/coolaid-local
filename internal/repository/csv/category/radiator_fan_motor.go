package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type RadiatorFanMotorRow struct {
	PartNo           string
	FanBladeDiameter string
	NumberOfBlades   int32
	Voltage          string
	NumberOfSockets  int32
	ConnectorType    string
	Notes            string
}

type RadiatorFanMotorRepository interface {
	List(ctx context.Context) ([]RadiatorFanMotorRow, error)
}

type radiatorFanMotorRepo struct{ q *db.DBContext }

func NewRadiatorFanMotorRepository(q *db.DBContext) RadiatorFanMotorRepository {
	return &radiatorFanMotorRepo{q}
}

func (r *radiatorFanMotorRepo) List(ctx context.Context) ([]RadiatorFanMotorRow, error) {
	rows, err := r.q.Queries.GetRadFanMotorsForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]RadiatorFanMotorRow, 0, len(rows))
	for _, v := range rows {
		out = append(out, RadiatorFanMotorRow{
			PartNo:           v.PartNo,
			FanBladeDiameter: v.FanBladeDiameter.String,
			NumberOfBlades:   v.NumberOfBlades.Int32,
			Voltage:          v.Voltage.String,
			NumberOfSockets:  v.NumberOfSockets.Int32,
			ConnectorType:    v.ConnectorType.String,
			Notes:            v.Notes.String,
		})
	}
	return out, nil
}
