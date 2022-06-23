package router

import (
	"github.com/gin-gonic/gin"
	"mlinaa/pkg/logger"
	"net/http"
)

func NewRouter(handler *gin.Engine, logger logger.Interface) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/check", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	handler.POST("/set", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

}
