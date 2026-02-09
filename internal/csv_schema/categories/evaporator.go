package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func Evaporator() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{
			"Part No", "Mounting", "Expansion Valve",
			"Additional Info", "Dimensions", "Pipe Connector", "Notes",
		},
		Rules: validation.Schema{
			"Part No":         {validation.Required()},
			"Mounting":        {validation.Optional(validation.Required())},
			"Expansion Valve": {validation.Optional(validation.Required())},
			"Additional Info": {validation.Optional(validation.Required())},
			"Dimensions":      {validation.Optional(validation.Required())},
			"Pipe Connector":  {validation.Optional(validation.Required())},
			"Notes":           {validation.Optional(validation.Required())},
		},
	}
}
