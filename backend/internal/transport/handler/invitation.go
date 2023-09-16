package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	BIP_project "github.com/nekitalek/bip_project/backend"
)

// CreateInvitation godoc
//
//	@Summary		CreateInvitation
//	@Description	Создание приглашения
//	@Tags			invite
//	@Accept			json
//	@Produce		json
//	@Router			/invitation/CreateInvitation [post]
//	 @Success 200 {integer} integer 1
//	 @Failure 400 object} errorResponse
func (h *Handler) CreateInvitation(c *gin.Context) {
	input := new(BIP_project.Event_invitations_input)

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user_id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	id, err := h.services.Invitation.CreateInvitation(user_id, input)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	//костыль
	if *input.Status == BIP_project.Confirmed {
		err = h.services.PushNotification.SendPushNotification(*input.Event_id)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status":        "ok",
		"invitation_id": id,
	})
}

// GetInvitation godoc
//
//	@Summary		GetInvitation
//	@Description	Получение приглашения
//	@Tags			invite
//	@Accept			json
//	@Produce		json
//	@Router			/invitation/GetInvitation [get]
//	 @Success 200 {integer} integer 1
//	 @Failure 400 object} errorResponse
func (h *Handler) GetInvitation(c *gin.Context) {
	user_id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	inv, err := h.services.Invitation.GetInvitation(user_id)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	c.JSON(http.StatusOK, inv)
}

// UpdateInvitation godoc
//
//	@Summary		UpdateInvitation
//	@Description	Обновление статуса приглашения
//	@Tags			invite
//	@Accept			json
//	@Produce		json
//	@Router			/invitation/UpdateInvitation [patch]
//	 @Success 200 {integer} integer 1
//	 @Failure 400 object} errorResponse
func (h *Handler) UpdateInvitation(c *gin.Context) {
	input := new(BIP_project.Event_invitations_input)

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	eventId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	user_id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	err = h.services.Invitation.UpdateInvitation(user_id, eventId, input)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// DeleteInvitation godoc
//
//	@Summary		DeleteInvitation
//	@Description	Удаление приглашения
//	@Tags			invite
//	@Accept			json
//	@Produce		json
//	@Router			/invitation/DeleteInvitation [delete]
//	 @Success 200 {integer} integer 1
//	 @Failure 400 object} errorResponse
func (h *Handler) DeleteInvitation(c *gin.Context) {
	input := new(BIP_project.Event_invitations_input)

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user_id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	err = h.services.Invitation.DeleteInvitation(user_id, input)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
