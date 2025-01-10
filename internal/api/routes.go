package api

import (
	"github.com/gin-gonic/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"

	"aqua-backend/internal/api/server"
	v1 "aqua-backend/internal/api/v1"
)

func InitRoutes(router *gin.Engine, si *Routes) {
	server.RegisterHandlers(router, si)
}

func NewRoutes(apiV1 *v1.API) *Routes {
	return &Routes{
		v1: apiV1,
	}
}

// Routes is the wrapper for all the versions of the API defined by server.ServerInterface.
type Routes struct {
	v1 *v1.API
}

func (r Routes) V1CreateCustomer(c *gin.Context) {
	r.v1.V1CreateCustomer(c)
}

func (r Routes) V1GetCustomerResources(c *gin.Context, customerId openapi_types.UUID) {
	r.v1.V1GetCustomerResources(c, customerId)
}

func (r Routes) V1CreateCustomerResources(c *gin.Context, customerId openapi_types.UUID) {
	r.v1.V1CreateCustomerResources(c, customerId)
}

func (r Routes) V1DeleteResource(c *gin.Context, resourceId openapi_types.UUID) {
	r.v1.V1DeleteResource(c, resourceId)
}

func (r Routes) V1UpdateResource(c *gin.Context, resourceId openapi_types.UUID) {
	r.v1.V1UpdateResource(c, resourceId)
}
