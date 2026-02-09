package validation

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/webomindapps-dev/coolaid-backend/oplog"
)

func ValidateRowsAgainstSchema(
	ctx context.Context,
	query string,
	category string,
	rows []map[string]string,
	rules Schema,
) []RowError {

	var out []RowError

	for i, row := range rows {
		re := RowError{RowIndex: i + 1}

		// Field-level validation
		if errs := rules.Validate(row); errs != nil {
			re.FieldErrors = map[string]string{}
			for k, v := range errs {
				re.FieldErrors[k] = v.Error()
			}
		}

		// Cross-field rules (example: pricing vendor triplets)
		if strings.EqualFold(query, "pricing") {
			for n := 1; n <= 5; n++ {
				if ge := vendorGroupErrors(row, n); ge != nil {
					if re.FieldErrors == nil {
						re.FieldErrors = map[string]string{}
					}
					for k, v := range ge {
						re.FieldErrors[k] = v
					}
				}
			}
		}

		if len(re.FieldErrors) > 0 {
			out = append(out, re)
		}

		oplog.Info(ctx, "row validated", re.RowIndex)
	}

	return out
}

func (s Schema) Validate(input map[string]string) map[string]error {
	errs := map[string]error{}

	for field, validators := range s {
		val := input[field]
		for _, v := range validators {
			if err := v(val); err != nil {
				errs[field] = err
				break
			}
		}
	}

	if len(errs) == 0 {
		return nil
	}
	return errs
}

func vendorGroupErrors(row map[string]string, n int) map[string]string {
	errs := map[string]string{}

	nameK := fmt.Sprintf("Vendor Name %d", n)
	partK := fmt.Sprintf("Vendor Part No %d", n)
	priceK := fmt.Sprintf("Vendor Price %d", n)

	name := strings.TrimSpace(row[nameK])
	part := strings.TrimSpace(row[partK])
	price := strings.TrimSpace(row[priceK])

	allEmpty := name == "" && part == "" && price == ""
	allFilled := name != "" && part != "" && price != ""

	switch {
	case allEmpty:
		return nil
	case allFilled:
		if _, err := strconv.ParseFloat(price, 64); err != nil {
			errs[priceK] = "must be number"
		}
	default:
		if name == "" {
			errs[nameK] = "required when vendor data present"
		}
		if part == "" {
			errs[partK] = "required when vendor data present"
		}
		if price == "" {
			errs[priceK] = "required when vendor data present"
		}
	}

	if len(errs) == 0 {
		return nil
	}
	return errs
}
