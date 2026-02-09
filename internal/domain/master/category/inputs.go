package category

type CreateCategoryInput struct {
	Name  string
	Image *string
}

type UpdateCategoryInput struct {
	ID    string
	Name  *string
	Image *string
}
