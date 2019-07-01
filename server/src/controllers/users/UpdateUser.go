package controllers

import (
	"encoding/json"

	data "github.com/estebanborai/go-server-sample/server/src/data"
	helpers "github.com/estebanborai/go-server-sample/server/src/helpers"
	"github.com/estebanborai/go-server-sample/server/src/helpers/gimlet"
	models "github.com/estebanborai/go-server-sample/server/src/models"
	"github.com/gin-gonic/gin"
)

// A UpdateUserPayload represents the expected HTTP Body in the request
// to update user data
type UpdateUserPayload struct {
	UserName  string
	FirstName string
	LastName  string
}

// UpdateUser updates user properties based on non-empty and different values
func UpdateUser(c *gin.Context) {
	var user models.User
	var decodedPayload UpdateUserPayload
	var encodedPayload = helpers.ContextRequestBody(c)
	userID := c.Param("id")

	if err := json.Unmarshal([]byte(encodedPayload), &decodedPayload); err != nil {
		var msg = "Invalid JSON " + err.Error()
		gimlet.BadRequest(c, msg)
	}

	db := data.Connection(c)

	if result := db.Table("users").Where(&models.User{ID: userID}).Updates(models.User{
		UserName:  decodedPayload.UserName,
		FirstName: decodedPayload.FirstName,
		LastName:  decodedPayload.LastName,
	}); result.Error == nil {
		db.Where(&models.User{ID: userID}).First(&user)
		c.JSON(200, user)
	} else {
		gimlet.BadRequest(c, result.Error.Error())
	}
}
