package middlewares

import (
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

func DBMiddleware(app firebase.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", app)
		c.Next()
	}
}
