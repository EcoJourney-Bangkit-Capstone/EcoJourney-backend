package controller

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"Ecojourney-backend/config"
	"Ecojourney-backend/helper"
)

type ResponseBody struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Details struct {
		TotalCount int         `json:"Total_count"`
		Articles   []helper.Article `json:"articles"`
		HistoryId  string      `json:"historyId"`
	} `json:"details"`
}

func WasteRecognitionHandler(c *gin.Context) {
	imageFile, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Invalid image format", "data": nil})
		return
	}

	image, err := imageFile.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to open image", "data": nil})
		return
	}
	defer image.Close()

	// Generate unique file name
	fileName := time.Now().Format("20060102150405") + filepath.Ext(imageFile.Filename)
	imageURL, err := config.UploadImageToGCS(image, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to upload image", "data": nil})
		return
	}

	uid := c.GetString("uid")
	wasteTypes := c.PostFormArray("type")

	// Save the history
	historyId, err := helper.SaveHistory(uid, imageURL, wasteTypes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to save history", "data": nil})
		return
	}

	// Create response
	response := gin.H{
		"total_type": len(wasteTypes),
		"type":       wasteTypes,
		"historyId":  historyId,
		"imageURL":   imageURL,
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Image analyzed successfully", "details": response})
}




func WasteHistoryHandler(c *gin.Context) {
	uid := c.GetString("uid")
	if uid == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Unauthorized", "details": nil})
		return
	}

	histories, err := helper.GetWasteHistory(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to retrieve history", "details": nil})
		return
	}

	response := gin.H{
		"total_item": len(histories),
		"history":    histories,
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "History retrieved successfully", "details": response})
}