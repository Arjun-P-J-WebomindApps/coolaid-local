package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func Intercooler() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{
			"Part No", "Size", "Temp Sensor", "Notes",
		},
		Rules: validation.Schema{
			"Part No":     {validation.Required()},
			"Size":        {validation.Optional(validation.Required())},
			"Temp Sensor": {validation.Optional(validation.Required())},
			"Notes":       {validation.Optional(validation.Required())},
		},
	}
}
