package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	exporter "github.com/mingcheng/prometheus-smart-exporter"
)

func main() {
	// Echo instance
	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		context.Response().Header().Add("Location", "/metrics")
		return context.NoContent(http.StatusFound)
	})

	// Routes
	e.GET("/metrics", func(context echo.Context) error {
		output, err := exporter.RunScript()
		if err != nil {
			return context.NoContent(http.StatusInternalServerError)
		} else {
			return context.String(http.StatusOK, output)
		}
	})

	// Start server
	_ = e.Start(":9111")
}
