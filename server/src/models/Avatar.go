package models

// Avatar represents an User Avatar
type Avatar struct {
	ID     string `gorm:"not null" json:"-"`
	Image  string `sql:"type:text" json:"base64"`
	UserID string `gorm:"not null" json:"-"`
}
