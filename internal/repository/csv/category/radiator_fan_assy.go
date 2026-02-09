package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

/*
RadiatorFanAssyRow represents a radiator fan assembly record.
*/
type RadiatorFanAssyRow struct {
	PartNo           string
	Voltage          string
	MotorType        string
	Resistance       string
	NumberOfSockets  int32
	Shroud           string
	ConnectorType    string
	FanBladeDiameter string
	NumberOfBlades   int32
	Notes            string
}

/*
RadiatorFanAssyRepository exposes typed access to radiator fan assemblies.
*/
type RadiatorFanAssyRepository interface {
	List(ctx context.Context) ([]RadiatorFanAssyRow, error)
}

type radiatorFanAssyRepo struct {
	q *db.DBContext
}

func NewRadiatorFanAssyRepository(q *db.DBContext) RadiatorFanAssyRepository {
	return &radiatorFanAssyRepo{q: q}
}

func (r *radiatorFanAssyRepo) List(ctx context.Context) ([]RadiatorFanAssyRow, error) {
	items, err := r.q.Queries.GetRadFanAssysForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]RadiatorFanAssyRow, 0, len(items))
	for _, it := range items {
		out = append(out, RadiatorFanAssyRow{
			PartNo:           it.PartNo,
			Voltage:          it.Voltage.String,
			MotorType:        it.MotorType.String,
			Resistance:       it.Resistance.String,
			NumberOfSockets:  it.NumberOfSockets.Int32,
			Shroud:           it.Shroud.String,
			ConnectorType:    it.ConnectorType.String,
			FanBladeDiameter: it.FanBladeDiameter.String,
			NumberOfBlades:   it.NumberOfBlades.Int32,
			Notes:            it.Notes.String,
		})
	}

	return out, nil
}
