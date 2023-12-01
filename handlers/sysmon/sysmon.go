package sysmon

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetMetrics() {
	metrics := echo.New()
	// Add/remove trailing slash to improve SEO
	metrics.Pre(middleware.RemoveTrailingSlash())
	// Hide banner
	metrics.HideBanner = true
	// Hide port
	metrics.HidePort = true
	// GET /metrics
	metrics.GET("/", echoprometheus.NewHandler())
	// Start metrics server
	if err := metrics.Start(":8085"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
