package brand

func mapRowToModel(r *BrandRow) *Brand {
	return &Brand{
		ID:    r.ID,
		Name:  r.Name,
		Image: r.Image,
	}
}
