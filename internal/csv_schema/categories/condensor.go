package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func Condenser() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{
			"Part No", "Size", "Pipe Connector", "Drier", "Pressure Switch", "Notes",
		},
		Rules: validation.Schema{
			"Part No":         {validation.Required()},
			"Size":            {validation.Optional(validation.Required())},
			"Pipe Connector":  {validation.Optional(validation.Required())},
			"Drier":           {validation.Optional(validation.Required())},
			"Pressure Switch": {validation.Optional(validation.Required())},
			"Notes":           {validation.Optional(validation.Required())},
		},
	}
}
