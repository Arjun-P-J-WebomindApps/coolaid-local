package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func Basic() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{
			"Part No", "Company", "Model", "Brand", "Category",
			"Type", "Year Start", "Year End", "Gen",
			"Fuel Type", "Transmission", "Engine CC", "Platform Code",
			"HSN Code", "Make", "Placement", "OEM Numbers",
			"Additional Info", "Unicode", "Description",
			"Vendor Name 1", "Vendor Part No 1",
			"Vendor Name 2", "Vendor Part No 2",
			"Vendor Name 3", "Vendor Part No 3",
			"Vendor Name 4", "Vendor Part No 4",
			"Vendor Name 5", "Vendor Part No 5",
		},

		Rules: validation.Schema{
			// Core identity
			"Part No":  {validation.Required()},
			"Company":  {validation.Optional(validation.Required())},
			"Model":    {validation.Optional(validation.Required())},
			"Brand":    {validation.Optional(validation.Required())},
			"Category": {validation.Optional(validation.Required())},

			// Vehicle metadata
			"Type":       {validation.Optional(validation.Required())},
			"Year Start": {validation.Optional(validation.Required()), validation.IsInt()},
			"Year End":   {validation.Optional(validation.Required()), validation.IsInt()},
			"Gen":        {validation.Optional(validation.Required())},

			// Enums / lists
			"Fuel Type": {
				validation.CSVListOfRequired("PETROL", "DIESEL", "CNG", "EV", "HYBRID"),
			},
			"Transmission": {
				validation.CSVListOfRequired("MANUAL", "AUTOMATIC"),
			},

			// Numeric / optional
			"Engine CC":     {validation.Optional(validation.IsInt())},
			"Platform Code": {validation.CSVListOfOptional()},

			// Regulatory / classification
			"HSN Code": {validation.Optional(validation.Required())},
			"Make": {
				validation.OneOfStrict("OEM", "After-Market"),
			},
			"Placement": {
				validation.OneOfStrict("NOT APPLICABLE", "FRONT", "REAR"),
			},

			// Lists / text
			"OEM Numbers":     {validation.CSVListOfOptional()},
			"Additional Info": {validation.Optional(validation.Required())},
			"Unicode":         {validation.CSVListOfOptional()},
			"Description":     {validation.Optional(validation.Required())},

			// Vendor refs (pair rules handled in engine if needed later)
			"Vendor Name 1":    {validation.Optional(validation.Required())},
			"Vendor Part No 1": {validation.Optional(validation.Required())},
			"Vendor Name 2":    {validation.Optional(validation.Required())},
			"Vendor Part No 2": {validation.Optional(validation.Required())},
			"Vendor Name 3":    {validation.Optional(validation.Required())},
			"Vendor Part No 3": {validation.Optional(validation.Required())},
			"Vendor Name 4":    {validation.Optional(validation.Required())},
			"Vendor Part No 4": {validation.Optional(validation.Required())},
			"Vendor Name 5":    {validation.Optional(validation.Required())},
			"Vendor Part No 5": {validation.Optional(validation.Required())},
		},
	}
}
