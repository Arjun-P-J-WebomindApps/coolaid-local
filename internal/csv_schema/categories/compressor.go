package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func Compressor() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{
			"Part No", "Compressor ID", "Oil", "Refrigerant", "Voltage",
			"Pulley Ribs", "Pulley Size", "Pipe Connector",
			"Compressor Type", "Compressor Mounting", "Connector Type", "Notes",
		},
		Rules: validation.Schema{
			"Part No":             {validation.Required()},
			"Compressor ID":       {validation.Optional(validation.Required())},
			"Oil":                 {validation.Optional(validation.Required())},
			"Refrigerant":         {validation.Optional(validation.Required())},
			"Voltage":             {validation.Optional(validation.Required())},
			"Pulley Ribs":         {validation.Optional(validation.Required())},
			"Pulley Size":         {validation.Optional(validation.Required())},
			"Pipe Connector":      {validation.Optional(validation.Required())},
			"Compressor Type":     {validation.Optional(validation.Required())},
			"Compressor Mounting": {validation.Optional(validation.Required())},
			"Connector Type":      {validation.Optional(validation.Required())},
			"Notes":               {validation.Optional(validation.Required())},
		},
	}
}
