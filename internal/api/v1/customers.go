package v1

import (
	"net/http"

	"aqua-backend/internal/api/server"
	"aqua-backend/internal/repositories/customers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CustomersHandler struct {
	customersRepo customers.Repository
}

func NewCustomersHandler(customersRepo customers.Repository) *CustomersHandler {
	return &CustomersHandler{
		customersRepo: customersRepo,
	}
}

func (a *API) V1CreateCustomer(c *gin.Context) {
	var reqBody server.V1CreateCustomerJSONRequestBody

	// Decode the incoming JSON request body
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		server.BadRequestError(err, c)
		return
	}

	customerData := reqBody.Data

	// Create a new customer object
	newCustomer := &customers.DBCustomer{
		ID:    uuid.New(),
		Name:  customerData.Name,
		Email: string(customerData.Email),
	}

	// Call the customer handler's CreateCustomer method
	result, err := a.customersHandler.customersRepo.CreateCustomer(c.Request.Context(), newCustomer)
	if err != nil {
		server.BadRequestError(err, c)
		return
	}

	// Serialize the customer to API response format
	response := serializeCustomerToAPIResponse(result)

	// Set the HTTP status and send the response as JSON
	c.JSON(http.StatusCreated, response)
}

func serializeCustomerToAPIResponse(customer *customers.Customer) server.CustomerResponseData {
	return server.CustomerResponseData{
		Email: customer.Email,
		Id:    customer.ID,
		Name:  customer.Name,
	}
}
