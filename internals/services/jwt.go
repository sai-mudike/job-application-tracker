package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKeys = "supersecret"

func GenerateToken(email string, id int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKeys))
}
