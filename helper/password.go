package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(pass []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ValidatePassword(hash, pass []byte) error {
	return bcrypt.CompareHashAndPassword(hash, pass)
}
