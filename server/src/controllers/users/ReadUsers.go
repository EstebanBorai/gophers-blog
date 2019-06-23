package controllers

import (
	data "github.com/estebanborai/songs-share-server/server/src/data"
	models "github.com/estebanborai/songs-share-server/server/src/models"
	"github.com/gin-gonic/gin"
)

// ReadUsers indexes all users in the database
func ReadUsers(c *gin.Context) {
	var users []models.User

	db := data.Connection(c)

	db.Preload("Avatar").Find(&users)

	c.JSON(200, users)
}
