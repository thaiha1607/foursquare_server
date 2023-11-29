package main

import (
	"net/http"

	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo/v4"
	"github.com/thaiha1607/foursquare_server/common/env"
)

func main() {
	e := echo.New()
	if !env.IsProdEnv() {
		pprof.Register(e)
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
