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
	"aqua-backend/internal/repositories/resources"
	resourcesMock "aqua-backend/internal/repositories/resources/mocks"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type resourcesSuite struct {
	suite.Suite
	mux           *gin.Engine
	resourcesRepo *resourcesMock.MockRepository
	customersRepo *customersMock.MockRepository
}

func (r *resourcesSuite) SetupSubTest() {
	r.mux = gin.Default()
	r.resourcesRepo = resourcesMock.NewMockRepository(r.T())
	r.customersRepo = customersMock.NewMockRepository(r.T())

	resourceHandler := v1.NewResourcesHandler(r.resourcesRepo, r.customersRepo)
	apiService := v1.NewAPI(nil, resourceHandler, nil)
	api.InitRoutes(r.mux, api.NewRoutes(apiService))
}

func TestResources(t *testing.T) {
	suite.Run(t, new(resourcesSuite))
}

func (r *resourcesSuite) Test_V1CreateCustomerResources() {
	path := "/v1/customers/388fa464-590b-4c4d-908b-d045e06bbe89/resources"

	r.Run("when request is successful", func() {
		requestBody := `{
			"data": [
				{
					"name": "AWS S3 Bucket",
					"type": "Storage",
					"region":"us-east-1"
				},
				{
					"name": "AWS EC2",
					"type": "Server",
					"region":"us-east-1"
				}
			]
		}`

		customerID := lo.Must(uuid.Parse("388fa464-590b-4c4d-908b-d045e06bbe89"))

		newResources := []*resources.Resource{
			{
				ID:         uuid.New(),
				Name:       "AWS S3 Bucket",
				Type:       "Storage",
				Region:     "us-east-1",
				CustomerID: customerID,
				CreatedAt:  time.Now().UTC(),
				UpdatedAt:  time.Now().UTC(),
			},
			{
				ID:         uuid.New(),
				Name:       "AWS EC2",
				Type:       "Server",
				Region:     "us-east-1",
				CustomerID: customerID,
				CreatedAt:  time.Now().UTC(),
				UpdatedAt:  time.Now().UTC(),
			},
		}

		r.resourcesRepo.On("CreateResourcesByCustomerID",
			mock.Anything, mock.Anything, mock.Anything,
		).Return(newResources, nil).Once()

		apitest.New().
			Handler(r.mux).
			Post(path).
			Header("Content-Type", "application/json").
			Body(requestBody).
			Expect(r.T()).
			Status(http.StatusOK).
			End()
	})

	r.Run("when bad request is passed", func() {
		requestBody := `{
			"data": [
				{
					"name": "AWS S3 Bucket"
					"type": "Storage",
					"region":"us-east-1"
				},
				{
					"name": "AWS EC2",
					"type": "Server",
					"region":"us-east-1"
				}
			]
		}`

		apitest.New().
			Handler(r.mux).
			Post(path).
			Header("Content-Type", "application/json").
			Body(requestBody).
			Expect(r.T()).
			Status(http.StatusBadRequest).
			End()
	})

	r.Run("when db fails", func() {
		requestBody := `{
			"data": [
				{
					"name": "AWS S3 Bucket",
					"type": "Storage",
					"region":"us-east-1"
				},
				{
					"name": "AWS EC2",
					"type": "Server",
					"region":"us-east-1"
				}
			]
		}`

		r.resourcesRepo.On("CreateResourcesByCustomerID",
			mock.Anything, mock.Anything, mock.Anything,
		).Return([]*resources.Resource{}, errors.New("")).Once()

		apitest.New().
			Handler(r.mux).
			Post(path).
			Header("Content-Type", "application/json").
			Body(requestBody).
			Expect(r.T()).
			Status(http.StatusUnprocessableEntity).
			End()
	})
}

func (r *resourcesSuite) Test_V1GetCustomerResources() {
	path := "/v1/customers/388fa464-590b-4c4d-908b-d045e06bbe89/resources"

	r.Run("when request is successful", func() {
		customerID := lo.Must(uuid.Parse("388fa464-590b-4c4d-908b-d045e06bbe89"))

		newResources := []*resources.Resource{
			{
				ID:         uuid.New(),
				Name:       "AWS S3 Bucket",
				Type:       "Storage",
				Region:     "us-east-1",
				CustomerID: customerID,
				CreatedAt:  time.Now().UTC(),
				UpdatedAt:  time.Now().UTC(),
			},
			{
				ID:         uuid.New(),
				Name:       "AWS EC2",
				Type:       "Server",
				Region:     "us-east-1",
				CustomerID: customerID,
				CreatedAt:  time.Now().UTC(),
				UpdatedAt:  time.Now().UTC(),
			},
		}

		r.resourcesRepo.On("GetResourcesByCustomerID",
			mock.Anything, mock.Anything,
		).Return(newResources, nil).Once()

		apitest.New().
			Handler(r.mux).
			Get(path).
			Header("Content-Type", "application/json").
			Expect(r.T()).
			Status(http.StatusOK).
			End()
	})

	r.Run("when db fails", func() {
		r.resourcesRepo.On("GetResourcesByCustomerID",
			mock.Anything, mock.Anything,
		).Return([]*resources.Resource{}, errors.New("")).Once()

		apitest.New().
			Handler(r.mux).
			Get(path).
			Header("Content-Type", "application/json").
			Expect(r.T()).
			Status(http.StatusUnprocessableEntity).
			End()
	})
}

func (r *resourcesSuite) Test_V1UpdateResource() {
	path := "/v1/resources/388fa464-590b-4c4d-908b-d045e06bbe81"

	r.Run("when request is successful", func() {
		requestBody := `{
			"name": "AWS S3 Bucket",
			"type": "Storage",
			"region":"us-east-1"
		}`

		customerID := lo.Must(uuid.Parse("388fa464-590b-4c4d-908b-d045e06bbe89"))

		newResources := &resources.Resource{
			ID:         uuid.New(),
			Name:       "AWS S3 Bucket",
			Type:       "Storage",
			Region:     "us-east-1",
			CustomerID: customerID,
			CreatedAt:  time.Now().UTC(),
			UpdatedAt:  time.Now().UTC(),
		}

		newCustomer := &customers.Customer{
			ID:        customerID,
			Name:      "John Does",
			Email:     "test@gmail.com",
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		}

		r.resourcesRepo.On("GetResourceByID",
			mock.Anything, mock.Anything,
		).Return(newResources, nil).Once()

		r.customersRepo.On("GetCustomerByID",
			mock.Anything, mock.Anything,
		).Return(newCustomer, nil).Once()

		r.resourcesRepo.On("UpdateResource",
			mock.Anything, mock.Anything,
		).Return(nil).Once()

		apitest.New().
			Handler(r.mux).
			Patch(path).
			Header("Content-Type", "application/json").
			Body(requestBody).
			Expect(r.T()).
			Status(http.StatusOK).
			End()
	})

	r.Run("when bad request is passed", func() {
		requestBody := `{
			"name": "AWS S3 Bucket"
			"type": "Storage",
			"region":"us-east-1"
		}`

		apitest.New().
			Handler(r.mux).
			Patch(path).
			Header("Content-Type", "application/json").
			Body(requestBody).
			Expect(r.T()).
			Status(http.StatusBadRequest).
			End()
	})

	r.Run("when db fails", func() {
		requestBody := `{
			"name": "AWS S3 Bucket",
			"type": "Storage",
			"region":"us-east-1"
		}`

		r.resourcesRepo.On("GetResourceByID",
			mock.Anything, mock.Anything,
		).Return(&resources.Resource{}, errors.New("")).Once()

		apitest.New().
			Handler(r.mux).
			Patch(path).
			Header("Content-Type", "application/json").
			Body(requestBody).
			Expect(r.T()).
			Status(http.StatusUnprocessableEntity).
			End()
	})
}

func (r *resourcesSuite) Test_V1DeleteResource() {
	path := "/v1/resources/388fa464-590b-4c4d-908b-d045e06bbe81"

	r.Run("when request is successful", func() {
		r.resourcesRepo.On("DeleteResource",
			mock.Anything, mock.Anything,
		).Return(nil).Once()

		apitest.New().
			Handler(r.mux).
			Delete(path).
			Header("Content-Type", "application/json").
			Expect(r.T()).
			Status(http.StatusNoContent).
			End()
	})

	r.Run("when db fails", func() {
		r.resourcesRepo.On("DeleteResource",
			mock.Anything, mock.Anything,
		).Return(errors.New("")).Once()

		apitest.New().
			Handler(r.mux).
			Delete(path).
			Header("Content-Type", "application/json").
			Expect(r.T()).
			Status(http.StatusUnprocessableEntity).
			End()
	})
}
