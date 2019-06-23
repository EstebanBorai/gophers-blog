package models

// UserCredentials represent the payload of user's Username/Password
type UserCredentials struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
