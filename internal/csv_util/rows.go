package csv_util

import "strings"

func ParseRows(
	records [][]string,
	headers []string,
	index map[string]int,
) []map[string]string {

	var rows []map[string]string

	for _, record := range records[1:] {
		row := make(map[string]string, len(headers))

		for _, h := range headers {
			i := index[strings.ToLower(h)]
			if i < len(record) {
				row[h] = strings.TrimSpace(record[i])
			} else {
				row[h] = ""
			}
		}
		rows = append(rows, row)
	}

	return rows
}
