package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func ClutchAssy() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{
			"Part No", "Pulley Ribs", "Pulley Size",
			"Compressor Details", "Connector Type",
			"Voltage", "Shaft Type", "Notes",
		},
		Rules: validation.Schema{
			"Part No":            {validation.Required()},
			"Pulley Ribs":        {validation.Optional(validation.Required())},
			"Pulley Size":        {validation.Optional(validation.Required())},
			"Compressor Details": {validation.Optional(validation.Required())},
			"Connector Type":     {validation.Optional(validation.Required())},
			"Voltage":            {validation.Optional(validation.Required())},
			"Shaft Type":         {validation.Optional(validation.Required())},
			"Notes":              {validation.Optional(validation.Required())},
		},
	}
}
