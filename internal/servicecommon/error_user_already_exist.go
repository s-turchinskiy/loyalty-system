package servicecommon

import (
	"errors"
	"fmt"
)

var ErrUserAlreadyExist = errors.New("user already exist")

type ErrorUserAlreadyExist struct {
	login string
}

func (e *ErrorUserAlreadyExist) Error() string {
	return fmt.Sprintf("user \"%v\" already exist", e.login)
}

func (e *ErrorUserAlreadyExist) Unwrap() error {
	return ErrUserAlreadyExist
}

func NewErrorUserAlreadyExist(login string) error {
	return &ErrorUserAlreadyExist{
		login: login,
	}
}
