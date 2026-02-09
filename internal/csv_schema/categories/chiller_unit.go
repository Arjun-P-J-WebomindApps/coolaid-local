package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func ChillerUnit() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{"Part No", "Type", "Voltage", "Notes"},
		Rules: validation.Schema{
			"Part No": {validation.Required()},
			"Type":    {validation.Optional(validation.Required())},
			"Voltage": {validation.Optional(validation.Required())},
			"Notes":   {validation.Optional(validation.Required())},
		},
	}
}
