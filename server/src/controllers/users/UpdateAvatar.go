package controllers

import (
	data "github.com/estebanborai/songs-share-server/server/src/data"
	"github.com/estebanborai/songs-share-server/server/src/helpers"
	"github.com/estebanborai/songs-share-server/server/src/helpers/gimlet"
	"github.com/estebanborai/songs-share-server/server/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UpdateAvatar stores an user's avatar
func UpdateAvatar(c *gin.Context, userID string) bool {
	var avatarID = uuid.New().String()
	decodedPayload, _ := c.MultipartForm()

	if decodedPayload == nil {
		return true
	}

	if len(decodedPayload.File) == 0 {
		return true
	}

	if _, ok := decodedPayload.File["avatar"]; !ok {
		return true
	}

	avatarFile := decodedPayload.File["avatar"][0]

	if avatarFile != nil {
		file, err := avatarFile.Open()

		if err != nil {
			gimlet.InternalServerError(c, err.Error())
		} else {
			avatarBase64 := helpers.FileToBase64(file)

			avatar := models.Avatar{
				ID:     avatarID,
				Image:  avatarBase64,
				UserID: userID,
			}

			db := data.Connection(c)

			db.AutoMigrate(&models.Avatar{})

			db.NewRecord(avatar)
			db.Create(&avatar)

			return true
		}

		return false
	}

	return false
}
