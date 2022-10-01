// Package routes contains all routes handler for GIN
package routes

import (
	core "auth-plus-notification/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

// EmailRequestBody is type for payload
type EmailRequestBody struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

// EmailHandler ia route handler for POST /email
func EmailHandler(c *gin.Context) {
	var requestBody EmailRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	go core.NewCore().EmailUsecase.Send(
		requestBody.Email,
		requestBody.Subject,
		requestBody.Content)
	c.String(http.StatusOK, "Ok")
}
