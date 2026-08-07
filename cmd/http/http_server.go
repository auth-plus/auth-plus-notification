// Package main is the mais file for starting http server or kafka or any trigger
package main

import (
	"context"

	http "auth-plus-notification/api/http"
	"auth-plus-notification/api/http/middlewares"
	"auth-plus-notification/config"

	"github.com/uptrace/uptrace-go/uptrace"
)

func main() {
	env := config.GetEnv()

	ctx := context.Background()
	uptrace.ConfigureOpentelemetry(
		uptrace.WithDSN(env.Trace.DSN),
		uptrace.WithServiceName(env.App.Name),
		uptrace.WithDeploymentEnvironment(env.App.Env),
	)
	defer uptrace.Shutdown(ctx)

	middlewares.MetricSetup()
	http.Server().Run(":" + env.App.Port)
}
