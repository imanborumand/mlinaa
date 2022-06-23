package app

import (
	"github.com/gin-gonic/gin"
	"mlinaa/config"
	router "mlinaa/internal/controller/http"
	"mlinaa/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New("start app.go")

	// HTTP Server
	handler := gin.New()
	router.NewRouter(handler, l)
	handler.Run(":" + cfg.Port)
}
