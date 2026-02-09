package category

type CreateCategoryParams struct {
	ID    string
	Name  string
	Image *string
}

type UpdateCategoryParams struct {
	ID    string
	Name  *string
	Image *string
}
