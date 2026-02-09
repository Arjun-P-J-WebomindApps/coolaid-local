package customer

import "errors"

var (
	ErrCustomerNotFound = errors.New("customer not found")
	ErrCustomerExists   = errors.New("customer already exists")
	ErrInvalidInput     = errors.New("invalid input")
	ErrInternal         = errors.New("internal error")
)
