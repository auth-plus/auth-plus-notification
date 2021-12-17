package routes

import (
	core "auth-plus-notification/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SmsRequestBody struct {
	phone   string
	content string
}

func SmsHandler(c *gin.Context) {
	var reqBody SmsRequestBody

	if err := c.BindJSON(&reqBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	go core.NewCore().SmsUsecase.Send(
		reqBody.phone,
		reqBody.content)
	c.String(http.StatusOK, "Ok")
}
