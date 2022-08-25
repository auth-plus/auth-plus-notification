package routes

import (
	core "auth-plus-notification/cmd"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmailRequestBody struct {
	Email   string `json:"email"`
	Content string `json:"content"`
}

func EmailHandler(c *gin.Context) {
	var requestBody EmailRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	go core.NewCore().EmailUsecase.Send(
		requestBody.Email,
		requestBody.Content)
	c.String(http.StatusOK, "Ok")
}
