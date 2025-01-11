package v1

import (
	"aqua-backend/internal/api/server"
	"aqua-backend/internal/repositories/customers"
	"aqua-backend/internal/repositories/resources"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/samber/lo"
	"net/http"
)

type ResourcesHandler struct {
	resourcesRepo resources.Repository
	customerRepo  customers.Repository
}

func NewResourcesHandler(resourcesRepo resources.Repository, customerRepo customers.Repository) *ResourcesHandler {
	return &ResourcesHandler{
		resourcesRepo: resourcesRepo,
		customerRepo:  customerRepo,
	}
}

func (a *API) V1CreateCustomerResources(c *gin.Context, customerId openapi_types.UUID) {
	var reqBody server.V1CreateCustomerResourcesJSONRequestBody

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		server.BadRequestError(err, c)
		return
	}

	resourcesData := reqBody.Data
	dBResources := lo.Map(resourcesData, func(resource server.CustomerResourceRequestBodyData, _ int) *resources.DBResource {
		return &resources.DBResource{
			ID:         uuid.New(),
			Name:       resource.Name,
			Type:       resource.Type,
			Region:     resource.Region,
			CustomerID: customerId,
		}
	})

	result, err := a.resourceHandler.resourcesRepo.CreateResourcesByCustomerID(c.Request.Context(), customerId, dBResources)
	if err != nil {
		server.ProcessingError(err, c)

		return
	}

	c.JSON(http.StatusOK, lo.Map(result, func(resource *resources.Resource, _ int) server.CustomerResourceResponseData {
		return serializeResourceToAPIResponseData(resource)
	}))
}

func (a *API) V1GetCustomerResources(c *gin.Context, customerId openapi_types.UUID) {
	result, err := a.resourceHandler.resourcesRepo.GetResourcesByCustomerID(c.Request.Context(), customerId)
	if err != nil {
		server.ProcessingError(err, c)

		return
	}

	c.JSON(http.StatusOK, lo.Map(result, func(resource *resources.Resource, _ int) server.CustomerResourceResponseData {
		return serializeResourceToAPIResponseData(resource)
	}))
}

func (a *API) V1UpdateResource(c *gin.Context, resourceId openapi_types.UUID) {
	var reqBody server.V1UpdateResourceJSONRequestBody

	// Decode the incoming JSON request body
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		server.BadRequestError(err, c)
		return
	}

	if *reqBody.Name == "" || *reqBody.Region == "" || *reqBody.Type == "" {
		server.BadRequestError(errors.New("missing fields"), c)
		return
	}

	resource, err := a.resourceHandler.resourcesRepo.GetResourceByID(c.Request.Context(), resourceId)
	if err != nil {
		server.ProcessingError(err, c)

		return
	}

	_, err = a.resourceHandler.customerRepo.GetCustomerByID(c.Request.Context(), resource.CustomerID)
	if err != nil {
		server.ProcessingError(err, c)

		return
	}

	resource.Name = *reqBody.Name
	resource.Type = *reqBody.Type
	resource.Region = *reqBody.Region

	err = a.resourceHandler.resourcesRepo.UpdateResource(c.Request.Context(), resource)
	if err != nil {
		server.ProcessingError(err, c)

		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "resource updated"})
}

func (a *API) V1DeleteResource(c *gin.Context, resourceId openapi_types.UUID) {
	err := a.resourceHandler.resourcesRepo.DeleteResource(c.Request.Context(), resourceId)
	if err != nil {
		server.ProcessingError(err, c)

		return
	}

	c.JSON(http.StatusNoContent, gin.H{"msg": "resource deleted"})
}

func serializeResourceToAPIResponseData(resource *resources.Resource) server.CustomerResourceResponseData {
	return server.CustomerResourceResponseData{
		Id:     resource.ID,
		Name:   resource.Name,
		Region: resource.Region,
		Type:   resource.Type,
	}
}
