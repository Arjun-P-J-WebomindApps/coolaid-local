package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func RadiatorFanAssy() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{
			"Part No", "Voltage", "Motor Type", "Resistance",
			"Number Of Sockets", "Shroud", "Connector Type",
			"Fan Blade Diameter", "Number Of Blades", "Notes",
		},
		Rules: validation.Schema{
			"Part No":            {validation.Required()},
			"Voltage":            {validation.Optional(validation.Required())},
			"Motor Type":         {validation.Optional(validation.Required())},
			"Resistance":         {validation.Optional(validation.Required())},
			"Number Of Sockets":  {validation.Optional(validation.IsInt())},
			"Shroud":             {validation.Optional(validation.Required())},
			"Connector Type":     {validation.Optional(validation.Required())},
			"Fan Blade Diameter": {validation.Optional(validation.Required())},
			"Number Of Blades":   {validation.Optional(validation.IsInt())},
			"Notes":              {validation.Optional(validation.Required())},
		},
	}
}
