package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func CompressorValve() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{
			"Part No", "Type", "Voltage", "Connector Type",
			"Compressor Details", "Notes",
		},
		Rules: validation.Schema{
			"Part No":            {validation.Required()},
			"Type":               {validation.Optional(validation.Required())},
			"Voltage":            {validation.Optional(validation.Required())},
			"Connector Type":     {validation.Optional(validation.Required())},
			"Compressor Details": {validation.Optional(validation.Required())},
			"Notes":              {validation.Optional(validation.Required())},
		},
	}
}
