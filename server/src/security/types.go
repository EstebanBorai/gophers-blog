package security

import (
  "time"
  jwt "github.com/dgrijalva/jwt-go"
)

var JwtSecret = []byte("songs-share-not-safe-key")

type Claims struct {
  UserName string `json:"userName"`
  jwt.StandardClaims
}

type JWTCookie struct {
  Name string
  Value string
  Expires time.Time
}
