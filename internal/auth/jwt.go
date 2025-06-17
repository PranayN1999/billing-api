package auth

import (
	"time"

	"github.com/PranayN1999/billing-api/internal/db"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const identityKey = "userID"

func NewGinJWT(dbx *gorm.DB, secret string) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "billing",
		Key:         []byte(secret),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: identityKey,

		Authenticator: func(c *gin.Context) (interface{}, error) {
			var creds struct{ Email, Password string }
			if err := c.ShouldBindJSON(&creds); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}
			var u db.User
			if err := dbx.First(&u, "email = ?", creds.Email).Error; err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			if err := CheckPassword(u.PasswordHash, creds.Password); err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			return &u, nil
		},

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if u, ok := data.(*db.User); ok {
				return jwt.MapClaims{identityKey: u.ID}
			}
			return jwt.MapClaims{}
		},

		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &db.User{ID: claims[identityKey].(string)}
		},

		Authorizator:  func(data interface{}, c *gin.Context) bool { return true },
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
	})
}
