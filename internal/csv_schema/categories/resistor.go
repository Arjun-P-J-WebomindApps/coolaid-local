package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func Resistor() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{
			"Part No", "Type", "Connector Type", "Voltage", "Notes",
		},
		Rules: validation.Schema{
			"Part No":        {validation.Required()},
			"Type":           {validation.Optional(validation.Required())},
			"Connector Type": {validation.Optional(validation.Required())},
			"Voltage":        {validation.Optional(validation.Required())},
			"Notes":          {validation.Optional(validation.Required())},
		},
	}
}
