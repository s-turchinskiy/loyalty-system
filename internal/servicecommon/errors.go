package servicecommon

import (
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrNotEnoughBalance                        = fmt.Errorf("not enough balance")
	ErrNoLuhnValidate                          = fmt.Errorf("order number no luhn validate")
	ErrOrderNumberAlreadyUploadedByThisUser    = fmt.Errorf("order number already uploaded by this user")
	ErrOrderNumberAlreadyUploadedByAnotherUser = fmt.Errorf("order number already uploaded by another user")
)

func IsErrorDuplicateKeyValue(err error) bool {

	if err == nil {
		return false
	}

	var pgErr *pgconn.PgError
	errors.As(err, &pgErr)
	if pgErr == nil {
		return false
	}

	return pgErr.Code == pgerrcode.UniqueViolation

}
