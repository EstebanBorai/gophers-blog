package routes

import (
	jwt "github.com/appleboy/gin-jwt"
	controllers "github.com/estebanborai/songs-share-server/server/src/controllers/users"
	"github.com/gin-gonic/gin"
)

// UsersRoute assigns api/v1/users controllers to routes
func UsersRoute(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	users := r.Group("api/v1/users")

	users.Use(authMiddleware.MiddlewareFunc())
	{
		users.GET("/", controllers.ReadUsers)

		users.GET("/id/:id", func(c *gin.Context) {
			// TODO: Refactor use single parameter function
			id := c.Param("id")

			controllers.FindUserByID(c, id)
		})

		users.GET("/username/:username", func(c *gin.Context) {
			// TODO: Refactor use single parameter function
			userName := c.Param("username")

			controllers.FindUserByUserName(c, userName)
		})

		users.PATCH("/id/:id", controllers.UpdateUser)

		users.PATCH("/update/password", controllers.UpdatePassword)

		users.DELETE("/id/:id", controllers.DeactivateUser)
	}
}
