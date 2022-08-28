package main

import (
	http "auth-plus-notification/api/http"
	"auth-plus-notification/config"
)

func main() {
	env := config.GetEnv()
	http.Server().Run(":" + env.App.Port)
}
