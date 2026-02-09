package company

func mapRowToModel(r *CompanyRow) *Company {
	return &Company{
		ID:       r.ID,
		Name:     r.Name,
		Status:   r.Status,
		ImageURL: r.ImageURL,
	}
}
