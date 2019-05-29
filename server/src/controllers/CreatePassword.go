package controllers

import (
  uuid "github.com/google/uuid"
  models "models"
  "golang.org/x/crypto/bcrypt"
  "github.com/jinzhu/gorm"
)

func CreatePassword(password string, userId string) (response bool, err error) {
  pwd := []byte(password)

  db, err := gorm.Open("mysql", "root:root@tcp(songs-share-db)/songs-share?charset=utf8mb4&parseTime=True&loc=Local")

  if err != nil {
    return false, err
  } else {
    db.AutoMigrate(&models.UserSecret{})
  }

  defer db.Close()

  hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

  if err != nil {
    panic(err.Error())
  }

  var id string = uuid.New().String()
  userSecret := models.UserSecret {
    Id: id,
    UserId: userId,
    Hash: string(hash),
  }

  db.NewRecord(userSecret)
  db.Create(&userSecret)

  return true, nil
}
