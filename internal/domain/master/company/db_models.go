package company

import "time"

type CompanyRow struct {
	ID        string
	Name      string
	Status    bool
	ImageURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
