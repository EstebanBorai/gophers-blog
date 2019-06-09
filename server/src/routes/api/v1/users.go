package routes

import (
  "github.com/gin-gonic/gin"
  controllers "github.com/estebanborai/songs-share-server/controllers"
)

func UsersRoute(r *gin.Engine) {
  users := r.Group("api/v1/users")

  users.GET("/", func (c *gin.Context) {
    controllers.ReadUsers(c)
  })

  users.POST("/", func (c *gin.Context) {
    controllers.CreateUser(c)
  })
}
