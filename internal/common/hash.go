package common

import (
	"golang.org/x/crypto/bcrypt"
)

func HashFromPassword(password string) (hashedPassword string, err error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil

}

func HashAndPasswordEqual(password, hashedPassword string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil

}
