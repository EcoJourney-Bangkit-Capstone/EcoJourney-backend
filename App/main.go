package main

import (
	// Standard packages

	"os"

	// Local packages

	middlewares "Ecojourney-backend/middleware"
	routes "Ecojourney-backend/routes"
	utils "Ecojourney-backend/utils"

	// Third-party packages
	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadEnvVariables()
}

func main() {
	/*
	* Initialize Client
	 */

	/**
	 * Initialize Application
	 */
	app := gin.Default()
	app.Use(middlewares.CorsMiddleware())

	routes.ConfigureRouter(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Run(":" + port)
}
