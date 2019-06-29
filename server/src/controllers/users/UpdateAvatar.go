package controllers

import (
	"mime/multipart"

	data "github.com/estebanborai/songs-share-server/server/src/data"
	"github.com/estebanborai/songs-share-server/server/src/helpers"
	"github.com/estebanborai/songs-share-server/server/src/helpers/gimlet"
	"github.com/estebanborai/songs-share-server/server/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func isAvatarAvailable(payload *multipart.Form) bool {
	if payload.File["Avatar"] == nil {
		return false
	}

	if payload.File["Avatar"][0] == nil {
		return false
	}

	return true
}

// UpdateAvatar stores an user's avatar
func UpdateAvatar(c *gin.Context, userID string) bool {
	var avatarID = uuid.New().String()
	decodedPayload, _ := c.MultipartForm()

	if !isAvatarAvailable(decodedPayload) {
		return true
	}

	avatarFile := decodedPayload.File["Avatar"][0]

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
