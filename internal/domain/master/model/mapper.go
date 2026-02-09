package models

func mapRowToModel(r *ModelRow) *Model {
	return &Model{
		ID:        r.ID,
		CompanyID: r.CompanyID,
		Name:      r.Name,
		ImageURL:  r.ImageURL,
	}
}

func mapRowToCompanyModel(r *ModelWithCompanyRow) *ModelWithCompanyRow {
	return &ModelWithCompanyRow{
		ID:          r.ID,
		CompanyName: r.CompanyName,
		Name:        r.Name,
		ImageURL:    r.ImageURL,
	}
}
