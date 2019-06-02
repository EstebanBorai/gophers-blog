package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
  controllers "github.com/estebanborai/songs-share-server/controllers"
)

func StartServer() {
  router := gin.Default()

  router.Use(cors.Default())

  router.POST("/api/v1/users", func (c *gin.Context) {
    controllers.CreateUser(c)
  })

  router.POST("/login", func (c *gin.Context) {
    controllers.LogIn(c)
  })

  router.Run(":8080")
}
