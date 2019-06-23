package models

// UserSecret represents the User (model) related Hash
type UserSecret struct {
	ID     string `gorm:"not null" json:"id"`
	UserID string `gorm:"size:200" json:"-"`
	Hash   string `gorm:"size:200" json:"-"`
}
