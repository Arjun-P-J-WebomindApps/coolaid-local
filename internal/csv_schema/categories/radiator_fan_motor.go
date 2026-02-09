package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func RadiatorFanMotor() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{
			"Part No", "Fan Blade Diameter", "Number Of Blades",
			"Voltage", "Number Of Sockets", "Connector Type", "Notes",
		},
		Rules: validation.Schema{
			"Part No":            {validation.Required()},
			"Fan Blade Diameter": {validation.Optional(validation.Required())},
			"Number Of Blades":   {validation.Optional(validation.IsInt())},
			"Voltage":            {validation.Optional(validation.Required())},
			"Number Of Sockets":  {validation.Optional(validation.IsInt())},
			"Connector Type":     {validation.Optional(validation.Required())},
			"Notes":              {validation.Optional(validation.Required())},
		},
	}
}
