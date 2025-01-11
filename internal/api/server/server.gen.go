// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// CustomerRequestBodyData defines model for CustomerRequestBodyData.
type CustomerRequestBodyData struct {
	Email openapi_types.Email `json:"email"`
	Name  string              `json:"name"`
}

// CustomerResourceRequestBodyData defines model for CustomerResourceRequestBodyData.
type CustomerResourceRequestBodyData struct {
	// Name The name of the resource.
	Name string `json:"name"`

	// Region The region where the resource is located.
	Region string `json:"region"`

	// Type The type of the resource.
	Type string `json:"type"`
}

// CustomerResourceResponseData defines model for CustomerResourceResponseData.
type CustomerResourceResponseData struct {
	// Id The unique identifier for the resource.
	Id openapi_types.UUID `json:"id"`

	// Name The name of the resource.
	Name string `json:"name"`

	// Region The region where the resource is located.
	Region string `json:"region"`

	// Type The type of the resource.
	Type string `json:"type"`
}

// CustomerResponseData defines model for CustomerResponseData.
type CustomerResponseData struct {
	Email string             `json:"email"`
	Id    openapi_types.UUID `json:"id"`
	Name  string             `json:"name"`
}

// Error defines model for Error.
type Error struct {
	Code   string                  `json:"code"`
	Detail string                  `json:"detail"`
	Meta   *map[string]interface{} `json:"meta,omitempty"`
	Status int                     `json:"status"`
	Title  string                  `json:"title"`
}

// ErrorResponse Response that contains the list of errors
type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

// UpdateResource defines model for UpdateResource.
type UpdateResource struct {
	Name   *string `json:"name,omitempty"`
	Region *string `json:"region,omitempty"`
	Type   *string `json:"type,omitempty"`
}

// UserNotificationResponseData defines model for UserNotificationResponseData.
type UserNotificationResponseData struct {
	// Id The unique identifier for the notification.
	Id openapi_types.UUID `json:"id"`

	// Message The content of the notification.
	Message string `json:"message"`

	// UserId The user if.
	UserId string `json:"user_id"`
}

// CustomerResourceResponse defines model for CustomerResourceResponse.
type CustomerResourceResponse struct {
	Data []CustomerResourceResponseData `json:"data"`
}

// CustomerResponse defines model for CustomerResponse.
type CustomerResponse struct {
	Data CustomerResponseData `json:"data"`
}

// UserNotificationsResponse defines model for UserNotificationsResponse.
type UserNotificationsResponse struct {
	Data []UserNotificationResponseData `json:"data"`
}

// CreateCustomerRequestBody defines model for CreateCustomerRequestBody.
type CreateCustomerRequestBody struct {
	Data CustomerRequestBodyData `json:"data"`
}

// CreateCustomerResourceRequestBody defines model for CreateCustomerResourceRequestBody.
type CreateCustomerResourceRequestBody struct {
	Data []CustomerResourceRequestBodyData `json:"data"`
}

// V1CreateCustomerJSONBody defines parameters for V1CreateCustomer.
type V1CreateCustomerJSONBody struct {
	Data CustomerRequestBodyData `json:"data"`
}

// V1CreateCustomerResourcesJSONBody defines parameters for V1CreateCustomerResources.
type V1CreateCustomerResourcesJSONBody struct {
	Data []CustomerResourceRequestBodyData `json:"data"`
}

// V1CreateCustomerJSONRequestBody defines body for V1CreateCustomer for application/json ContentType.
type V1CreateCustomerJSONRequestBody V1CreateCustomerJSONBody

// V1CreateCustomerResourcesJSONRequestBody defines body for V1CreateCustomerResources for application/json ContentType.
type V1CreateCustomerResourcesJSONRequestBody V1CreateCustomerResourcesJSONBody

// V1UpdateResourceJSONRequestBody defines body for V1UpdateResource for application/json ContentType.
type V1UpdateResourceJSONRequestBody = UpdateResource

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Register a new customer
	// (POST /v1/customers)
	V1CreateCustomer(c *gin.Context)
	// Get all cloud resources for a customer
	// (GET /v1/customers/{customer_id}/resources)
	V1GetCustomerResources(c *gin.Context, customerId openapi_types.UUID)
	// Add cloud resources to a customer
	// (POST /v1/customers/{customer_id}/resources)
	V1CreateCustomerResources(c *gin.Context, customerId openapi_types.UUID)
	// Delete a notification
	// (DELETE /v1/notification/{notification_id})
	V1DeleteNotification(c *gin.Context, notificationId openapi_types.UUID)
	// Delete a notification
	// (DELETE /v1/notifications/{user_id})
	V1DeleteUserNotifications(c *gin.Context, userId string)
	// Get all notifications for a user
	// (GET /v1/notifications/{user_id})
	V1GetUserNotifications(c *gin.Context, userId string)
	// Delete a resource
	// (DELETE /v1/resources/{resource_id})
	V1DeleteResource(c *gin.Context, resourceId openapi_types.UUID)
	// Update a cloud resource
	// (PATCH /v1/resources/{resource_id})
	V1UpdateResource(c *gin.Context, resourceId openapi_types.UUID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// V1CreateCustomer operation middleware
func (siw *ServerInterfaceWrapper) V1CreateCustomer(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.V1CreateCustomer(c)
}

// V1GetCustomerResources operation middleware
func (siw *ServerInterfaceWrapper) V1GetCustomerResources(c *gin.Context) {

	var err error

	// ------------- Path parameter "customer_id" -------------
	var customerId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "customer_id", c.Param("customer_id"), &customerId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter customer_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.V1GetCustomerResources(c, customerId)
}

// V1CreateCustomerResources operation middleware
func (siw *ServerInterfaceWrapper) V1CreateCustomerResources(c *gin.Context) {

	var err error

	// ------------- Path parameter "customer_id" -------------
	var customerId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "customer_id", c.Param("customer_id"), &customerId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter customer_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.V1CreateCustomerResources(c, customerId)
}

// V1DeleteNotification operation middleware
func (siw *ServerInterfaceWrapper) V1DeleteNotification(c *gin.Context) {

	var err error

	// ------------- Path parameter "notification_id" -------------
	var notificationId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "notification_id", c.Param("notification_id"), &notificationId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter notification_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.V1DeleteNotification(c, notificationId)
}

// V1DeleteUserNotifications operation middleware
func (siw *ServerInterfaceWrapper) V1DeleteUserNotifications(c *gin.Context) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", c.Param("user_id"), &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter user_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.V1DeleteUserNotifications(c, userId)
}

// V1GetUserNotifications operation middleware
func (siw *ServerInterfaceWrapper) V1GetUserNotifications(c *gin.Context) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", c.Param("user_id"), &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter user_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.V1GetUserNotifications(c, userId)
}

// V1DeleteResource operation middleware
func (siw *ServerInterfaceWrapper) V1DeleteResource(c *gin.Context) {

	var err error

	// ------------- Path parameter "resource_id" -------------
	var resourceId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "resource_id", c.Param("resource_id"), &resourceId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter resource_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.V1DeleteResource(c, resourceId)
}

// V1UpdateResource operation middleware
func (siw *ServerInterfaceWrapper) V1UpdateResource(c *gin.Context) {

	var err error

	// ------------- Path parameter "resource_id" -------------
	var resourceId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "resource_id", c.Param("resource_id"), &resourceId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter resource_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.V1UpdateResource(c, resourceId)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/v1/customers", wrapper.V1CreateCustomer)
	router.GET(options.BaseURL+"/v1/customers/:customer_id/resources", wrapper.V1GetCustomerResources)
	router.POST(options.BaseURL+"/v1/customers/:customer_id/resources", wrapper.V1CreateCustomerResources)
	router.DELETE(options.BaseURL+"/v1/notification/:notification_id", wrapper.V1DeleteNotification)
	router.DELETE(options.BaseURL+"/v1/notifications/:user_id", wrapper.V1DeleteUserNotifications)
	router.GET(options.BaseURL+"/v1/notifications/:user_id", wrapper.V1GetUserNotifications)
	router.DELETE(options.BaseURL+"/v1/resources/:resource_id", wrapper.V1DeleteResource)
	router.PATCH(options.BaseURL+"/v1/resources/:resource_id", wrapper.V1UpdateResource)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AP//ie+K+94pAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}