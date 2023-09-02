package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nekitalek/bip_project/backend/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(h.CSRFIdentity())
	router.GET("/CSRF", h.GetCSRF)

	auth := router.Group("/auth")
	{
		sign_up := auth.Group("/sign-up")
		{
			sign_up.POST("/password", h.signUp)
			sign_up.POST("/sec_factor", h.signUpSecondFactor)
		}

		//auth.POST("/sign-up", h.signUp)
		email := auth.Group("/email")
		{
			//email.POST("/verification", h.EmailVerification)
			email.POST("/resend", h.ResendCode)
		}

		sign_in := auth.Group("/sign-in")
		{
			sign_in.POST("/password", h.signInPass)
			sign_in.POST("/sec_factor", h.signInSecondFactor)
		}
		//CSRF := csrf.Protect([]byte("a-32-byte-long-key-goes-here"))
		//csrf.MaxAge(0))
		//sign_in.Use(adapter.Wrap(CSRF))

		change := auth.Group("/change")
		{
			password := change.Group("/password")
			{
				password.POST("/first_factor", h.ChangePassFirstFactor)
				password.POST("/sec_factor", h.ChangePassSecondFactor)
			}
			email := change.Group("/email")
			{
				email.POST("/first_factor", h.ChangeLoginFirstFactor)
				email.POST("/sec_factor", h.ChangeLoginSecondFactor)
				email.POST("/verification_new_email", h.VerificationNewEmail)
			}

		}
	}

	api := router.Group("/api", h.userIdentity)
	{
		event := api.Group("event")
		{
			event.POST("/", h.CreateEvent)
			event.GET("/", h.GetEvents)
			//может только admin события
			event.PATCH("/:id", h.UpdateEvent)
			//может только admin события
			event.DELETE("/:id", h.DeleteEvent)
		}
		invitation := api.Group("invitation")
		{
			invitation.POST("/", h.CreateInvitation)
			invitation.GET("/", h.GetInvitation)
			invitation.PATCH("/:id", h.UpdateInvitation)
			invitation.DELETE("/", h.DeleteInvitation)
		}

		// lists := api.Group("/lists")
		// {
		// 	lists.POST("/", h.createList)
		// 	lists.GET("/", h.getAllLists)
		// 	lists.GET("/:id", h.getListById)
		// 	lists.PUT("/:id", h.updateList)
		// 	lists.DELETE("/:id", h.deleteList)
		// }

		// 	items := lists.Group(":id/items")
		// 	{
		// 		items.POST("/", h.createItem)
		// 		items.GET("/", h.getAllItems)
		// 	}
		// }

	}

	return router
}
