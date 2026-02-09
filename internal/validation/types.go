package validation

type Schema map[string][]Validator

type TableSchema struct {
	Headers []string
	Rules   Schema
}

type RowError struct {
	RowIndex    int
	FieldErrors map[string]string
}
