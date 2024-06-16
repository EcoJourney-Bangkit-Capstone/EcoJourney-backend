package main

import (
	// Standard packages
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
	/*
	* Initialize Firebase
	 */
	db, err := config.Connect()
	if err != nil {
		panic(err)
	}
	defer os.Exit(0)

	/**
	 * Initialize Gin
	 */
	app := gin.Default()
	app.Use(middlewares.CorsMiddleware())
	app.Use(middlewares.DBMiddleware(*db))

	routes.ConfigureRouter(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Run(":" + port)
}
