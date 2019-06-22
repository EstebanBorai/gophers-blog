package security

import (
	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(hash string, plain string) bool {
	byteHash := []byte(hash)
	bytePlain := []byte(plain)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)

	if err != nil {
		return false
	}

	return true
}
