package routes

import(
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  controllers "controllers"
)

func StartServer() {
  router := gin.Default()

  router.Use(cors.Default())

  router.POST("/v1/users", func (ctx *gin.Context) {
    ctx.String(200, controllers.CreateUser())
  })

  router.Run(":8080")
}
