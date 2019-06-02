package controllers

// FIXME: Enhance Error Handlers to stop function instead of using return 

import (
  "time"
  "strings"
  "encoding/json"
  "github.com/jinzhu/gorm"
  "github.com/gin-gonic/gin"
  uuid "github.com/google/uuid"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  models "github.com/estebanborai/songs-share-server/models"
  helpers "github.com/estebanborai/songs-share-server/helpers"
  eh "github.com/estebanborai/songs-share-server/lib/error_handlers"
)

type Payload struct {
  models.User
  Password string
}

func CreateUser(c *gin.Context) {
  var id string = uuid.New().String()
  var decodedPayload Payload
  var encodedUser string = helpers.ContextRequestBody(c)

  if err := json.Unmarshal([]byte(encodedUser), &decodedPayload); err != nil {
    var msg string = "Invalid JSON " + err.Error()
    eh.ResponseWithError(c, 400, msg)
  }

  user := models.User {
    Id: id,
    UserName: decodedPayload.UserName,
    FirstName: decodedPayload.FirstName,
    LastName: decodedPayload.LastName,
    Birthday: decodedPayload.Birthday,
    DateJoined: time.Now(),
  }

  db, err := gorm.Open("mysql", "root:root@tcp(songs-share-db)/songs-share?charset=utf8mb4&parseTime=True&loc=Local")

  if err != nil {
    eh.ResponseWithError(c, 500, "Unable to connect to the database")
    return
  } else {
    // TODO: Automigrate on Init
    db.AutoMigrate(&models.User{})
  }

  defer db.Close()
  
  db.NewRecord(user)

  if dbc := db.Create(&user); dbc.Error != nil {
    errorString := dbc.Error.Error()
    var isUserNameTaken bool = strings.Contains(errorString, "Error 1062")

    if isUserNameTaken == true {
      eh.ResponseWithError(c, 400, "Username " + user.UserName + " is already taken.")
      return
    } else {
      eh.ResponseWithError(c, 400, errorString)
      return
    }
  }
  
  _, passwordError := CreatePassword(decodedPayload.Password, user.Id)

  if passwordError != nil {
    eh.ResponseWithError(c, 400, "Invalid Password")
    return
  } else {
    c.JSON(200, user)
  }
}
