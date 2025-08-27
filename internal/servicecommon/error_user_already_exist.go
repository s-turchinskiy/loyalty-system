package servicecommon

import (
	"errors"
	"fmt"
)

var ErrUserAlreadyExist = errors.New("user already exist")

type TypeErrorUserAlreadyExist struct {
	login string
}

func (e *TypeErrorUserAlreadyExist) Error() string {
	return fmt.Sprintf("user \"%v\" already exist", e.login)
}

func (te *TypeErrorUserAlreadyExist) Unwrap() error {
	return ErrUserAlreadyExist
}

func NewErrorUserAlreadyExist(login string) error {
	return &TypeErrorUserAlreadyExist{
		login: login,
	}
}
