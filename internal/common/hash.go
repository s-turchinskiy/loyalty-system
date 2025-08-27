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

func HashAndPasswordEqual(password, hashedPassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil

}
