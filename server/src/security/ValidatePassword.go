package security

import (
	"golang.org/x/crypto/bcrypt"
)

// ValidatePassword verifies a stored hash with a given plain text password
func ValidatePassword(hash string, plain string) bool {
	byteHash := []byte(hash)
	bytePlain := []byte(plain)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)

	if err != nil {
		return false
	}

	return true
}
