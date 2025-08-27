package servicecommon

import (
	"errors"
	"fmt"
)

var ErrLoginPasswordIncorrect = errors.New("login-password incorrect")

type ErrorLoginPasswordIncorrect struct {
	login string
}

func (e *ErrorLoginPasswordIncorrect) Error() string {
	return fmt.Sprintf("login-password for user \"%v\" incorrect", e.login)
}

func (e *ErrorLoginPasswordIncorrect) Unwrap() error {
	return ErrLoginPasswordIncorrect
}

func NewErrorLoginPasswordIncorrect(login string) error {
	return &ErrorLoginPasswordIncorrect{
		login: login,
	}
}
