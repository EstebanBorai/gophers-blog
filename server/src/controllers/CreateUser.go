package controllers

import (
  uuid "github.com/google/uuid"
  "models"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateUser() string {
  db, err := gorm.Open("mysql", "admin:admin@/songs-share?charset=utf8mb4&parseTime=True&loc=Local")

  if err != nil {
    return nil, err
  }

  defer db.Close()

  var id string = uuid.New().String()

  user := User {
    Id: id
    UserName: "estebanborai"
    FirstName: "Esteban"
    LastName: "Borai"
    Birthday: time.Date(
      1990, 11, 17, 20, 34, 58, 651387237, time.UTC)
    DateJoined: time.Now()
  }

  db.NewRecord(user)
  db.Create(&user)

  return "ok"
}
