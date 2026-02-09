package category

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

/*
ActuatorRow is the strongly typed row returned by the repository.
*/
type ActuatorRow struct {
	PartNo        string
	ConnectorType string
	Mounting      string
	Voltage       string
	RotationAngle string
	Notes         string
}

/*
ActuatorRepository exposes typed access to actuator data.
*/
type ActuatorRepository interface {
	List(ctx context.Context) ([]ActuatorRow, error)
}

type actuatorRepo struct {
	q *db.DBContext
}

func NewActuatorRepository(q *db.DBContext) ActuatorRepository {
	return &actuatorRepo{q: q}
}

func (r *actuatorRepo) List(ctx context.Context) ([]ActuatorRow, error) {
	items, err := r.q.Queries.GetActuatorForDownload(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]ActuatorRow, 0, len(items))
	for _, it := range items {
		out = append(out, ActuatorRow{
			PartNo:        it.PartNo,
			ConnectorType: it.ConnectorType.String,
			Mounting:      it.Mounting.String,
			Voltage:       it.Voltage.String,
			RotationAngle: it.RotationAngle.String,
			Notes:         it.Notes.String,
		})
	}

	return out, nil
}
