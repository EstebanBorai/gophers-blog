package models

import (
  "time"
)

type User struct {
	Id string `gorm:"not null" json:"id"`
	Avatar string `sql:"type:text" json:"avatar"`
  UserName string `gorm:"type:varchar(100);not null;unique_index" json:"userName"`
  FirstName string `gorm:"type:varchar(100);not null;" json:"firstName"`
  LastName string `gorm:"type:varchar(100);not null;" json:"lastName"`
  Email string `gorm:"type:varchar(190);not null;unique_index" json:"email"`
  Birthday time.Time `gorm:"not null" json:"birthday"`
  DateJoined time.Time `gorm:"not null" json:"dateJoined"`
	UserSecret UserSecret `gorm:"foreignkey:UserId" json:"-"`
}
