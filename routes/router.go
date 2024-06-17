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
		}

		// Public Endpoint
		api.GET("/public", controller.PublicEndpoint)

		// Protected Endpoint
		api.GET("/protected", middlewares.AuthMiddleware, controller.AuthenticatedEndpoint)
	}
}
