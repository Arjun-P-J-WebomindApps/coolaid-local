package company

import "errors"

var (
	ErrCompanyNotFound = errors.New("company not found")
	ErrCompanyExists   = errors.New("company already exists")
	ErrInvalidInput    = errors.New("invalid input")
	ErrInternal        = errors.New("internal error")
)
