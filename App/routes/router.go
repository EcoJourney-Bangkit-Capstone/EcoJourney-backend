package routes

import (
	// "crypto-academy-backend/controller"
	// middlewares "crypto-academy-backend/middleware"

	"github.com/gin-gonic/gin"
)

func ConfigureRouter(router *gin.Engine) {
	// General Endpoint
	router.GET("/", func(ctx *gin.Context) {

		isValidated := true

		if !isValidated {
			ctx.AbortWithStatusJSON(400, gin.H{
				"message": "bad request, some field not valid.",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"hello": "hello, world!",
		})

	})

}
