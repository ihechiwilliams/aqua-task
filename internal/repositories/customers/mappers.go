package customers

func FromDBCustomer(dbCustomer *DBCustomer) *Customer {
	return &Customer{
		ID:        dbCustomer.ID,
		Name:      dbCustomer.Name,
		Email:     dbCustomer.Email,
		CreatedAt: dbCustomer.CreatedAt,
		UpdatedAt: dbCustomer.UpdatedAt,
	}
}
