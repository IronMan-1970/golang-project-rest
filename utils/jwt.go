package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "SuperSecretKey"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secret))
}

func ValidateToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Token isn't valid")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, errors.New("Could not parse token")
	}

	isTokenValid := parsedToken.Valid

	if !isTokenValid {
		return 0, errors.New("Ivalid token!")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Token isn't valid")
	}

	//email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil

}
