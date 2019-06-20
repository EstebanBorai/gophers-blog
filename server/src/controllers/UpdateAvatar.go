package controllers

import (
	multipart "mime/multipart"
	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	data "github.com/estebanborai/songs-share-server/data"
	models "github.com/estebanborai/songs-share-server/models"
	helpers "github.com/estebanborai/songs-share-server/helpers"
)

func UpdateAvatar(c *gin.Context, avatarImage *multipart.FileHeader, userId string) (response bool, err error) {
	file, err := avatarImage.Open(); if err != nil {
		return false, err
	}

	avatarBase64 := helpers.FileToBase64(file)

	db := data.Connection(c)

	db.AutoMigrate(&models.Avatar{})

	var id string = uuid.New().String()
	avatar := models.Avatar {
		Id: id,
		Image: avatarBase64,
		UserId: userId,
	}

	db.NewRecord(avatar)
	db.Create(&avatar)

	return true, nil
}
