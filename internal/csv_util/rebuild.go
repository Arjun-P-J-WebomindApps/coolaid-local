package csv_util

import "strings"

func RebuildToFormat(
	canonicalHeaders []string,
	csvHeader []string,
	rows []map[string]string,
) ([]string, []map[string]string) {

	normalize := func(s string) string {
		return strings.ToLower(strings.TrimSpace(s))
	}

	csvHeaderMap := make(map[string]string, len(csvHeader))
	for _, h := range csvHeader {
		csvHeaderMap[normalize(h)] = h
	}

	outRows := make([]map[string]string, 0, len(rows))

	for _, row := range rows {
		nr := make(map[string]string, len(canonicalHeaders))

		for _, field := range canonicalHeaders {
			if orig, ok := csvHeaderMap[normalize(field)]; ok {
				nr[field] = row[orig]
			} else {
				nr[field] = ""
			}
		}
		outRows = append(outRows, nr)
	}

	return canonicalHeaders, outRows
}
