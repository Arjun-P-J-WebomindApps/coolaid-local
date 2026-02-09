package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func FilterDrier() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{
			"Part No", "Pipe Connector", "Size", "Pressure Switch", "Notes",
		},
		Rules: validation.Schema{
			"Part No":         {validation.Required()},
			"Pipe Connector":  {validation.Optional(validation.Required())},
			"Size":            {validation.Optional(validation.Required())},
			"Pressure Switch": {validation.Optional(validation.Required())},
			"Notes":           {validation.Optional(validation.Required())},
		},
	}
}
