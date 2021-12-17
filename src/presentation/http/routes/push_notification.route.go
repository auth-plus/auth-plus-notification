package routes

import (
	core "auth-plus-notification/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PushNotificationRequestBody struct {
	deviceId string
	title    string
	content  string
}

func PushNotificationHandler(c *gin.Context) {
	var reqBody PushNotificationRequestBody

	if err := c.BindJSON(&reqBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	go core.NewCore().PushNotificationUsecase.Send(
		reqBody.deviceId,
		reqBody.title,
		reqBody.content)
	c.String(http.StatusOK, "Ok")
}
