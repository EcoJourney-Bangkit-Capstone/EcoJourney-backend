package controller

import (
	"Ecojourney-backend/config"
	"Ecojourney-backend/helper"

	"github.com/gin-gonic/gin"
)

func GetSelf(c *gin.Context) {
	// config.AuthClient.GetUser(c, c.GetString("uid"))
	user, err := config.AuthClient.GetUser(c, c.GetString("uid"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, helper.GenerateResponse(false, "Successfully get user", user))
}
