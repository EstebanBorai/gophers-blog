package middleware

import (
	"fmt"
	"time"

	"github.com/appleboy/gin-jwt"
	models "github.com/estebanborai/songs-share-server/server/src/models"
	security "github.com/estebanborai/songs-share-server/server/src/security"
	"github.com/gin-gonic/gin"
)

var identityKey string = "id"
var email string = "email"
var userName string = "userName"

func AuthMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "songs-share-auth",
		Key:         []byte("songs-share-key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Id,
					email:       v.Email,
					userName:    v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				Id:       claims["id"].(string),
				Email:    claims["email"].(string),
				UserName: claims["userName"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var credentials models.UserCredentials
			if err := c.ShouldBind(&credentials); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userName := credentials.UserName
			password := credentials.Password
			user, err := security.ValidateUser(c, userName, password)

			if err == nil {
				return &models.User{
					Id:       user.Id,
					UserName: user.UserName,
					Email:    user.Email,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// Theres no roles in this application
			// if v, ok := data.(*models.User); ok && v.UserName == "admin" {
			//   return true
			// }

			// return false
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			// TODO: Use eh instead
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	return authMiddleware
}
