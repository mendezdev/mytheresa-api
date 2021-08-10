package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func routes(hdls handlers) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", ping)

	router.GET("/products", hdls.productHdl.GetAll)

	return router
}
