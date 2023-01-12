// Package main is the mais file for starting http server or kafka or any trigger
package main

import (
	http "auth-plus-notification/api/http"
	"auth-plus-notification/api/http/middlewares"
	"auth-plus-notification/config"
)

func main() {
	env := config.GetEnv()
	middlewares.MetricSetup()
	http.Server().Run(":" + env.App.Port)
}
