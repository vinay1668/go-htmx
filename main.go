package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/vinay1668/go-htmx/templates"
	"github.com/vinay1668/go-htmx/templates/common"
)

func main() {
	e := echo.New()
	component := templates.Index()
	fmt.Println("router")
	e.GET("/", func(c echo.Context) error {
		return component.Render(context.Background(), c.Response())
	})
	e.GET("/login", func(c echo.Context) error {
		loginComponent := common.Login()
		return loginComponent.Render(context.Background(), c.Response())
	})
	e.GET("/signup", func(c echo.Context) error {
		signupComponent := common.Signup()
		return signupComponent.Render(context.Background(), c.Response())
	})
	e.POST("/api/login", handleLogin)
	e.POST("/api/signup", handleSignup)

	e.Static("/css", "css")
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":3000"))

}

func handleLogin(c echo.Context) error {
	fmt.Println("Login attempt")
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Here you would typically validate the credentials against a database
	// For this example, we'll use a simple check
	if email == "user@example.com" && password == "password" {
		return c.HTML(http.StatusOK, `<div id="login-response" class="text-green-500">Login successful!</div>`)
	} else {
		return c.HTML(http.StatusUnauthorized, `<div id="login-response" class="text-red-500">Invalid credentials</div>`)
	}
}
func handleSignup(c echo.Context) error {
	fmt.Println("Signup attempt")
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Here you would typically create a new user in your database
	// For this example, we'll use a simple check
	if username != "" && password != "" {
		return c.HTML(http.StatusOK, `<div id="signup-response" class="text-green-500">Signup successful!</div>`)
	} else {
		return c.HTML(http.StatusBadRequest, `<div id="signup-response" class="text-red-500">Invalid username or password</div>`)
	}
}
