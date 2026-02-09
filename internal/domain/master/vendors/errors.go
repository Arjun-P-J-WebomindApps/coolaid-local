package vendor

import "errors"

var (
	ErrVendorNotFound = errors.New("vendor not found")
	ErrVendorExists   = errors.New("vendor already exists")
	ErrInternal       = errors.New("internal error")
)
