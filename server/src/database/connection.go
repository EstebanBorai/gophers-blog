package database

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() (*gorm.DB, error) {
  db, err := gorm.Open("mysql", "admin:admin@/songs-share?charset=utf8mb4&parseTime=True&loc=Local")
  if err != nil {
		return nil, err
  }

  defer db.Close()

  return db, nil
}
