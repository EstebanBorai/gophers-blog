package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/appleboy/gin-jwt"
  controllers "github.com/estebanborai/songs-share-server/controllers"
)

func UsersRoute(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
  users := r.Group("api/v1/users")

  users.Use(authMiddleware.MiddlewareFunc())
  {
    users.GET("/", func (c *gin.Context) {
      controllers.ReadUsers(c)
    })
  }
}
