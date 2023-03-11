package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterApiV1(router *gin.Engine) {
	apiV1 := router.Group("/v1")

	apiV1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 1. Create a credential
	// 2.

}
