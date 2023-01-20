package middleware

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func CreateAccessToken(permission string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":   "go-jwt-middleware-example",
		"aud":   "audience-example",
		"sub":   "1234567890",
		"iat":   time.Now().Unix(),
		"scope": permission,
	})
	tokenString, err := token.SignedString(signingKey)
	return "Bearer " + tokenString, err
}

// References
// https://pkg.go.dev/github.com/golang-jwt/jwt/v4#example-New-Hmac
