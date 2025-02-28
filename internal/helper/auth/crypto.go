package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var ErrPasswordLength = errors.New("password length must be between 1 and 72 characters")

func HashPassword(password string) (string, error) {
	if len(password) <= 0 || len(password) > 72 {
		return "", ErrPasswordLength
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CompareHashAndPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
