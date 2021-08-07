package httpapp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func routes() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", ping)

	return router
}
