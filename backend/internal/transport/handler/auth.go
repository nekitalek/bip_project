package handler

import (
	"net/http"

	"github.com/gorilla/csrf"
	BIP_project "github.com/nekitalek/bip_project/backend"

	"github.com/gin-gonic/gin"
)

// GetCSRF godoc
//
//	@Summary		CSRF
//	@Description	Получение CSRF токена
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Router			/CSRF [get]
//	 @Success 200 {integer} integer 1
func (h *Handler) GetCSRF(c *gin.Context) {
	token_CSRF := csrf.Token(c.Request)
	c.JSON(http.StatusOK, map[string]interface{}{
		"token_CSRF": token_CSRF,
	})
}

// signUp godoc
//
//	@Summary		signUp
//	@Description	Вывод страницы регистрации
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//
// @Param input body BIP_project.User_auth true "account info"
//
//		@Router			/auth/signUp [post]
//	 @Success 200 {integer} integer 1
//	 @Failure 400 object} errorResponse
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

// signUpSecondFactor godoc
//
//	@Summary		signUpSecondFactor
//	@Description	Второй фактор при регистрации
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//
// @Param input body BIP_project.Email_confirmation true "account info"
//
//	@Router			/auth/signUpSecondFactor [post]
//	 @Success 200 {integer} integer 1
//	 @Failure 400 object} errorResponse
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
//
//	@Summary		signInPass
//	@Description	Страница входа пользвателя
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//
// @Param input body LoginAndPass true "account info"
//
//	@Router			/auth/signInPass [post]
//	 @Success 200 {integer} integer 1
//	 @Failure 400 object} errorResponse
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
//
//	@Summary		signInSecondFactor
//	@Description	Страница вывода второго фактора при входе
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//
// @Param input body BIP_project.Email_confirmation true "account info"
//
//	@Router			/auth/signInSecondFactor [post]
//	 @Success 200 {integer} integer 1
//	 @Failure 400 object} errorResponse
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

// ResendCode godoc
//
//	@Summary		ResendCode
//	@Description	Повторная отправка кода второго фактора на почту
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//
// @Param input body BIP_project.Email_confirmation true "account info"
//
//	@Router			/auth/ResendCode [post]
//	 @Success 200 {integer} integer 1
//	 @Failure 400 object} errorResponse
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

// ChangePassFirstFactor godoc
//
//	@Summary		ChangePassFirstFactor
//	@Description	Проверка первого фактора при изменении пароля
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//
// @Param input body LoginAndPass true "account info"
//
//	@Router			/auth/ChangePassFirstFactor [post]
//	 @Success 200 {integer} integer 1
//	 @Failure 400 object} errorResponse
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

// ChangePassSecondFactor godoc
//
//	@Summary		ChangePassSecondFactor
//	@Description	Проверка и изменение второго фактора при измеинении пароля
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//
// @Param input body ChangePassStruct true "account info"
//
//	@Router			/auth/ChangePassSecondFactor [post]
//	 @Success 200 {integer} integer 1
//	 @Failure 400 object} errorResponse
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

// ChangeLoginFirstFactor godoc
//
//	@Summary		ChangeLoginFirstFactor
//	@Description	Проверка первого фактора при изменении логина
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//
// @Param input body LoginAndPass true "account info"
//
//	@Router			/auth/ChangeLoginFirstFactor [post]
//	 @Success 200 {integer} integer 1
//	 @Failure 400 object} errorResponse
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

// ChangeLoginSecondFactor godoc
//
//	@Summary		ChangeLoginSecondFactor
//	@Description	Проверка второго фактора при изменении логина
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//
// @Param input body ChangeLoginStruct true "account info"
//
//	@Router			/auth/ChangeLoginSecondFactor [post]
//	 @Success 200 {integer} integer 1
//	 @Failure 400 object} errorResponse
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

// VerificationNewEmail godoc
//
//	@Summary		VerificationNewEmail
//	@Description	Подтверждение новой почты
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//
// @Param input body BIP_project.Email_confirmation true "account info"
//
//	@Router			/auth/VerificationNewEmail [post]
//	 @Success 200 {integer} integer 1
//	 @Failure 400 object} errorResponse
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
