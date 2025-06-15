package middleware

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ZapLogger returns a Gin middleware that logs each request.
func ZapLogger(l *zap.Logger) gin.HandlerFunc {
	return ginzap.Ginzap(l, time.RFC3339, true)
}
