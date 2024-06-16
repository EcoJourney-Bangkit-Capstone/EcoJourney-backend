package routes

import (
	// "crypto-academy-backend/controller"
	// middlewares "crypto-academy-backend/middleware"

	"Ecojourney-backend/controller"
	middlewares "Ecojourney-backend/middleware"

	"github.com/gin-gonic/gin"
)

func ConfigureRouter(router *gin.Engine) {
	// Main Endpoint
	router.GET("/api", controller.MainController)
	apis := router.Group("/api")

	// Auth Endpoint
	apis.POST("/register", controller.Register)
	apis.POST("/login", controller.Login)

	// Check Endpoint
	apis.GET("/public", controller.PublicEndpoint)

	// Apply authentication middleware to protected routes
	protected := router.Group("/protected")
	protected.Use(middlewares.AuthMiddleware)
	protected.GET("/authenticated", controller.AuthenticatedEndpoint)

}
