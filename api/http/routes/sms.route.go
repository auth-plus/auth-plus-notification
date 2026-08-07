package routes

import (
	core "auth-plus-notification/internal"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SmsRequestBody is type for payload
type SmsRequestBody struct {
	Phone   string `json:"phone"`
	Content string `json:"content"`
}

// SmsHandler ia route handler for POST /sms
func SmsHandler(c *gin.Context) {
	var reqBody SmsRequestBody

	if err := c.BindJSON(&reqBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	ctx := context.WithoutCancel(c.Request.Context())
	go core.NewCore().SmsUsecase.Send(
		ctx,
		reqBody.Phone,
		reqBody.Content)
	c.String(http.StatusOK, "Ok")
}
