package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func PressureSwitch() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{
			"Part No", "Connector Type", "Thread Type", "Notes",
		},
		Rules: validation.Schema{
			"Part No":        {validation.Required()},
			"Connector Type": {validation.Optional(validation.Required())},
			"Thread Type":    {validation.Optional(validation.Required())},
			"Notes":          {validation.Optional(validation.Required())},
		},
	}
}
