package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mlinaa/internal/usecase"
	"net/http"
)

func NewRouter(handler *gin.Engine, log usecase.LogInterface) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/check", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	handler.POST("/set", func(c *gin.Context) {
		fmt.Println(log.Set(c))

		c.Status(http.StatusOK)
	})

}
