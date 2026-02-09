package csv_util

import (
	"encoding/csv"
	"fmt"
	"io"
)

func Write(headers []string, rows [][]string, w io.Writer) error {
	if len(headers) == 0 {
		return fmt.Errorf("headers cannot be empty")
	}

	cw := csv.NewWriter(w)

	if err := cw.Write(headers); err != nil {
		return err
	}

	for _, row := range rows {
		if len(row) != len(headers) {
			return fmt.Errorf("row length mismatch")
		}
		if err := cw.Write(row); err != nil {
			return err
		}
	}

	cw.Flush()
	return cw.Error()
}
