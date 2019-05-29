package controllers

import (
  "time"
  uuid "github.com/google/uuid"
  models "models"
  "encoding/json"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Payload struct {
  models.User
  Password string
}

func CreateUser(encodedUser string) (response string, err error) {
  db, err := gorm.Open("mysql", "root:root@tcp(songs-share-db)/songs-share?charset=utf8mb4&parseTime=True&loc=Local")

  if err != nil {
    return "", err
  } else {
    db.AutoMigrate(&models.User{})
  }

  defer db.Close()

  var id string = uuid.New().String()
  var decodedPayload Payload

  json.Unmarshal([]byte(encodedUser), &decodedPayload)

  user := models.User {
    Id: id,
    UserName: decodedPayload.UserName,
    FirstName: decodedPayload.FirstName,
    LastName: decodedPayload.LastName,
    Birthday: decodedPayload.Birthday,
    DateJoined: time.Now(),
  }

  db.NewRecord(user)
  db.Create(&user)

  _, passwordError := CreatePassword(decodedPayload.Password, user.Id)

  if passwordError != nil {
    return "", passwordError
  }

  jsonEncoded, err := json.Marshal(user)

  if err != nil {
    return "", err
  }

  return string(jsonEncoded), nil
}
