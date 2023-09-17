package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func getAnimal(c echo.Context) error {
	txn := nrecho.FromContext(c)
	txn.SetOption(newrelic.WithThisCodeLocation())
	defer txn.StartSegment("getAnimal").End()

	return c.String(http.StatusOK, "panda")
}

func hello(c echo.Context) error {
	txn := nrecho.FromContext(c)
	txn.SetOption(newrelic.WithThisCodeLocation())
	defer txn.StartSegment("hello").End()

	return c.String(http.StatusOK, "hello world")
}

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
	e.Debug = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(nrecho.Middleware(app))

	e.GET("/hello", hello)

	g := e.Group("/animals")
	g.GET("/:id", getAnimal)

	e.Start(":8000")
}
