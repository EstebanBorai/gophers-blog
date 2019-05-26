package queries

import (
  "golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd string) string {
  hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
  
  if err != nil {
    panic(err.Error())
  }

  return string(hash)
}

func CheckPassword(hash string, pwd string) bool {
  byteHash := []byte(hash)
  bytePwd := []byte(pwd)

  err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)

  if err != nil {
    return false
  }

  return true
}
