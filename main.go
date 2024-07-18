package main

import (
	"context"

	"github.com/labstack/echo"
	"github.com/vinay1668/go-htmx/templates"
)

func main() {
	e := echo.New()
	component := templates.Index()
	e.GET("/", func(c echo.Context) error {
		return component.Render(context.Background(), c.Response())
	})
	e.Static("/css", "css")
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":3000"))

}
