package routes

import(
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  controllers "controllers"
)

func StartServer() {
  router := gin.Default()

  router.Use(cors.Default())

  router.POST("/api/v1/users", func (c *gin.Context) {
    buf := make([]byte, 1024)
    num, _ := c.Request.Body.Read(buf)
    reqBody := string(buf[0:num])

    response, err := controllers.CreateUser(reqBody)

    if err != nil {
      c.JSON(500, err)
    } else {
      c.JSON(200, response)
    }
  })

  router.Run(":8080")
}
