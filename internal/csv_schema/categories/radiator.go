package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func Radiator() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{
			"Part No", "Size", "Transmission", "Temp Sensor", "Tank", "Notes",
		},
		Rules: validation.Schema{
			"Part No":      {validation.Required()},
			"Size":         {validation.Optional(validation.Required())},
			"Transmission": {validation.Optional(validation.Required())},
			"Temp Sensor":  {validation.Optional(validation.Required())},
			"Tank":         {validation.Optional(validation.Required())},
			"Notes":        {validation.Optional(validation.Required())},
		},
	}
}
