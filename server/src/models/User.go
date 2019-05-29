package models

import (
  "time"
)

type User struct {
  Id string `gorm:"not null" json:"id"`
  UserName string `gorm:"type:varchar(100);unique_index" json:"userName"`
  FirstName string `gorm:"type:varchar(100);" json:"firstName"`
  LastName string `gorm:"type:varchar(100);" json:"lastName"`
  Birthday time.Time `gorm:"not null" json:"birthday"`
  DateJoined time.Time `gorm:"not null" json:"dateJoined"`
}
