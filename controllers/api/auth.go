package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func HandleLogin(c echo.Context) error {
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

func HandleSignup(c echo.Context) error {
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
