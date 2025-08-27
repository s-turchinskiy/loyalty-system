package servicecommon

import (
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

var ErrorNotEnoughBalance = fmt.Errorf("not enough balance")

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
