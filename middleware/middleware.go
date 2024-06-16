package middlewares

import (
	"Ecojourney-backend/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	idToken := c.GetHeader("Authorization")
	if idToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization token required"})
		c.Abort()
		return
	}

	// Strip the Bearer prefix
	idToken = idToken[7:]
	println(idToken)

	// Verify the ID token
	token, err := config.AuthClient.VerifyIDToken(c, idToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}

	c.Set("uid", token.UID)
	c.Next()
}
