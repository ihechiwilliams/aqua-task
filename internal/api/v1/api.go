package v1

type API struct {
	activitiesHandler *ActivitiesHandler
	customersHandler  *CustomersHandler
	invoicesHandler   *InvoiceHandler
	resourceHandler   *ResourcesHandler
}

func NewAPI(
	activitiesHandler *ActivitiesHandler,
	customersHandler *CustomersHandler,
	invoicesHandler *InvoiceHandler,
	resourcesHandler *ResourcesHandler,
) *API {
	return &API{
		activitiesHandler: activitiesHandler,
		customersHandler:  customersHandler,
		invoicesHandler:   invoicesHandler,
		resourceHandler:   resourcesHandler,
	}
}
