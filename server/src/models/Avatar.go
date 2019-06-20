package models

type Avatar struct {
	Id string `gorm:"not null" json:"-"`
	Image string `sql:"type:text" json:"base64"`
	UserId string `gorm:"not null" json:"-"`
}
