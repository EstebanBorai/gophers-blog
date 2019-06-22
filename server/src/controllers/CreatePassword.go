package controllers

import (
	models "github.com/estebanborai/songs-share-server/server/src/models"
	security "github.com/estebanborai/songs-share-server/server/src/security"
	uuid "github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func CreatePassword(password string, userId string) (response bool, err error) {
	db, err := gorm.Open("mysql", "root:root@tcp(songs-share-db)/songs-share?charset=utf8mb4&parseTime=True&loc=Local")

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
