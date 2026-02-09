package models

type ModelRow struct {
	ID        string
	CompanyID string
	Name      string
	ImageURL  string
}

type ModelWithCompanyRow struct {
	ID          string
	CompanyName string
	Name        string
	ImageURL    string
}
