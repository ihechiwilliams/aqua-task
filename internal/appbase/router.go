package appbase

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	ginTrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
)

func NewRouterGin(serviceName string, timeout time.Duration) *gin.Engine {
	r := gin.Default()

	// Request ID Middleware
	r.Use(requestid.New())

	// Recovery Middleware
	r.Use(gin.Recovery())

	// Timeout Middleware
	r.Use(func(c *gin.Context) {
		// Create a new context with a timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// Set the new context with timeout into the Gin context
		c.Request = c.Request.WithContext(ctx)

		// Continue with the next middleware/handler
		c.Next()
	})

	// DataDog Tracing Middleware
	r.Use(ginTrace.Middleware(
		serviceName,
		ginTrace.WithIgnoreRequest(func(c *gin.Context) bool {
			return c.Request.URL.Path == "/" || c.Request.URL.Path == "/healthz" || c.Request.URL.Path == "/readyz"
		}),
	))

	// Heartbeat Routes
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	r.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	r.GET("/readyz", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	return r
}
