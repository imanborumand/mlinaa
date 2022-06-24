package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mlinaa/config"
	router "mlinaa/internal/controller/http"
	"mlinaa/internal/usecase"
	"mlinaa/internal/usecase/repo"
	"mlinaa/pkg/logger"
	"mlinaa/pkg/mongodb"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New("start app.go")

	// HTTP Server
	handler := gin.New()
	mongo, err := mongodb.New()
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - mongo.New: %w", err))
	}
	// Use case
	logUseCase := usecase.New(
		repo.New(mongo),
	)

	router.NewRouter(handler, logUseCase)
	handler.Run(":" + cfg.Port)
}
