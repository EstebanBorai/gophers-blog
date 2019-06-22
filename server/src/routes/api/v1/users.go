package routes

import (
	"github.com/appleboy/gin-jwt"
	controllers "github.com/estebanborai/songs-share-server/server/src/controllers/users"
	"github.com/gin-gonic/gin"
)

func UsersRoute(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	users := r.Group("api/v1/users")

	users.Use(authMiddleware.MiddlewareFunc())
	{
		users.GET("/", controllers.ReadUsers)

		users.GET("/id/:id", func(c *gin.Context) {
			// TODO: Refactor use single parameter function
			id := c.Param("id")

			controllers.FindUserById(c, id)
		})

		users.GET("/username/:username", func(c *gin.Context) {
			// TODO: Refactor use single parameter function
			userName := c.Param("username")

			controllers.FindUserByUserName(c, userName)
		})

		users.PATCH("/id/:id", controllers.UpdateUser)

		users.PATCH("/update/password", controllers.UpdatePassword)
	}
}
