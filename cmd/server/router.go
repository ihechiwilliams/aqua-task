package main

import (
	"github.com/gin-gonic/gin"

	"aqua-backend/internal/api"
	"aqua-backend/internal/appbase"

	"github.com/samber/do"
)

func buildRouter(app *appbase.AppBase) *gin.Engine {
	mux := do.MustInvokeNamed[*gin.Engine](app.Injector, appbase.InjectorApplicationRouter)
	routes := do.MustInvoke[*api.Routes](app.Injector)

	api.InitRoutes(mux, routes)

	return mux
}
