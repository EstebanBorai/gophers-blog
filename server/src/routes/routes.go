package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
  "github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
  controllers "github.com/estebanborai/songs-share-server/controllers"
)

func StartServer() {
  r := gin.Default()
  store := cookie.NewStore([]byte("songs-share-store"))
  
  r.Use(cors.Default())
  r.Use(sessions.Sessions("sessions", store))

  r.GET("/api/v1/users", func (c *gin.Context) {
    controllers.ReadUsers(c)
  })

  r.POST("/api/v1/users", func (c *gin.Context) {
    controllers.CreateUser(c)
  })

  r.POST("/login", func (c *gin.Context) {
    session := sessions.Default(c)
    var count int

    v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
    }

		session.Set("count", count)
    session.Save()

    controllers.LogIn(c)
  })

  r.Run(":8080")
}
