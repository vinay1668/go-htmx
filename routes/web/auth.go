package web

import (
	"github.com/labstack/echo"
	auth_controller "github.com/vinay1668/go-htmx/controllers/web"
)

func RegisterAuthRoutes(g *echo.Group) {
	g.GET("/login", auth_controller.RenderLogin)
	g.GET("/signup", auth_controller.RenderSignup)
	// Other auth-related frontend routes
}
