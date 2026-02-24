package search

import "errors"

var (
	ErrMissingID         = errors.New("search: missing document ID")
	ErrMissingCollection = errors.New("search: missing collection")
	ErrInvalidQuery      = errors.New("search: invalid query")
	ErrMissingPayload    = errors.New("search: missing payload")
	ErrMissingQueryBy    = errors.New("search: missing query_by")
)
