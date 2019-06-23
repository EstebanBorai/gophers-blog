package routes

import (
	controllers "github.com/estebanborai/songs-share-server/server/src/controllers"
	middleware "github.com/estebanborai/songs-share-server/server/src/middleware"
	v1 "github.com/estebanborai/songs-share-server/server/src/routes/api/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// StartServer assigns routes and initializes server
func StartServer() {
	r := gin.Default()
	store := cookie.NewStore([]byte("songs-share-store"))
	authMiddleware := middleware.AuthMiddleware()

	r.Use(cors.Default())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(sessions.Sessions("sessions", store))

	v1.UsersRoute(r, authMiddleware)

	r.POST("/login", authMiddleware.LoginHandler)

	r.POST("/signup", func(c *gin.Context) {
		controllers.SignUp(c)
	})

	r.Run(":8080")
}
