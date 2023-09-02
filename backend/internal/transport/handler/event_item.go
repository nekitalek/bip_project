package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	BIP_project "github.com/nekitalek/bip_project/backend"
)

func (h *Handler) CreateEvent(c *gin.Context) {
	input := new(BIP_project.Event_items)

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user_id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	id, err := h.services.EventItem.CreateEvent(user_id, input)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "ok",
		"event_id": id,
	})
}

func (h *Handler) GetEvents(c *gin.Context) {
	input := new(BIP_project.Event_items_input)

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	events, err := h.services.EventItem.GetEvents(input)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, events)
}

func (h *Handler) UpdateEvent(c *gin.Context) {
	input := new(BIP_project.Event_items_input)

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	eventId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.EventItem.UpdateEvent(userId, eventId, input)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) DeleteEvent(c *gin.Context) {
	eventId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.EventItem.DeleteEvent(userId, eventId)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
