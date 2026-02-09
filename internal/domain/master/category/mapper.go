package category

func mapRowToModel(r *CategoryRow) *Category {
	return &Category{
		ID:    r.ID,
		Name:  r.Name,
		Image: r.Image,
	}
}
