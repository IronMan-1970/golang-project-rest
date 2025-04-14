package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	byteArr, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(byteArr), err
}

func CheckCompliance(hashPasword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPasword), []byte(password))
	return err == nil
}
