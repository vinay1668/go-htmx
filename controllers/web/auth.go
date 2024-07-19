package web

import (
	"context"

	"github.com/labstack/echo"
	"github.com/vinay1668/go-htmx/templates/common"
)

func RenderLogin(c echo.Context) error {
	loginComponent := common.Login()
	return loginComponent.Render(context.Background(), c.Response())
}

func RenderSignup(c echo.Context) error {
	signupComponent := common.Signup()
	return signupComponent.Render(context.Background(), c.Response())
}
