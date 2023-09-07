package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUser godoc
//
//	@Summary		GetUser
//	@Description	Получение информации о пользователе
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Router			/user/GetUser [get]
//	 @Success 200 {integer} integer 1
//	 @Failure 400 object} errorResponse
func (h *Handler) GetUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	if userId == 0 {
		userId, err = getUserId(c)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	user_data, err := h.services.User.GetUser(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user_data)

}

// func (h *Handler) UpdateUser(c *gin.Context) {

// }
