package controller

import (
	"Ecojourney-backend/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainController(c *gin.Context) {
	c.JSON(http.StatusBadRequest, helper.GenerateResponse(true, "This is the main controller", nil))
}

func PublicEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateResponse(false, "This is a public endpoint", nil))
}

func AuthenticatedEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateResponse(false, "This is an authenticated endpoint", nil))
}
