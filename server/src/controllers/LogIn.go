package controllers

import (
  models "models"
  "encoding/json"
  jwt "github.com/dgrijalva/jwt-go"
  "fmt"
  "time"
)

var jwtKey = []byte("songs-share-not-safe-key")

type Claims struct {
  UserName string `json:"userName"`
  jwt.StandardClaims
}

type JWTCookie struct {
  Name string
  Value string
  Expires time.Time
}

func LogIn(credentials string) JWTCookie {
  var creds models.UserCredentials

  json.Unmarshal([]byte(credentials), &creds)

  // Check for Password
  fmt.Sprintf("Username: %s, Password: %s", creds.UserName, creds.Password)

  tokenExpirationTime := time.Now().Add(5 * time.Minute)
  
  claims := &Claims {
    UserName: creds.UserName,
    StandardClaims: jwt.StandardClaims {
      ExpiresAt: tokenExpirationTime.Unix(),
    },
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

  tokenString, err := token.SignedString(jwtKey)

  if err != nil {
    panic(err.Error())
  }
  
  return JWTCookie {
    Name: "token",
		Value: tokenString,
		Expires: tokenExpirationTime,
  }
}
