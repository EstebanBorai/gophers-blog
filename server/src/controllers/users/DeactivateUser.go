package controllers

import (
	data "github.com/estebanborai/songs-share-server/server/src/data"
	"github.com/estebanborai/songs-share-server/server/src/helpers/gimlet"
	models "github.com/estebanborai/songs-share-server/server/src/models"
	"github.com/gin-gonic/gin"
)

// DeactivationResults represents the response of the API
// after deactivating an user
type DeactivationResults struct {
	ID string
}

// DeactivateUser deactivates the user with the given ID
func DeactivateUser(c *gin.Context) {
	userID := c.Param("id")

	db := data.Connection(c)

	if result := db.Table("users").Where(&models.User{ID: userID}).Update("Active", false); result.Error == nil {
		c.JSON(200, DeactivationResults{
			ID: userID,
		})
	} else {
		gimlet.BadRequest(c, result.Error.Error())
	}
}
