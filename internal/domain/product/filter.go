package product

type FilterParams struct {
	CompanyID  *string
	BrandID    *string
	CategoryID *string
	ModelID    *string
	IsActive   *bool
	Search     *string
	Page       int
	Limit      int
}

type SimilarPricingParams struct {
	PartNo string
}
