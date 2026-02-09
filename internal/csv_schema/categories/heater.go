package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func Heater() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{
			"Part No", "Size", "Pipe", "Type", "Notes",
		},
		Rules: validation.Schema{
			"Part No": {validation.Required()},
			"Size":    {validation.Optional(validation.Required())},
			"Pipe":    {validation.Optional(validation.Required())},
			"Type":    {validation.Optional(validation.Required())},
			"Notes":   {validation.Optional(validation.Required())},
		},
	}
}
