package controller

import (
    "Ecojourney-backend/config"
    "Ecojourney-backend/helper"
    "Ecojourney-backend/models"
    "net/http"
	"bytes"
    "encoding/json"
    "io/ioutil"

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

func UploadUserProfilePicture(c *gin.Context) {
    // Menerima file gambar dari form-data
    imageFile, err := c.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, helper.GenerateResponse(true, "Invalid image format", nil))
        return
    }

    // Membuka file gambar
    image, err := imageFile.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, "Failed to open image", nil))
        return
    }
    defer image.Close()

    // Membuat nama file unik menggunakan timestamp
    fileName := helper.GenerateUniqueFileName(imageFile.Filename)

    // Mengunggah gambar ke Google Cloud Storage (GCS)
    imageURL, err := config.UploadImageToGCS(image, fileName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, "Failed to upload image", nil))
        return
    }

    // Memperbarui URL gambar profil pengguna menggunakan fungsi UpdateUser
    req := models.User{PhotoURL: imageURL}
    jsonData, err := json.Marshal(req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, "Failed to marshal JSON", nil))
        return
    }

    c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonData))
    c.Request.Header.Set("Content-Type", "application/json")
    UpdateUser(c)
}
