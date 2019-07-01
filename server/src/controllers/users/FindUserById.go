package controllers

import (
	"fmt"

	data "github.com/estebanborai/go-server-sample/server/src/data"
	"github.com/estebanborai/go-server-sample/server/src/helpers/gimlet"
	models "github.com/estebanborai/go-server-sample/server/src/models"
	"github.com/gin-gonic/gin"
)

// FindUserByID returs the user with the specified ID
func FindUserByID(c *gin.Context, ID string) {
	var user models.User

	db := data.Connection(c)

	if userResult := db.Where(&models.User{ID: ID}).First(&user); userResult.Error == nil {
		c.JSON(200, user)
	} else {
		gimlet.NotFound(c, fmt.Sprintf("%s doesn't exists", ID))
	}
}
