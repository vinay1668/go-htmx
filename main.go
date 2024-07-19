package main

import (
	"github.com/labstack/echo"
	"github.com/vinay1668/go-htmx/routes"
)

func main() {
	e := echo.New()

	// Register all routes
	routes.RegisterRoutes(e)

	// Static files
	e.Static("/css", "css")
	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":3000"))

}
