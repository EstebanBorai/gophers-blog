package controllers

import (
	"encoding/json"

	data "github.com/estebanborai/songs-share-server/server/src/data"
	helpers "github.com/estebanborai/songs-share-server/server/src/helpers"
	eh "github.com/estebanborai/songs-share-server/server/src/lib/error_handlers"
	models "github.com/estebanborai/songs-share-server/server/src/models"
	"github.com/gin-gonic/gin"
)

type UpdateUserPayload struct {
	UserName  string
	FirstName string
	LastName  string
}

func UpdateUser(c *gin.Context) {
	var user models.User
	var decodedPayload UpdateUserPayload
	var encodedPayload string = helpers.ContextRequestBody(c)
	userId := c.Param("id")

	if err := json.Unmarshal([]byte(encodedPayload), &decodedPayload); err != nil {
		var msg string = "Invalid JSON " + err.Error()
		eh.BadRequest(c, msg)
	}

	db := data.Connection(c)

	if result := db.Table("users").Where(&models.User{Id: userId}).Updates(models.User{
		UserName:  decodedPayload.UserName,
		FirstName: decodedPayload.FirstName,
		LastName:  decodedPayload.LastName,
	}); result.Error == nil {
		db.Where(&models.User{Id: userId}).First(&user)
		c.JSON(200, user)
	} else {
		eh.BadRequest(c, result.Error.Error())
	}
}
