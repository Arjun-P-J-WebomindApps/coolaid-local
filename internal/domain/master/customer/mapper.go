package customer

func mapRowToModel(r *CustomerRow) *Customer {
	return &Customer{
		ID:                  r.ID,
		CustomerCompanyName: r.CustomerCompanyName,
		ContactPerson:       r.ContactPerson,
		Mobile:              r.Mobile,
		Type:                r.Type,
		CustomerDesignation: r.CustomerDesignation,
		Address:             r.Address,
		Flat:                r.Flat,
		Street:              r.Street,
		City:                r.City,
		State:               r.State,
		Pincode:             r.Pincode,
		PaymentMode:         r.PaymentMode,
	}
}
