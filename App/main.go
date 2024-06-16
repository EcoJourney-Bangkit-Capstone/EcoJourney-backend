package main

import (
	// Standard packages
	"context"
	"os"

	// Local packages
	config "Ecojourney-backend/config"
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
	ctx := context.Background()
	/*
	* Initialize Client
	 */
	firebase, err := config.InitFirebaseApp(ctx)
	if err != nil {
		panic(err)
	}

	/**
	 * Initialize Application
	 */
	app := gin.Default()
	app.Use(middlewares.CorsMiddleware())
	app.Use(middlewares.DBMiddleware(*firebase))

	routes.ConfigureRouter(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Run(":" + port)
}
