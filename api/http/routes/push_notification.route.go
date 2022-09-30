package routes

import (
	core "auth-plus-notification/internal"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// PushNotificationRequestBody is type for payload
type PushNotificationRequestBody struct {
	DeviceID string `json:"deviceId"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

// PushNotificationHandler ia route handler for POST /push_notification
func PushNotificationHandler(c *gin.Context) {
	reqBody := PushNotificationRequestBody{}

	if err := c.ShouldBindBodyWith(&reqBody, binding.JSON); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	go core.NewCore().PushNotificationUsecase.Send(
		reqBody.DeviceID,
		reqBody.Title,
		reqBody.Content)
	c.String(http.StatusOK, "Ok")
}
