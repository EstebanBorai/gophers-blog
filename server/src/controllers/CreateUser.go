package controllers

import (
  "time"
  uuid "github.com/google/uuid"
  models "models"
  "encoding/json"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateUser(encodedUser string) (response string, err error) {
  db, err := gorm.Open("mysql", "root:root@tcp(songs-share-db)/songs-share?charset=utf8mb4&parseTime=True&loc=Local")

  if err != nil {
    return "", err
  } else {
    db.AutoMigrate(&models.User{})
  }

  defer db.Close()

  var id string = uuid.New().String()
  var decodedUser models.User

  json.Unmarshal([]byte(encodedUser), &decodedUser)

  user := models.User {
    Id: id,
    UserName: decodedUser.UserName,
    FirstName: decodedUser.FirstName,
    LastName: decodedUser.LastName,
    Birthday: decodedUser.Birthday,
    DateJoined: time.Now(),
  }

  db.NewRecord(user)
  db.Create(&user)

  jsonEncoded, err := json.Marshal(user)
  return string(jsonEncoded), nil
}
