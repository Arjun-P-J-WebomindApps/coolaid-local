package brand

type CreateBrandParams struct {
	ID    string
	Name  string
	Image string
}

type UpdateBrandParams struct {
	ID    string
	Name  *string
	Image *string
}
