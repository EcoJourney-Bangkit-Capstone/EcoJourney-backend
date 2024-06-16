package main

import (
	// Standard packages

	"log"
	"os"

	// Local packages

	"Ecojourney-backend/config"
	middlewares "Ecojourney-backend/middleware"
	routes "Ecojourney-backend/routes"
	utils "Ecojourney-backend/utils"

	// Third-party packages

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

var (
	firebaseApp     *firebase.App
	authClient      *auth.Client
	firestoreClient *firestore.Client
	storageClient   *storage.Client
)

func init() {
	utils.LoadEnvVariables()
}

func main() {
	if err := config.InitFirebase(); err != nil {
		log.Fatalf("firebase initialization error: %v", err)
	}

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
