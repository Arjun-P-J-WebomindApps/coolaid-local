package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func ExpansionValve() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{
			"Part No", "Type", "Material", "Refrigerant", "Notes",
		},
		Rules: validation.Schema{
			"Part No":     {validation.Required()},
			"Type":        {validation.Optional(validation.Required())},
			"Material":    {validation.Optional(validation.Required())},
			"Refrigerant": {validation.Optional(validation.Required())},
			"Notes":       {validation.Optional(validation.Required())},
		},
	}
}
