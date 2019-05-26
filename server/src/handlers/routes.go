package handlers

import(
  "github.com/gin-gonic/gin"
  queries "queries"
)

func StartServer() {
  router := gin.Default()
  router.GET("/api/users", func (ctx *gin.Context) {
    ctx.String(200, queries.UserIndex())
  })

  router.Run(":8080")
}
