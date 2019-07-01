package controllers

import (
	data "github.com/estebanborai/go-server-sample/server/src/data"
	models "github.com/estebanborai/go-server-sample/server/src/models"
	security "github.com/estebanborai/go-server-sample/server/src/security"
	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
)

// CreatePassword stores user hashed password
func CreatePassword(c *gin.Context, password string, userID string) (response bool, err error) {
	db := data.Connection(c)

	if err != nil {
		return false, err
	}

	db.AutoMigrate(&models.UserSecret{})

	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	var id = uuid.New().String()
	userSecret := models.UserSecret{
		ID:     id,
		UserID: userID,
		Hash:   security.EncryptPassword(password),
	}

	db.NewRecord(userSecret)
	db.Create(&userSecret)

	return true, nil
}
