package routes

import (
	// "crypto-academy-backend/controller"
	// middlewares "crypto-academy-backend/middleware"

	"Ecojourney-backend/controller"

	"github.com/gin-gonic/gin"
)

func ConfigureRouter(router *gin.Engine) {
	// Main Endpoint
	router.GET("/api", controller.MainController)
	apis := router.Group("/api")

	// Auth Endpoint

	// Check Endpoint
	apis.GET("/public", controller.PublicEndpoint)
	apis.GET("/authenticated", controller.AuthenticatedEndpoint)

}
