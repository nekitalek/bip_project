package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/nekitalek/bip_project/backend/docs"
	"github.com/nekitalek/bip_project/backend/internal/service"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Handler struct {
	services *service.Service
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "https://51.250.24.31")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("X-Frame-Options", "SAMEORIGIN")
		c.Header("Strict-Transport-Security", "max-age=2592000; includeSubDomains")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(CORSMiddleware())
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
		user := api.Group("user")
		{
			//user.POST("/", )
			user.GET("/:id", h.GetUser)
			//user.PATCH("/:id", h.UpdateUser)
			//user.DELETE("/:id",)
		}
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
		push := api.Group("push_notification")
		{
			push.POST("/", h.CreatePushNotification)
			push.DELETE("/", h.DeletePushNotification)
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
