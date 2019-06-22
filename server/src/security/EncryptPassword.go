package security

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) string {
	pwd := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	if err != nil {
		panic(err.Error())
	}

	return string(hash)
}
