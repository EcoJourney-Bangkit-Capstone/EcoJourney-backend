package controller

import (
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

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

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, err.Error(), nil))
		return
	}

	params := (&auth.UserToCreate{}).
		Email(req.Email).
		Password(req.Password)

	userRecord, err := config.AuthClient.CreateUser(c, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, err.Error(), nil))
		return
	}

	// Store user information in Firestore
	_, err = config.FirestoreClient.Collection("users").Doc(userRecord.UID).Set(c, map[string]interface{}{
		"username": req.Username,
		"email":    req.Email,
		"password": string(hashedPassword),
	})
	if err != nil {
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

	// Get user data from Firestore
	doc, err := config.FirestoreClient.Collection("users").Where("email", "==", req.Email).Limit(1).Documents(c).Next()
	if err != nil {
		c.JSON(http.StatusUnauthorized, helper.GenerateResponse(true, "Invalid email or password", nil))
		return
	}

	userData := doc.Data()
	hashedPassword := userData["password"].(string)

	// Verify the password using bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, helper.GenerateResponse(true, "Invalid email or password", nil))
		return
	}

	// token, err := config.AuthClient.CustomToken(c, doc.Ref.ID)
	token, err := helper.ConvertCustomTokenToIDToken(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, "Failed to generate token", nil))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateResponse(false, "Successfully logged in", gin.H{"token": token, "user": gin.H{
		"uid": doc.Ref.ID,
	}}))
}

func Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, helper.GenerateResponse(true, "Missing token", nil))
		return
	}

	if err := config.AuthClient.RevokeRefreshTokens(c, token); err != nil {
		c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, "Failed to revoke token", nil))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateResponse(false, "Successfully logged out", nil))
}
