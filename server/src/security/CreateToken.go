package security

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken creates the JWT token with the given properties as
// the JWT token claims
func CreateToken(userName string, email string, id string) JWTCookie {
	tokenExpirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		ID:       id,
		Email:    email,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JwtSecret)

	if err != nil {
		panic(err.Error())
	}

	return JWTCookie{
		Name:    "token",
		Value:   tokenString,
		Expires: tokenExpirationTime,
	}
}
