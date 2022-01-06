package routes

import (
	core "auth-plus-notification/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SmsRequestBody struct {
	Phone   string `json:"phone"`
	Content string `json:"content"`
}

func SmsHandler(c *gin.Context) {
	var reqBody SmsRequestBody

	if err := c.BindJSON(&reqBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	go core.NewCore().SmsUsecase.Send(
		reqBody.Phone,
		reqBody.Content)
	c.String(http.StatusOK, "Ok")
}
