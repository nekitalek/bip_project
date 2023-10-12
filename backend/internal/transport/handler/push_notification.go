package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	BIP_project "github.com/nekitalek/bip_project/backend"
)

func (h *Handler) CreatePushNotification(c *gin.Context) {
	var input BIP_project.Push_notification_input

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user_id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	err = h.services.PushNotification.CreatePushNotification(user_id, input.Token, input.Device)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) DeletePushNotification(c *gin.Context) {
	var input BIP_project.Push_notification_input
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user_id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	err = h.services.PushNotification.DeletePushNotification(user_id, input.Token)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
