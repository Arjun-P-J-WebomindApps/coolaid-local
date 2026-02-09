package csv_util

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
)

func Read(r io.Reader) ([][]string, error) {
	br := bufio.NewReader(r)

	//Handle BOM
	if b, _ := br.Peek(3); len(b) == 3 && b[0] == 0xEE && b[1] == 0xBB && b[2] == 0xBF {
		_, _ = br.Read(make([]byte, 3))
	}

	cr := csv.NewReader(br)
	cr.FieldsPerRecord = -1 // allow ragged rows

	records, err := cr.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("read csv: %w", err)
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("empty csv")
	}

	return records, nil
}
