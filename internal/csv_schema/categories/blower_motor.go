package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func BlowerMotor() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{
			"Part No", "Mounting", "Connector Type", "Impeller",
			"Resistance", "Motor Mounting", "Motor Type",
			"Voltage", "Notes",
		},
		Rules: validation.Schema{
			"Part No":        {validation.Required()},
			"Mounting":       {validation.Optional(validation.Required())},
			"Connector Type": {validation.Optional(validation.Required())},
			"Impeller":       {validation.Optional(validation.Required())},
			"Resistance":     {validation.Optional(validation.Required())},
			"Motor Mounting": {validation.Optional(validation.Required())},
			"Motor Type":     {validation.Optional(validation.Required())},
			"Voltage":        {validation.Optional(validation.Required())},
			"Notes":          {validation.Optional(validation.Required())},
		},
	}
}
