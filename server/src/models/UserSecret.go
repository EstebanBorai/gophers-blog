package models

import (
  "github.com/jinzhu/gorm"
)

type UserSecret struct {
  gorm.Model
  Id string `gorm:"not null" json:"id"`
  UserId string `gorm:"size:200" json:"-"`
  Hash string `gorm:"size:200" json:"-"`
}
