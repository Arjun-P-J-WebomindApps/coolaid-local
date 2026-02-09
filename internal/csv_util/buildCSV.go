package csv_util

import (
	"fmt"
	"reflect"
	"strings"
)

func BuildCSV(headers []string, rows []any) ([][]string, error) {
	out := make([][]string, 0, len(rows))

	for _, r := range rows {
		v := reflect.ValueOf(r)
		if v.Kind() == reflect.Pointer {
			v = v.Elem()
		}

		row := make([]string, len(headers))

		for i, h := range headers {
			field := v.FieldByNameFunc(func(name string) bool {
				return normalize(name) == normalize(h)
			})

			if field.IsValid() {
				row[i] = fmt.Sprint(field.Interface())
			} else {
				row[i] = ""
			}
		}

		out = append(out, row)
	}

	return out, nil
}

func normalize(s string) string {
	return strings.ToLower(strings.ReplaceAll(strings.TrimSpace(s), " ", ""))
}
