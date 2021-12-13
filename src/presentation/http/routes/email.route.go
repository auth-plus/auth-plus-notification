package routes

import (
	core "auth-plus-notification/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmailRequestBody struct {
	email   string
	content string
}

func EmailHandler(c *gin.Context) {
	var requestBody EmailRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	co := core.NewCore()
	co.EmailUsecase.Send(requestBody.email, requestBody.content)
	c.String(http.StatusOK, "Ok")
}
