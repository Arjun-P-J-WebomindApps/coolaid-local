package customer

type CreateCustomerParams struct {
	ID                  string
	CustomerCompanyName string
	ContactPerson       string
	Mobile              string
	Type                string
	CustomerDesignation *string
	Address             *string
	Flat                *string
	Street              *string
	City                *string
	State               *string
	Pincode             *string
	PaymentMode         *string
}

type UpdateCustomerParams struct {
	ID                  string
	CustomerCompanyName *string
	ContactPerson       *string
	Mobile              *string
	Type                *string
	CustomerDesignation *string
	Address             *string
	Flat                *string
	Street              *string
	City                *string
	State               *string
	Pincode             *string
	PaymentMode         *string
}
