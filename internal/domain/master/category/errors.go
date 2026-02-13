package category

import "errors"

var (
	ErrCategoryNotFound = errors.New("category not found")
	ErrCategoryExists   = errors.New("category already exists")
	ErrInternal         = errors.New("internal error")
	ErrInvalidInput     = errors.New("invalid input")
)
