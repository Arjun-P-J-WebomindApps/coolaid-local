package models

import "errors"

var (
	ErrModelNotFound   = errors.New("model not found")
	ErrModelExists     = errors.New("model already exists")
	ErrCompanyNotFound = errors.New("company not found")
	ErrInvalidInput    = errors.New("invalid input")
	ErrInternal        = errors.New("internal error")
)
