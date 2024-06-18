package routes

import (
	// "crypto-academy-backend/controller"
	// middlewares "crypto-academy-backend/middleware"

	"Ecojourney-backend/controller"
	middlewares "Ecojourney-backend/middleware"

	"github.com/gin-gonic/gin"
)

func ConfigureRouter(router *gin.Engine) {
	// Make sure every request should go from /api
	api := router.Group("/api")
	{
		// Auth Endpoint
		auth := api.Group("/auth")
		{
			auth.POST("/register", controller.Register)
			auth.POST("/login", controller.Login)
			auth.POST("/logout", middlewares.AuthMiddleware, controller.Logout)
			auth.POST("/forgot-password", controller.ForgotPassword)
		}

		// User Endpoint
		user := api.Group("/user")
		{
			user.GET("/self", middlewares.AuthMiddleware, controller.GetSelf)
			user.POST("/update", middlewares.AuthMiddleware, controller.UpdateUser)
		}

		// Article Endpoint
		articles := api.Group("/articles")
		{
			articles.POST("/create", middlewares.AuthMiddleware, controller.AddArticle)
		}

		// Public Endpoint
		api.GET("/public", controller.PublicEndpoint)

		// Protected Endpoint
		api.GET("/protected", middlewares.AuthMiddleware, controller.AuthenticatedEndpoint)
	}
}
