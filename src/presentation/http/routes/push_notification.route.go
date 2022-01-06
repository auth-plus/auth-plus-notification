package routes

import (
	core "auth-plus-notification/core"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type PushNotificationRequestBody struct {
	DeviceId string `json:"deviceId"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

func PushNotificationHandler(c *gin.Context) {
	reqBody := PushNotificationRequestBody{}

	if err := c.ShouldBindBodyWith(&reqBody, binding.JSON); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	go core.NewCore().PushNotificationUsecase.Send(
		reqBody.DeviceId,
		reqBody.Title,
		reqBody.Content)
	c.String(http.StatusOK, "Ok")
}
