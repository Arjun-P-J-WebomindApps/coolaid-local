package csv_schema

import "github.com/webomindapps-dev/coolaid-backend/internal/validation"

func Pricing() *validation.TableSchema {
	return &validation.TableSchema{
		Headers: []string{
			"Part No", "Company", "Model", "Brand", "Category",
			"Basic Price", "Freight", "GST",
			"Ac Workshop", "Ac Workshop %", "Ac Workshop Margin",
			"Multibrand Workshop", "Multibrand Workshop %", "Multibrand Workshop Margin",
			"Auto Trader", "Auto Trader %", "Auto Trader Margin",
			"Ac Trader", "Ac Trader %", "Ac Trader Margin",
			"OEM MRP", "Unit Measure",
			"Vendor Name 1", "Vendor Part No 1", "Vendor Price 1",
			"Vendor Name 2", "Vendor Part No 2", "Vendor Price 2",
			"Vendor Name 3", "Vendor Part No 3", "Vendor Price 3",
			"Vendor Name 4", "Vendor Part No 4", "Vendor Price 4",
			"Vendor Name 5", "Vendor Part No 5", "Vendor Price 5",
		},
		Rules: validation.Schema{
			"Part No":  {validation.Required()},
			"Company":  {validation.Optional(validation.Required())},
			"Model":    {validation.Optional(validation.Required())},
			"Brand":    {validation.Optional(validation.Required())},
			"Category": {validation.Optional(validation.Required())},

			"Basic Price": {validation.Optional(validation.Required()), validation.IsFloat()},
			"Freight":     {validation.Optional(validation.Required()), validation.IsFloat()},
			"GST":         {validation.Optional(validation.Required()), validation.Percent()},

			"Ac Workshop":        {validation.Optional(validation.Required()), validation.IsFloat()},
			"Ac Workshop %":      {validation.Optional(validation.Required()), validation.Percent()},
			"Ac Workshop Margin": {validation.Optional(validation.Required()), validation.IsFloat()},

			"Multibrand Workshop":        {validation.Optional(validation.Required()), validation.IsFloat()},
			"Multibrand Workshop %":      {validation.Optional(validation.Required()), validation.Percent()},
			"Multibrand Workshop Margin": {validation.Optional(validation.Required()), validation.IsFloat()},

			"Auto Trader":        {validation.Optional(validation.Required()), validation.IsFloat()},
			"Auto Trader %":      {validation.Optional(validation.Required()), validation.Percent()},
			"Auto Trader Margin": {validation.Optional(validation.Required()), validation.IsFloat()},

			"Ac Trader":        {validation.Optional(validation.Required()), validation.IsFloat()},
			"Ac Trader %":      {validation.Optional(validation.Required()), validation.Percent()},
			"Ac Trader Margin": {validation.Optional(validation.Required()), validation.IsFloat()},

			"OEM MRP":      {validation.Optional(validation.Required()), validation.IsFloat()},
			"Unit Measure": {validation.Optional(validation.Required())},

			// vendor fields — cross-field rules handled in validation engine
			"Vendor Name 1":    {validation.Optional(validation.Required())},
			"Vendor Part No 1": {validation.Optional(validation.Required())},
			"Vendor Price 1":   {validation.Optional(validation.IsFloat())},

			"Vendor Name 2":    {validation.Optional(validation.Required())},
			"Vendor Part No 2": {validation.Optional(validation.Required())},
			"Vendor Price 2":   {validation.Optional(validation.IsFloat())},

			"Vendor Name 3":    {validation.Optional(validation.Required())},
			"Vendor Part No 3": {validation.Optional(validation.Required())},
			"Vendor Price 3":   {validation.Optional(validation.IsFloat())},

			"Vendor Name 4":    {validation.Optional(validation.Required())},
			"Vendor Part No 4": {validation.Optional(validation.Required())},
			"Vendor Price 4":   {validation.Optional(validation.IsFloat())},

			"Vendor Name 5":    {validation.Optional(validation.Required())},
			"Vendor Part No 5": {validation.Optional(validation.Required())},
			"Vendor Price 5":   {validation.Optional(validation.IsFloat())},
		},
	}
}
