package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func Actuator() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{
			"Part No", "Connector Type", "Mounting",
			"Voltage", "Rotation Angle", "Notes",
		},
		Rules: validation.Schema{
			"Part No":        {validation.Required()},
			"Connector Type": {validation.Optional(validation.Required())},
			"Mounting":       {validation.Optional(validation.Required())},
			"Voltage":        {validation.Optional(validation.Required())},
			"Rotation Angle": {validation.Optional(validation.Required())},
			"Notes":          {validation.Optional(validation.Required())},
		},
	}
}
