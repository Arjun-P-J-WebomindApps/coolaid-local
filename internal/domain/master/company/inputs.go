package company

type CreateCompanyInput struct {
	Name     string
	Status   bool
	ImageURL string
}

type UpdateCompanyInput struct {
	ID       string
	Name     *string
	Status   *bool
	ImageURL *string
}
