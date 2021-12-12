package http

import (
	email "auth-plus-notification/presentation/http/routes"
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
	router.POST("/email", email.EmailHandler)

	router.Run(":5000")
}
