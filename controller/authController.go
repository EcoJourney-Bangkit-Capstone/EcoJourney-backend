package controller

import (
	"Ecojourney-backend/config"
	"Ecojourney-backend/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	params := (&auth.UserToCreate{}).
		Email(req.Email).
		Password(req.Password)

	userRecord, err := config.AuthClient.CreateUser(c, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Store user information in Firestore
	_, err = config.FirestoreClient.Collection("users").Doc(userRecord.UID).Set(c, map[string]interface{}{
		"username": req.Username,
		"email":    req.Email,
		"password": string(hashedPassword),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to store user data"})
		return
	}

	token, err := config.AuthClient.CustomToken(c, userRecord.UID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"uid": userRecord.UID, "token": token})
}

func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user data from Firestore
	doc, err := config.FirestoreClient.Collection("users").Where("email", "==", req.Email).Limit(1).Documents(c).Next()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	userData := doc.Data()
	hashedPassword := userData["password"].(string)

	// Verify the password using bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// token, err := config.AuthClient.CustomToken(c, doc.Ref.ID)
	token, err := convertCustomTokenToIDToken(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func convertCustomTokenToIDToken(email string, password string) (string, error) {
	url := "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=" + os.Getenv("FIREBASE_API_KEY")
	payload := strings.NewReader(fmt.Sprintf(`{"email":"%s","password":"%s","returnSecureToken":true}`, email, password))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", errors.New(response["error"].(map[string]interface{})["message"].(string))
	}

	return response["idToken"].(string), nil
}
