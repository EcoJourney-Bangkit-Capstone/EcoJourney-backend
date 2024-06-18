package controller

import (
	"net/http"
	"Ecojourney-backend/config"
	"Ecojourney-backend/helper"
	"Ecojourney-backend/models"
	"github.com/gin-gonic/gin"
)

func AddArticle(c *gin.Context) {
	var req models.ArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Validation error", "details": err.Error()})
		return
	}

	// Validate required fields
	if req.Title == "" || req.Content == "" || req.AuthorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Validation error",
			"details": gin.H{
				"title":    "Title is required",
				"content":  "Content is required",
				"authorId": "Author ID is required",
			},
		})
		return
	}

	// Create a new article in Firestore
	docRef, _, err := config.FirestoreClient.Collection("articles").Add(c, map[string]interface{}{
		"title":    req.Title,
		"content":  req.Content,
		"authorId": req.AuthorID,
		"topic":    req.Topic,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, err.Error(), nil))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"error":   false,
		"message": "Article created successfully",
		"details": gin.H{
			"articleId": docRef.ID,
		},
	})
}

func DeleteArticle(c *gin.Context) {
	articleID := c.Param("articleId")
	if articleID == "" {
		c.JSON(http.StatusBadRequest, helper.GenerateResponse(true, "Invalid article ID", nil))
		return
	}

	// Reference to the article document
	docRef := config.FirestoreClient.Collection("articles").Doc(articleID)
	_, err := docRef.Get(c)
	if err != nil {
		c.JSON(http.StatusNotFound, helper.GenerateResponse(true, "Article not found", nil))
		return
	}

	_, err = docRef.Delete(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateResponse(false, "Article deleted successfully", nil))
}