package routes

import (
  controllers "controllers"
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
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

  router.POST("/login", func (c *gin.Context) {
    // buf := make([]byte, 1024)
    // num, _ := c.Request.Body.Read(buf)
    // reqBody := string(buf[0:num])

    // cookie := controllers.LogIn(reqBody)

    // c.SetCookie(
    //   cookie.Name,
    //   cookie.Value,
    //   3600,
    //   "/",
    //   "localhost",
    //   false,
    //   true,
    // )
    // controllers.LogIn(c)
    controllers.LogIn(c)
  })

  router.Run(":8080")
}
