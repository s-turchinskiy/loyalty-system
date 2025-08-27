package servicecommon

import (
	"errors"
	"fmt"
)

var ErrUserNoExist = errors.New("user no exist")

type ErrorUserNoExist struct {
	login string
}

func (e *ErrorUserNoExist) Error() string {
	return fmt.Sprintf("user \"%v\" no exist", e.login)
}

func (e *ErrorUserNoExist) Unwrap() error {
	return ErrUserNoExist
}

func NewErrorUserNoExist(login string) error {
	return &ErrorUserAlreadyExist{
		login: login,
	}
}
