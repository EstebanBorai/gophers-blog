package controllers

import (
	"encoding/json"
	"fmt"

	data "github.com/estebanborai/go-server-sample/server/src/data"
	helpers "github.com/estebanborai/go-server-sample/server/src/helpers"
	"github.com/estebanborai/go-server-sample/server/src/helpers/gimlet"
	models "github.com/estebanborai/go-server-sample/server/src/models"
	security "github.com/estebanborai/go-server-sample/server/src/security"
	"github.com/gin-gonic/gin"
)

// UpdatePasswordPayload represents expected HTTP Body to meet
// password update process
type UpdatePasswordPayload struct {
	ID       string
	Password string
}

// UpdatePassword replaces current user password
func UpdatePassword(c *gin.Context) {
	var user models.User
	var userSecret models.UserSecret
	var decodedPayload UpdatePasswordPayload
	var encodedPayload = helpers.ContextRequestBody(c)

	if err := json.Unmarshal([]byte(encodedPayload), &decodedPayload); err != nil {
		var msg = "Invalid JSON " + err.Error()
		gimlet.BadRequest(c, msg)
	}

	db := data.Connection(c)

	if userResult := db.Where(&models.User{ID: decodedPayload.ID}).First(&user); userResult.Error == nil {
		if userSecretResult := db.Where(&models.UserSecret{UserID: decodedPayload.ID}).First(&userSecret); userSecretResult.Error == nil {
			if decodedPayload.Password == "" || decodedPayload.ID == "" {
				gimlet.BadRequest(c, "Missing fields")
			} else {
				db.Model(&userSecret).Update("hash", security.EncryptPassword(decodedPayload.Password))
				c.JSON(200, nil)
			}
		} else {
			gimlet.InternalServerError(c, userSecretResult.Error.Error())
		}
	} else {
		gimlet.NotFound(c, fmt.Sprintf("%s doesn't exists", decodedPayload.ID))
	}
}
