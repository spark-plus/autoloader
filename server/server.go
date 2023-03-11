package server

import (
	"net/http"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/sparkster/autoloader/server/routes"
)

func GinServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	routes.RegisterApiV1(r)
	endless.ListenAndServe(":8080", r) // listen and serve on http://localhost:8080 (for windows "localhost:8080")
}
