package http

import (
	routes "auth-plus-notification/presentation/http/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server() {
	router := gin.New()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
	})
	// This handler will match /user/john but will not match /user/ or /user
	router.POST("/email", routes.EmailHandler)
	router.POST("/push_notification", routes.PushNotificationHandler)
	router.POST("/sms", routes.SmsHandler)

	router.Run(":5000")
}
