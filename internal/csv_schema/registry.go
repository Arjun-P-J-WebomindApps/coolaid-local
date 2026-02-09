package csv_schema

import (
	"fmt"
	"strings"

	categories "github.com/webomindapps-dev/coolaid-backend/internal/csv_schema/categories"
	"github.com/webomindapps-dev/coolaid-backend/internal/validation"
)

// SchemaFor returns the correct TableSchema based on upload type and category.
func SchemaFor(query, category string) (*validation.TableSchema, error) {
	switch strings.ToLower(strings.TrimSpace(query)) {

	// -------------------- Top-level sheets --------------------

	case "inventory":
		return Inventory(), nil

	case "pricing":
		return Pricing(), nil

	case "basic":
		return Basic(), nil

	// -------------------- Category-based sheets --------------------

	case "category":
		cat := strings.ToUpper(strings.TrimSpace(category))

		switch cat {
		case "ACTUATOR":
			return categories.Actuator(), nil
		case "BLOWER MOTOR":
			return categories.BlowerMotor(), nil
		case "CABIN FILTER":
			return categories.CabinFilter(), nil
		case "CHILLER UNIT":
			return categories.ChillerUnit(), nil
		case "CLUTCH ASSY":
			return categories.ClutchAssy(), nil
		case "COMPRESSOR":
			return categories.Compressor(), nil
		case "COMPRESSOR VALVE":
			return categories.CompressorValve(), nil
		case "CONDENSER":
			return categories.Condenser(), nil
		case "CONDENSER FAN ASSY":
			return categories.CondenserFanAssy(), nil
		case "EVAPORATOR":
			return categories.Evaporator(), nil
		case "EXPANSION VALVE":
			return categories.ExpansionValve(), nil
		case "FILTER DRIER":
			return categories.FilterDrier(), nil
		case "HEATER":
			return categories.Heater(), nil
		case "INTERCOOLER":
			return categories.Intercooler(), nil
		case "PRESSURE SWITCH":
			return categories.PressureSwitch(), nil
		case "RADIATOR":
			return categories.Radiator(), nil
		case "RADIATOR FAN ASSY":
			return categories.RadiatorFanAssy(), nil
		case "RADIATOR FAN MOTOR":
			return categories.RadiatorFanMotor(), nil
		case "RESISTOR":
			return categories.Resistor(), nil
		case "ROTOR":
			return categories.Rotor(), nil
		case "STATOR":
			return categories.Stator(), nil

		default:
			return nil, fmt.Errorf("invalid category %q", category)
		}

	// -------------------- Invalid --------------------

	default:
		return nil, fmt.Errorf("invalid upload type %q", query)
	}
}
