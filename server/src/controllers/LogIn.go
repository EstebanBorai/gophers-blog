package controllers

import (
  "fmt"
  security "security"
  models "models"
  "encoding/json"
)

func LogIn(credentials string) security.JWTCookie {
  var creds models.UserCredentials

  json.Unmarshal([]byte(credentials), &creds)

  // Check for Password
  fmt.Println(fmt.Sprintf("Username: %s, Password: %s", creds.UserName, creds.Password))
  
  return security.CreateToken(creds.UserName)
}
