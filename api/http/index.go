// Package http contain the function to start HTTP server
package http

import (
	"auth-plus-notification/api/http/middlewares"
	"auth-plus-notification/api/http/routes"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Server for initiate http server
func Server() *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.Metric())
	router.Use(middlewares.Trace())

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

	return router
}
