package handler

import (
	"net/http"

	"github.com/gorilla/csrf"
	BIP_project "github.com/nekitalek/bip_project/backend"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCSRF(c *gin.Context) {
	token_CSRF := csrf.Token(c.Request)
	c.JSON(http.StatusOK, map[string]interface{}{
		"token_CSRF": token_CSRF,
	})
}

// signUp godoc
// @Summary Вывод стандартной страницы
// @Description Вывод стандартной страницы
// @Tags signUp
// @Accept  json
// @Produce  json
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {

	var input BIP_project.User_auth

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	user_id, e_conf_id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":   user_id,
		"e_conf_id": e_conf_id,
		"status":    "ok",
	})
}

func (h *Handler) signUpSecondFactor(c *gin.Context) {
	var input BIP_project.Email_confirmation
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	err := h.services.Authorization.SignUpSecondFactor(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})

}

type LoginAndPass struct {
	Login    string `json:"Login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// signInPass godoc
// @Summary Вывод стандартной страницы
// @Description Вывод стандартной страницы
// @Tags signInPass
// @Accept  json
// @Produce  json
// @Router /auth/sign-in/pass [post]
func (h *Handler) signInPass(c *gin.Context) {
	var input LoginAndPass

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user_id, e_conf_id, err := h.services.Authorization.SingInByPass(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":   user_id,
		"e_conf_id": e_conf_id,
		"status":    "ok",
	})
}

type signInInputSecondFactor struct {
	Login    string `json:"Login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Code     int    `json:"Code" binding:"required"`
}

// signInSecondFactor godoc
// @Summary Вывод стандартной страницы
// @Description Вывод стандартной страницы
// @Tags signInSecondFactor
// @Accept  json
// @Produce  json
// @Router /auth/sign-in/factor [post]
func (h *Handler) signInSecondFactor(c *gin.Context) {
	var input BIP_project.Email_confirmation

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token_JWT, err := h.services.Authorization.SignInSecondFactor(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"auth_token": token_JWT,
	})
}

func (h *Handler) ResendCode(c *gin.Context) {
	var input BIP_project.Email_confirmation
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Authorization.ReSendCode(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) ChangePassFirstFactor(c *gin.Context) {
	var input LoginAndPass
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user_id, e_conf_id, err := h.services.Authorization.ChangePassFirstFactor(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":   user_id,
		"e_conf_id": e_conf_id,
		"status":    "ok",
	})
}

type ChangePassStruct struct {
	E_conf      BIP_project.Email_confirmation `json:"e_conf"`
	NewPassword string                         `json:"new_password"`
}

func (h *Handler) ChangePassSecondFactor(c *gin.Context) {
	var input ChangePassStruct

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Authorization.ChangePassSecondFactor(input.E_conf, input.NewPassword)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) ChangeLoginFirstFactor(c *gin.Context) {
	var input LoginAndPass
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user_id, e_conf_id, err := h.services.Authorization.ChangeLoginFirstFactor(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":   user_id,
		"e_conf_id": e_conf_id,
		"status":    "ok",
	})
}

type ChangeLoginStruct struct {
	E_conf   BIP_project.Email_confirmation `json:"e_conf"`
	NewLogin string                         `json:"new_login" binding:"required"`
}

func (h *Handler) ChangeLoginSecondFactor(c *gin.Context) {
	var input ChangeLoginStruct

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user_id, e_conf_id, err := h.services.Authorization.ChangeLoginSecondFactor(input.E_conf, input.NewLogin)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":   user_id,
		"e_conf_id": e_conf_id,
		"status":    "ok",
	})
}

func (h *Handler) VerificationNewEmail(c *gin.Context) {
	var input BIP_project.Email_confirmation

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Authorization.VerificationNewEmail(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
