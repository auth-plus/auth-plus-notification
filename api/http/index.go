package http

import (
	"auth-plus-notification/api/http/routes"
	"auth-plus-notification/config"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Server() {
	env := config.GetEnv()
	router := gin.New()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Default
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
	})
	// This handler will match /user/john but will not match /user/ or /user
	router.POST("/email", routes.EmailHandler)
	router.POST("/push_notification", routes.PushNotificationHandler)
	router.POST("/sms", routes.SmsHandler)

	router.Run(":" + env.App.Port)
}
