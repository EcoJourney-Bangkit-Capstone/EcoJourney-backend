package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainController(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  "error",
		"message": "Endpoint not found",
	})
}

func PublicEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "This is a public 2.0 endpoint",
	})
}

func AuthenticatedEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "This is an authenticated endpoint",
	})
}
