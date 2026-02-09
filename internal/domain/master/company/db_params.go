package company

import "time"

type CreateCompanyParams struct {
	ID        string
	Name      string
	Status    bool
	ImageURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateCompanyParams struct {
	ID        string
	Name      *string
	Status    *bool
	ImageURL  *string
	UpdatedAt time.Time
}
