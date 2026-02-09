package models

type CreateModelParams struct {
	ID        string
	CompanyID string
	Name      string
	ImageURL  string
}

type UpdateModelParams struct {
	ID        string
	CompanyID *string
	Name      *string
	ImageURL  *string
}
