package api

import (
	"net/http"

	"github.com/PranayN1999/billing-api/internal/auth"
	"github.com/PranayN1999/billing-api/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SignUpHandler(dbx *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct{ Email, Password string }
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad JSON"})
			return
		}
		hash, _ := auth.HashPassword(req.Password)
		u := db.User{ID: uuid.NewString(), Email: req.Email, PasswordHash: hash}

		if err := dbx.Create(&u).Error; err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "email exists"})
			return
		}
		c.Status(http.StatusCreated)
	}
}
