package web

import (
	"github.com/labstack/echo"
	dashboard_controller "github.com/vinay1668/go-htmx/controllers/web"
)

func RegisterDashboardRoutes(e *echo.Group) {
	e.GET("/", dashboard_controller.RenderDashboard)
}
