package brand

type CreateBrandInput struct {
	Name  string
	Image string
}

type UpdateBrandInput struct {
	ID    string
	Name  *string
	Image *string
}
