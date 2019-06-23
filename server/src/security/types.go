package security

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtSecret represents the JWT Signature
var JwtSecret = []byte("songs-share-not-safe-key")

// Claims represents the Claims of the JWT generated
// at the login
type Claims struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"userName"`
	jwt.StandardClaims
}

// JWTCookie represents the properties of the JWTCookie
// generated at the login
type JWTCookie struct {
	Name    string
	Value   string
	Expires time.Time
}
