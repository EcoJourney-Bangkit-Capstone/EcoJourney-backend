package controller

import (
	"log"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"

	"Ecojourney-backend/config"
	"Ecojourney-backend/helper"
	"Ecojourney-backend/models"
)

func Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	params := (&auth.UserToCreate{}).
		Email(req.Email).
		Password(req.Password).
		DisplayName(req.Username)

	userRecord, err := config.AuthClient.CreateUser(c, params)
	if err != nil {
		if err.Error() == "INVALID_LOGIN_CREDENTIALS" {
			c.JSON(http.StatusUnauthorized, helper.GenerateResponse(true, err.Error(), nil))
			return
		}
		c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, err.Error(), nil))
		return
	}

	token, err := config.AuthClient.CustomToken(c, userRecord.UID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, err.Error(), nil))
		return
	}

	c.JSON(http.StatusCreated, helper.GenerateResponse(false, "User registered successfully", gin.H{"token": token, "user": gin.H{
		"uid": userRecord.UID,
	}}))
}

func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.GenerateResponse(true, err.Error(), nil))
		return
	}

	token, err := helper.ConvertCustomTokenToIDToken(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateResponse(false, "Successfully logged in", gin.H{"token": token}))
}

func Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, helper.GenerateResponse(true, "Authorization token is required", nil))
		return
	}

	// Remove bearer
	token = token[7:]

	if err := config.AuthClient.RevokeRefreshTokens(c, c.GetString("uid")); err != nil {
		c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, err.Error(), nil))
		return
	}

	_, err := config.AuthClient.VerifyIDTokenAndCheckRevoked(c, token)
	if err != nil {
		if err.Error() == "ID token has been revoked" {
			// Token is revoked. Change the existing uid token to invalid
			user, err := config.AuthClient.GetUser(c, c.GetString("uid"))
			if err != nil {
				c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, err.Error(), nil))
				return
			}

			c.JSON(http.StatusOK, helper.GenerateResponse(false, "Successfully logged out", gin.H{"user": gin.H{
				"uid":      user.UID,
				"email":    user.Email,
				"username": user.DisplayName,
			}}))

		} else {
			// Token is invalid
			c.JSON(http.StatusBadRequest, helper.GenerateResponse(true, "Invalid token", nil))
		}
	}
}

func ForgotPassword(c *gin.Context) {
	var req models.ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.GenerateResponse(true, err.Error(), nil))
		return
	}

	link, err := config.AuthClient.PasswordResetLinkWithSettings(c, req.Email, nil)
	if err != nil {
		log.Fatalf("error generating email link: %v\n", err)
		c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, err.Error(), nil))
	}

	c.JSON(http.StatusOK, helper.GenerateResponse(false, "Password reset link generated!", gin.H{"link": link}))
}
