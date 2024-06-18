package controller

import (
	"Ecojourney-backend/config"
	"Ecojourney-backend/helper"
	"Ecojourney-backend/models"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

func GetSelf(c *gin.Context) {
	user, err := config.AuthClient.GetUser(c, c.GetString("uid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, err.Error(), nil))
		return
	}

	c.JSON(200, helper.GenerateResponse(false, "Successfully get user", user))
}

func UpdateUser(c *gin.Context) {
	var req models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.GenerateResponse(true, err.Error(), nil))
		return
	}

	// Verify the request body with the user model so update only on unempty fields
	params := (&auth.UserToUpdate{})

	if req.Email != "" {
		params = params.Email(req.Email)
	}

	if req.DisplayName != "" {
		params = params.DisplayName(req.DisplayName)
	}

	if req.PhotoURL != "" {
		params = params.PhotoURL(req.PhotoURL)
	}

	userRecord, err := config.AuthClient.UpdateUser(c, c.GetString("uid"), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, err.Error(), nil))
		return
	}

	c.JSON(200, helper.GenerateResponse(false, "Successfully updated user", userRecord))
}
