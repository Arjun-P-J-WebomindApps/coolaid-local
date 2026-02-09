package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func CabinFilter() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{
			"Part No", "Type", "Dimensions", "Material", "Notes",
		},
		Rules: validation.Schema{
			"Part No":    {validation.Required()},
			"Type":       {validation.Optional(validation.Required())},
			"Dimensions": {validation.Optional(validation.Required())},
			"Material":   {validation.Optional(validation.Required())},
			"Notes":      {validation.Optional(validation.Required())},
		},
	}
}
