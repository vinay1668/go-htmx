package api

import (
	"github.com/labstack/echo"
	auth_controller "github.com/vinay1668/go-htmx/controllers/api"
)

func RegisterAuthRoutes(g *echo.Group) {
	g.POST("/login", auth_controller.HandleLogin)
	g.POST("/signup", auth_controller.HandleSignup)
	// Other auth-related API routes
}
