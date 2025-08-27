package postgresql

import (
	"context"
	"github.com/s-turchinskiy/loyalty-system/internal/servicecommon"
)

func (p *PostgreSQL) NewUser(ctx context.Context, login, hash string) error {

	request, err := getRequest("create_user.sql")
	if err != nil {
		return err
	}

	_, err = p.db.ExecContext(ctx, request, login, hash)
	if servicecommon.IsErrorDuplicateKeyValue(err) {
		return servicecommon.NewErrorUserAlreadyExist(login)
	}
	return err

}
