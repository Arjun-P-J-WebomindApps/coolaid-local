package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func Inventory() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{
			"Part No",
			"Company",
			"Model",
			"Brand",
			"Qty In Stock",
		},
		Rules: validation.Schema{
			"Part No":      {validation.Required()},
			"Company":      {validation.Required()},
			"Model":        {validation.Optional(validation.Required())},
			"Brand":        {validation.Optional(validation.Required())},
			"Qty In Stock": {validation.Optional(validation.IsInt())},
		},
	}
}
