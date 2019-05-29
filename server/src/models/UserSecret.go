package models

type UserSecret struct {
  Id string `gorm:"not null" json:"id"`
  UserId string `gorm:"size:200" json:"-"`
  Hash string `gorm:"size:200" json:"-"`
}
