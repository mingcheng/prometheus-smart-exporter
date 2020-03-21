package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	exporter "github.com/mingcheng/prometheus-smart-exporter"
	flag "github.com/spf13/pflag"
)

var (
	version   string
	buildTime string
)

func main() {
	printVersion := flag.Bool("version", false, "help message for flagname")
	flag.Parse()

	if *printVersion {
		fmt.Printf("Version %s, Build on %s\n", version, buildTime)
		return
	}

	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		context.Response().Header().Add("Location", "/metrics")
		return context.NoContent(http.StatusFound)
	})

	e.GET("/metrics", func(context echo.Context) error {
		output, err := exporter.RunScript()
		if err != nil {
			return context.NoContent(http.StatusInternalServerError)
		} else {
			return context.String(http.StatusOK, output)
		}
	})

	e.GET("/version", func(context echo.Context) error {
		return context.String(http.StatusOK, fmt.Sprintf("Version %s, Build on %s", version, buildTime))
	})

	e.GET("/script", func(context echo.Context) error {
		if script, err := exporter.GetScript(); err != nil {
			return context.NoContent(http.StatusNotFound)
		} else {
			return context.String(http.StatusOK, script)
		}
	})

	_ = e.Start(":9111")
}
