// internal/repository/category/registry.go
package category

import (
	"context"
	"fmt"

	"github.com/webomindapps-dev/coolaid-backend/db"
)

type CategoryRepository interface {
	List(ctx context.Context) ([]Row, error)
}

func RepositoryFor(category string, q *db.DBContext) (CategoryRepository, error) {
	switch category {

	case "ACTUATOR":
		r := NewActuatorRepository(q)
		return adapter[ActuatorRow]{list: r.List}, nil

	case "BLOWER MOTOR":
		r := NewBlowerMotorRepository(q)
		return adapter[BlowerMotorRow]{list: r.List}, nil

	case "CABIN FILTER":
		r := NewCabinFilterRepository(q)
		return adapter[CabinFilterRow]{list: r.List}, nil

	case "CHILLER UNIT":
		r := NewChillerUnitRepository(q)
		return adapter[ChillerUnitRow]{list: r.List}, nil

	case "CLUTCH ASSY":
		r := NewClutchAssyRepository(q)
		return adapter[ClutchAssyRow]{list: r.List}, nil

	case "COMPRESSOR":
		r := NewCompressorRepository(q)
		return adapter[CompressorRow]{list: r.List}, nil

	case "COMPRESSOR VALVE":
		r := NewCompressorValveRepository(q)
		return adapter[CompressorValveRow]{list: r.List}, nil

	case "CONDENSER":
		r := NewCondenserRepository(q)
		return adapter[CondenserRow]{list: r.List}, nil

	case "CONDENSER FAN ASSY":
		r := NewCondenserFanAssyRepository(q)
		return adapter[CondenserFanAssyRow]{list: r.List}, nil

	case "EVAPORATOR":
		r := NewEvaporatorRepository(q)
		return adapter[EvaporatorRow]{list: r.List}, nil

	case "EXPANSION VALVE":
		r := NewExpansionValveRepository(q)
		return adapter[ExpansionValveRow]{list: r.List}, nil

	case "FILTER DRIER":
		r := NewFilterDrierRepository(q)
		return adapter[FilterDrierRow]{list: r.List}, nil

	case "HEATER":
		r := NewHeaterRepository(q)
		return adapter[HeaterRow]{list: r.List}, nil

	case "INTERCOOLER":
		r := NewIntercoolerRepository(q)
		return adapter[IntercoolerRow]{list: r.List}, nil

	case "PRESSURE SWITCH":
		r := NewPressureSwitchRepository(q)
		return adapter[PressureSwitchRow]{list: r.List}, nil

	case "RADIATOR":
		r := NewRadiatorRepository(q)
		return adapter[RadiatorRow]{list: r.List}, nil

	case "RADIATOR FAN ASSY":
		r := NewRadiatorFanAssyRepository(q)
		return adapter[RadiatorFanAssyRow]{list: r.List}, nil

	case "RADIATOR FAN MOTOR":
		r := NewRadiatorFanMotorRepository(q)
		return adapter[RadiatorFanMotorRow]{list: r.List}, nil

	case "RESISTOR":
		r := NewResistorRepository(q)
		return adapter[ResistorRow]{list: r.List}, nil

	case "ROTOR":
		r := NewRotorRepository(q)
		return adapter[RotorRow]{list: r.List}, nil

	case "STATOR":
		r := NewStatorRepository(q)
		return adapter[StatorRow]{list: r.List}, nil

	default:
		return nil, fmt.Errorf("unsupported category %q", category)
	}
}
