package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/PranayN1999/billing-api/internal/config"
	"github.com/PranayN1999/billing-api/internal/middleware"
)

func main() {
	cfg := config.Load()

	// logger for middleware
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	r := gin.New()
	r.Use(middleware.ZapLogger(logger), gin.Recovery())

	// health probe
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	logger.Info("listening", zap.String("port", cfg.Port))
	if err := r.Run(":" + cfg.Port); err != nil {
		logger.Fatal("server error", zap.Error(err))
	}
}
