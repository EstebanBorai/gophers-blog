package controllers

import (
	data "github.com/estebanborai/songs-share-server/server/src/data"
	models "github.com/estebanborai/songs-share-server/server/src/models"
	security "github.com/estebanborai/songs-share-server/server/src/security"
	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
)

func CreatePassword(c *gin.Context, password string, userId string) (response bool, err error) {
	db := data.Connection(c)

	if err != nil {
		return false, err
	} else {
		db.AutoMigrate(&models.UserSecret{})
	}

	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	var id string = uuid.New().String()
	userSecret := models.UserSecret{
		Id:     id,
		UserId: userId,
		Hash:   security.EncryptPassword(password),
	}

	db.NewRecord(userSecret)
	db.Create(&userSecret)

	return true, nil
}
