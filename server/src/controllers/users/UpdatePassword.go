package controllers

import (
	"fmt"
	"encoding/json"
  "github.com/gin-gonic/gin"
  data "github.com/estebanborai/songs-share-server/data"
  models "github.com/estebanborai/songs-share-server/models"
	helpers "github.com/estebanborai/songs-share-server/helpers"
	security "github.com/estebanborai/songs-share-server/security"
	eh "github.com/estebanborai/songs-share-server/lib/error_handlers"
)

type UpdatePasswordPayload struct {
	Id string
	Password string
}

func UpdatePassword(c *gin.Context) {
	var user models.User
	var userSecret models.UserSecret
	var decodedPayload UpdatePasswordPayload
	var encodedPayload string = helpers.ContextRequestBody(c)

	if err := json.Unmarshal([]byte(encodedPayload), &decodedPayload); err != nil {
		var msg string = "Invalid JSON " + err.Error()
		eh.BadRequest(c, msg)
	}

	db := data.Connection(c)

	if userResult := db.Where(&models.User{ Id: decodedPayload.Id }).First(&user); userResult.Error == nil {
		if userSecretResult := db.Where(&models.UserSecret{ UserId: decodedPayload.Id }).First(&userSecret); userSecretResult.Error == nil {
			if decodedPayload.Password == "" || decodedPayload.Id == "" {
				eh.BadRequest(c, "Missing fields")
			} else {
				db.Model(&userSecret).Update("hash", security.EncryptPassword(decodedPayload.Password))
				c.JSON(200, nil)
			}
		} else {
			eh.InternalServerError(c, userSecretResult.Error.Error())
		}
	} else {
		eh.NotFound(c, fmt.Sprintf("%s doesn't exists", decodedPayload.Id))
	}
}
