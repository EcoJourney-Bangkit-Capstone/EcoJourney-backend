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

	// Auth Endpoint
	api.POST("/register", controller.Register)
	api.POST("/login", controller.Login)

	// Public Endpoint
	api.GET("/public", controller.PublicEndpoint)

	// Protected Endpoint
	api.GET("/protected", middlewares.AuthMiddleware, controller.AuthenticatedEndpoint)

}
