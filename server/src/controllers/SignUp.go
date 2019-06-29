package controllers

import (
	"fmt"
	"strings"
	"time"

	users "github.com/estebanborai/songs-share-server/server/src/controllers/users"
	data "github.com/estebanborai/songs-share-server/server/src/data"
	"github.com/estebanborai/songs-share-server/server/src/helpers/gimlet"
	models "github.com/estebanborai/songs-share-server/server/src/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
)

// Payload holds the spected properties to be sent from the Front-End
// at the time the user signs up
type Payload struct {
	models.User
	Password string
	Avatar   string
}

// SignUp creates a new user, hashes/salt the user password and stores user avatar
func SignUp(c *gin.Context) {
	var id = uuid.New().String()
	decodedPayload, _ := c.MultipartForm()

	birthDay, dateError := time.Parse(time.RFC3339, decodedPayload.Value["birthday"][0])

	if dateError != nil {
		gimlet.BadRequest(c, "Invalid Date")
		return
	}

	user := models.User{
		ID:         id,
		UserName:   decodedPayload.Value["userName"][0],
		FirstName:  decodedPayload.Value["firstName"][0],
		LastName:   decodedPayload.Value["lastName"][0],
		Email:      decodedPayload.Value["email"][0],
		Birthday:   birthDay,
		DateJoined: time.Now(),
	}

	db := data.Connection(c)

	// TODO: Automigrate on Init
	db.AutoMigrate(&models.User{})

	defer db.Close()

	db.NewRecord(user)

	if dbc := db.Create(&user); dbc.Error != nil {
		errorString := dbc.Error.Error()
		var isUserNameTaken = strings.Contains(errorString, "Error 1062")

		if isUserNameTaken == true {
			gimlet.BadRequest(c, fmt.Sprintf("Username %s is already taken", user.UserName))
			return
		}

		gimlet.InternalServerError(c, errorString)
		return
	}

	_, passwordError := CreatePassword(c, decodedPayload.Value["password"][0], user.ID)
	if passwordError != nil {
		gimlet.BadRequest(c, "Invalid Password")
		return
	}

	avatarUpdate := users.UpdateAvatar(c, id)

	if avatarUpdate == false {
		gimlet.BadRequest(c, "Unable to set avatar")
	}

	c.JSON(200, user)
}
