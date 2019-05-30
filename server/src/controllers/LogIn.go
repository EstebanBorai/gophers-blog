package controllers

import (
  models "models"
  "encoding/json"
  security "security"
  "github.com/jinzhu/gorm"
)

func LogIn(credentials string) security.JWTCookie {
  var user models.User
  var creds models.UserCredentials
  var userSecret models.UserSecret

  if err := json.Unmarshal([]byte(credentials), &creds); err != nil {
    panic(err.Error())
  }

  db, err := gorm.Open("mysql", "root:root@tcp(songs-share-db)/songs-share?charset=utf8mb4&parseTime=True&loc=Local")

  if err != nil {
    panic(err.Error())
  }

  defer db.Close()

  db.Where(models.User{UserName: creds.UserName}).First(&user)
  db.Where(models.UserSecret{UserId: user.Id}).First(&userSecret)

  if security.ValidatePassword(userSecret.Hash, creds.Password) {
    return security.CreateToken(creds.UserName)
  } else {
    panic("Password is not correct")
  }
}
