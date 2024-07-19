package routes

import (
	"github.com/labstack/echo"
	"github.com/vinay1668/go-htmx/routes/api"
	"github.com/vinay1668/go-htmx/routes/web"
)

func RegisterRoutes(e *echo.Echo) {

	// grouping routes
	apiRoutes := e.Group("/api")
	appRoutes := e.Group("")
	// Frontend/web routes
	api.RegisterAuthRoutes(apiRoutes)
	web.RegisterAuthRoutes(appRoutes)
	web.RegisterDashboardRoutes(appRoutes)
	// web.RegisterProfileRoutes(web)

	// ... other frontend route registrations

	// API routes
	api.RegisterAuthRoutes(apiRoutes)

}
