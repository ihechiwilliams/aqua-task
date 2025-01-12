package v1_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"aqua-backend/internal/api"
	v1 "aqua-backend/internal/api/v1"
	"aqua-backend/internal/repositories/customers"
	customersMock "aqua-backend/internal/repositories/customers/mocks"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type customersSuite struct {
	suite.Suite
	mux           *gin.Engine
	customersRepo *customersMock.MockRepository
}

func (c *customersSuite) SetupSubTest() {
	c.mux = gin.Default()
	c.customersRepo = customersMock.NewMockRepository(c.T())

	customerHandler := v1.NewCustomersHandler(c.customersRepo)
	apiService := v1.NewAPI(customerHandler, nil, nil)
	api.InitRoutes(c.mux, api.NewRoutes(apiService))
}

func TestCustomers(t *testing.T) {
	suite.Run(t, new(customersSuite))
}

func (c *customersSuite) Test_V1CreateCustomer() {
	path := "/v1/customers"

	c.Run("when request is successful", func() {
		requestBody := `{
			"data": {
				"email": "test@gmail.com",
				"name": "John Doe"
			}
		}`

		newCustomer := &customers.Customer{
			ID:        uuid.New(),
			Name:      "John Does",
			Email:     "test@gmail.com",
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		}

		c.customersRepo.On("CreateCustomer", mock.Anything, mock.Anything).Return(newCustomer, nil).Once()

		apitest.New().
			Handler(c.mux).
			Post(path).
			Header("Content-Type", "application/json").
			Body(requestBody).
			Expect(c.T()).
			Status(http.StatusCreated).
			End()
	})

	c.Run("when bad request is passed", func() {
		requestBody := `{
			"data": {
				"email": "test@gmail.com"
				"name": "John Doe"
			}
		}`

		apitest.New().
			Handler(c.mux).
			Post(path).
			Header("Content-Type", "application/json").
			Body(requestBody).
			Expect(c.T()).
			Status(http.StatusBadRequest).
			End()
	})

	c.Run("when db fails", func() {
		requestBody := `{
			"data": {
				"email": "test@gmail.com",
				"name": "John Doe"
			}
		}`

		c.customersRepo.On("CreateCustomer", mock.Anything, mock.Anything).Return(&customers.Customer{}, errors.New("")).Once()

		apitest.New().
			Handler(c.mux).
			Post(path).
			Header("Content-Type", "application/json").
			Body(requestBody).
			Expect(c.T()).
			Status(http.StatusUnprocessableEntity).
			End()
	})
}
