package web

import (
	"context"

	"github.com/labstack/echo"
	"github.com/vinay1668/go-htmx/templates"
)

func RenderDashboard(c echo.Context) error {
	dashboardComponent := templates.Index()
	return dashboardComponent.Render(context.Background(), c.Response())
}
