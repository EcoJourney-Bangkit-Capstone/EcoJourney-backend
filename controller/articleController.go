package controller

import (
	"net/http"
	// "path/filepath"
	// "time"
	

	"Ecojourney-backend/config"
	"Ecojourney-backend/helper"
	"Ecojourney-backend/models"
	"github.com/gin-gonic/gin"
	"cloud.google.com/go/firestore" 
	"google.golang.org/api/iterator"
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

func EditArticle(c *gin.Context) {
	articleID := c.Param("articleId")
	if articleID == "" {
		c.JSON(http.StatusBadRequest, helper.GenerateResponse(true, "Invalid article ID", nil))
		return
	}

	var req models.ArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Validation error", "details": err.Error()})
		return
	}

	// Validate required fields
	if req.Title == "" || req.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Validation error",
			"details": gin.H{
				"title":   "Title is required",
				"content": "Content is required",
			},
		})
		return
	}

	// Reference to the article document
	docRef := config.FirestoreClient.Collection("articles").Doc(articleID)
	_, err := docRef.Get(c)
	if err != nil {
		c.JSON(http.StatusNotFound, helper.GenerateResponse(true, "Article not found", nil))
		return
	}

	_, err = docRef.Set(c, map[string]interface{}{
		"title":   req.Title,
		"content": req.Content,
		"topic":   req.Topic,
	}, firestore.MergeAll)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Article updated successfully",
		"details": gin.H{
			"articleId": articleID,
		},
	})
}


func GetArticles(c *gin.Context) {
	iter := config.FirestoreClient.Collection("articles").Documents(c)
	var articles []models.ArticleResponse
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, err.Error(), nil))
			return
		}

		var article models.ArticleResponse
		doc.DataTo(&article)
		article.ID = doc.Ref.ID
		articles = append(articles, article)
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Articles retrieved successfully",
		"details": gin.H{
			"Total_count": len(articles),
			"articles":    articles,
		},
	})
}

func WasteRecognitionHistoryHandler(c *gin.Context) {
	uid := c.GetString("uid")
	history, err := helper.GetWasteHistory(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, "Failed to retrieve history", nil))
		return
	}

	response := gin.H{
		"Total_item": len(history),
		"history":    history,
	}

	c.JSON(http.StatusOK, helper.GenerateResponse(false, "History retrieved successfully", response))
}

func WasteRecognitionHandler(c *gin.Context) {
	// Parse the form
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, helper.GenerateResponse(true, "Unable to parse form", nil))
		return
	}

	// Get the waste type
	wasteType := c.PostForm("type")
	if wasteType == "" {
		c.JSON(http.StatusBadRequest, helper.GenerateResponse(true, "Waste type is required", nil))
		return
	}

	// Fetch articles related to waste type
	articles, err := helper.FetchArticlesByKeyword(wasteType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, "Failed to fetch articles", nil))
		return
	}

	// Get user ID from context
	uid := c.GetString("uid")

	// Here you can create a mock URL since we're skipping the upload step
	imageURL := "mock_image_url"

	// Store the history
	historyId, err := helper.SaveHistory(uid, imageURL, wasteType, articles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.GenerateResponse(true, "Failed to save history", nil))
		return
	}

	// Create response
	response := gin.H{
		"image_url": imageURL,
		"articles":  articles,
		"historyId": historyId,
	}

	c.JSON(http.StatusOK, helper.GenerateResponse(false, "Image analyzed successfully", response))
}