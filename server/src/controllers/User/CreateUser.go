package controllers

import (
  database "database"
  uuid "github.com/google/uuid"
  database "database"
  "models"
)

func CreateUser() string {
  db := database.Connect()

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
