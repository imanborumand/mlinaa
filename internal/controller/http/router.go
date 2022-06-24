package router

import (
	"github.com/gin-gonic/gin"
	"mlinaa/internal/entity"
	"mlinaa/internal/usecase"
	"net/http"
)

func NewRouter(handler *gin.Engine, log usecase.LogInterface) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.POST("/set", func(c *gin.Context) {
		var request entity.StoreRequestBody
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{})
		}

		c.JSON(http.StatusOK, gin.H{
			"id": log.Set(request),
		})
	})

	handler.GET("/get/:id", func(c *gin.Context) {
		id := c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			"data": log.GetById(id),
		})
	})

	handler.DELETE("/delete/:id", func(c *gin.Context) {
		id := c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			"status": log.Delete(id),
		})
	})

}
