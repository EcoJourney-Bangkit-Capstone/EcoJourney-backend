package main

import (
	middlewares "Ecojourney-backend/app/middleware"
	routes "Ecojourney-backend/app/routes"

	"context"
	"log"

	// config "Ecojourney-backend/config"

	// "Ecojourney-backend/config"
	// "log"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}

	// Create a new gin router
	app := gin.Default()
	app.Use(middlewares.CorsMiddleware())
	app.Use(dbMiddleware(*db))

	routes.ConfigureRouter(app)
	app.Run(":8000")
}

func connectDB() (c *firebase.App, err error) {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "ecojourney"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer client.Close()

	return app, err
}

// create a middleware that attach the db connection to the context
func dbMiddleware(app firebase.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", app)
		c.Next()
	}
}
