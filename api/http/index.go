package http

import (
	"auth-plus-notification/api/http/middlewares"
	"auth-plus-notification/api/http/routes"
	"auth-plus-notification/config"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Server() {
	env := config.GetEnv()
	router := gin.New()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.Metric())

	// Default
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
	})
	router.GET("/metrics", func(c *gin.Context) {
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	})

	// Application
	router.POST("/email", routes.EmailHandler)
	router.POST("/push_notification", routes.PushNotificationHandler)
	router.POST("/sms", routes.SmsHandler)

	router.Run(":" + env.App.Port)
}
