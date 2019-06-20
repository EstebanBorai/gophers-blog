package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/appleboy/gin-jwt"
  controllers "github.com/estebanborai/songs-share-server/controllers/users"
  
)

func UsersRoute(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
  users := r.Group("api/v1/users")

  users.Use(authMiddleware.MiddlewareFunc())
  {
    users.GET("/", func (c *gin.Context) {
      controllers.ReadUsers(c)
    })

    users.GET("/id/:id", func (c *gin.Context) {
      id := c.Param("id")

      controllers.FindUserById(c, id)
    })

    users.GET("/username/:username", func (c *gin.Context) {
      userName := c.Param("username")

      controllers.FindUserByUserName(c, userName)
    })

    users.PATCH("/update/password", func (c *gin.Context) {
      controllers.UpdatePassword(c)
    })
  }
}
