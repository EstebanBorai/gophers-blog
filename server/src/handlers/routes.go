package handlers

import(
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  queries "queries"
)

func StartServer() {
  router := gin.Default()

  router.Use(cors.Default())

  router.GET("/api/users", func (ctx *gin.Context) {
    ctx.String(200, queries.UserIndex())
  })

  router.POST("/api/users", func (ctx *gin.Context) {
    buf := make([]byte, 1024)
    num, _ := ctx.Request.Body.Read(buf)
    reqBody := string(buf[0:num])

    result := queries.InsertUser(reqBody)

    ctx.JSON(200, result)
  })

  router.Run(":8080")
}
