package controllers

import (
	multipart "mime/multipart"

	data "github.com/estebanborai/songs-share-server/server/src/data"
	helpers "github.com/estebanborai/songs-share-server/server/src/helpers"
	models "github.com/estebanborai/songs-share-server/server/src/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
)

// UpdateAvatar creates a base64 encoding image and saves it to the specified
// user as avatar
func UpdateAvatar(c *gin.Context, avatarImage *multipart.FileHeader, userID string) (response bool, err error) {
	file, err := avatarImage.Open()
	if err != nil {
		return false, err
	}

	avatarBase64 := helpers.FileToBase64(file)

	db := data.Connection(c)

	db.AutoMigrate(&models.Avatar{})

	var id = uuid.New().String()
	avatar := models.Avatar{
		ID:     id,
		Image:  avatarBase64,
		UserID: userID,
	}

	db.NewRecord(avatar)
	db.Create(&avatar)

	return true, nil
}
