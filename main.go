package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("sasamuku-echo-test"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigDebugLogger(os.Stdout),
		newrelic.ConfigAppLogEnabled(false),
		newrelic.ConfigCodeLevelMetricsEnabled(true),
	)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	e := echo.New()

	e.Use(nrecho.Middleware(app))

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Start(":8000")
}
