package handlers

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello world")
	})
}
