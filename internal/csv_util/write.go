package csv_util

import (
	"encoding/csv"
	"fmt"
	"io"
)

/*
WriteCSV streams a CSV file to the provided writer.

Responsibilities:
- Writes headers first
- Writes rows in order
- Flushes writer
- Does NOT mutate data
- Does NOT add BOM (caller decides)

Used by:
- download service
- validation error export
*/
func WriteCSV(
	headers []string,
	rows [][]string,
	w io.Writer,
) error {

	if len(headers) == 0 {
		return fmt.Errorf("csv: headers cannot be empty")
	}

	cw := csv.NewWriter(w)

	// write header
	if err := cw.Write(headers); err != nil {
		return fmt.Errorf("csv: write header failed: %w", err)
	}

	// write rows
	for i, row := range rows {
		if len(row) != len(headers) {
			return fmt.Errorf(
				"csv: row %d length mismatch (got %d, want %d)",
				i,
				len(row),
				len(headers),
			)
		}

		if err := cw.Write(row); err != nil {
			return fmt.Errorf("csv: write row %d failed: %w", i, err)
		}
	}

	cw.Flush()

	if err := cw.Error(); err != nil {
		return fmt.Errorf("csv: flush failed: %w", err)
	}

	return nil
}
