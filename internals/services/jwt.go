package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	appErrors "github.com/sai-mudike/job-application-tracker/internals/errors"
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

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(secretKeys), nil
	})
	if err != nil {
		return 0, err
	}

	isTokenValid := parsedToken.Valid

	if !isTokenValid {
		return 0, appErrors.ErrInvalidToken
	}

	data, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, appErrors.ErrTokenExpired
	}

	userID := data["id"].(float64)

	return int64(userID), nil
}
