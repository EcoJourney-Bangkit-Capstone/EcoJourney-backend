package helper

import "github.com/gin-gonic/gin"

// Helper function to create payload template
func GenerateResponse(err bool, message string, data any) gin.H {
	return gin.H{
		"error":   err,
		"message": message,
		"data":    data,
	}
}
