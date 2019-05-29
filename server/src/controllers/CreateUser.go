package controllers

import (
  "time"
  uuid "github.com/google/uuid"
  models "models"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateUser() string {
  db, err := gorm.Open("mysql", "root:root@tcp(songs-share-db)/songs-share?charset=utf8mb4&parseTime=True&loc=Local")

  if err != nil {
    panic(err.Error())
  } else {
    db.Debug().DropTableIfExists(&models.User{}) 
    db.Debug().AutoMigrate(&models.User{}) 
  }

  defer db.Close()

  var id string = uuid.New().String()

  user := models.User {
    Id: id,
    UserName: "estebanborai",
    FirstName: "Esteban",
    LastName: "Borai",
    Birthday: time.Date(
      1990, 11, 17, 20, 34, 58, 651387237, time.UTC),
    DateJoined: time.Now(),
  }

  db.NewRecord(user)
  db.Create(&user)

  return "ok"
}
