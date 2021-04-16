package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/users/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/:id")
	})

	e.GET("/users/new", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/new")
	})

	e.GET("/users/files/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/files/*")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
