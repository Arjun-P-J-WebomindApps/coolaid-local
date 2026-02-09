package models

type CreateModelInput struct {
	CompanyName string
	Name        string
	ImageURL    string
}

type UpdateModelInput struct {
	ID          string
	CompanyName *string
	Name        *string
	ImageURL    *string
}
