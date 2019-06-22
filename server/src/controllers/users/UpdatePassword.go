package controllers

import (
	"encoding/json"
	"fmt"

	data "github.com/estebanborai/songs-share-server/server/src/data"
	helpers "github.com/estebanborai/songs-share-server/server/src/helpers"
	eh "github.com/estebanborai/songs-share-server/server/src/lib/error_handlers"
	models "github.com/estebanborai/songs-share-server/server/src/models"
	security "github.com/estebanborai/songs-share-server/server/src/security"
	"github.com/gin-gonic/gin"
)

type UpdatePasswordPayload struct {
	Id       string
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

	if userResult := db.Where(&models.User{Id: decodedPayload.Id}).First(&user); userResult.Error == nil {
		if userSecretResult := db.Where(&models.UserSecret{UserId: decodedPayload.Id}).First(&userSecret); userSecretResult.Error == nil {
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
