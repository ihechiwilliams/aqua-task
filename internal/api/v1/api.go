package v1

type API struct {
	customersHandler    *CustomersHandler
	resourceHandler     *ResourcesHandler
	notificationHandler *NotificationsHandler
}

func NewAPI(
	customersHandler *CustomersHandler,
	resourcesHandler *ResourcesHandler,
	notificationHandler *NotificationsHandler,
) *API {
	return &API{
		customersHandler:    customersHandler,
		resourceHandler:     resourcesHandler,
		notificationHandler: notificationHandler,
	}
}
