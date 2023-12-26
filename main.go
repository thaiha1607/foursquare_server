package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thaiha1607/foursquare_server/common/env"
	"github.com/thaiha1607/foursquare_server/handlers"
	"github.com/thaiha1607/foursquare_server/handlers/sysmon"
	"github.com/thaiha1607/foursquare_server/middlewares"
)

func main() {
	// Check whether we are running in production environment
	isProdEnv := env.IsProdEnv()
	e := echo.New()
	// Add/remove trailing slash to improve SEO
	e.Pre(middleware.RemoveTrailingSlash())
	// Hide banner
	e.HideBanner = true
	// Hide port
	e.HidePort = true
	// Set debug mode
	e.Debug = !isProdEnv
	// Request ID
	e.Use(middleware.RequestID())
	// Logging
	e.Use(middlewares.LoggerMiddleware())
	// Recover
	e.Use(middleware.Recover())
	// Profiler
	if !isProdEnv {
		pprof.Register(e)
	}
	// System monitor
	e.Use(echoprometheus.NewMiddleware("foursquare_server"))
	go sysmon.GetMetrics()
	// Rate limiter
	// Limit to 5000 requests per second
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(5000)))
	// Request timeout
	// Timeout after 10 seconds
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 10 * time.Second,
	}))
	// CORS
	e.Use(middleware.CORS())
	// Secure middleware
	e.Use(middleware.Secure())
	// Compress
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		// Minimum content length to trigger gzip
		// We set it to 1 MB
		MinLength: 1024 * 1024,
		// Gzip compression level
		// It should be 5 or 6 for the best size/efficiency tradeoff
		Level: 5,
		// Avoid gzip compression for certain paths
		Skipper: func(c echo.Context) bool {
			isProfiler := strings.Contains(c.Path(), "debug/pprof")
			return isProfiler
		},
	},
	))
	// Decompress HTTP request if Content-Encoding header is set to gzip
	e.Use(middleware.Decompress())
	// Routes
	handlers.RegisterRoutes(e)
	go func() {
		if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	// Graceful shutdown
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
