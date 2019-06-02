package controllers

import (
  models "github.com/estebanborai/songs-share-server/models"
  "encoding/json"
  helpers "github.com/estebanborai/songs-share-server/helpers"
  security "github.com/estebanborai/songs-share-server/security"
  eh "github.com/estebanborai/songs-share-server/lib/error_handlers"
  "github.com/jinzhu/gorm"
  "github.com/gin-gonic/gin"
)

func LogIn(c *gin.Context) {
  var user models.User
  var creds models.UserCredentials
  var userSecret models.UserSecret
  var credentials string = helpers.ContextRequestBody(c)

  if err := json.Unmarshal([]byte(credentials), &creds); err != nil {
    var msg string = "Invalid JSON " + err.Error()
    eh.ResponseWithError(c, 400, msg)
  }

  db, err := gorm.Open("mysql", "root:root@tcp(songs-share-db)/songs-share?charset=utf8mb4&parseTime=True&loc=Local")

  if err != nil {
    eh.ResponseWithError(c, 500, "Unable to connect to the database")
  }

  defer db.Close()

  db.Where(models.User{UserName: creds.UserName}).First(&user)
  db.Where(models.UserSecret{UserId: user.Id}).First(&userSecret)

  if security.ValidatePassword(userSecret.Hash, creds.Password) {
    cookie := security.CreateToken(creds.UserName)
    c.SetCookie(
      cookie.Name,
      cookie.Value,
      3600,
      "/",
      "localhost",
      false,
      true,
    )
  } else {
    eh.ResponseWithError(c, 403, "Invalid Username/Password")
  }
}
