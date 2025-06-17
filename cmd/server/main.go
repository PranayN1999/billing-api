package main

import (
	"net/http"

	"github.com/PranayN1999/billing-api/internal/api"
	"github.com/PranayN1999/billing-api/internal/auth"
	"github.com/PranayN1999/billing-api/internal/config"
	"github.com/PranayN1999/billing-api/internal/db"
	"github.com/PranayN1999/billing-api/internal/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	dbx := db.Connect(cfg.DBURL)
	// AutoMigrate for this iteration
	dbx.AutoMigrate(&db.User{})

	authMW, err := auth.NewGinJWT(dbx, cfg.JWTSecret)
	if err != nil {
		panic(err)
	}

	r := gin.New()
	r.Use(middleware.ZapLogger(logger), gin.Recovery())

	// health
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// public signup + login
	r.POST("/auth/signup", api.SignUpHandler(dbx))
	r.POST("/auth/login", authMW.LoginHandler)

	// Protected group
	authGrp := r.Group("/")
	authGrp.Use(authMW.MiddlewareFunc())
	authGrp.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "pong"})
	})

	logger.Info("listening", zap.String("port", cfg.Port))
	r.Run(":" + cfg.Port)
}
